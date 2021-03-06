// Copyright 2017 The etcd-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package writer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"sort"

	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/coreos/etcd-operator/pkg/backup/util"
	"github.com/pborman/uuid"
)

var _ Writer = &absWriter{}

type absWriter struct {
	abs *storage.BlobStorageClient
}

const (
	// AzureBlobBlockChunkLimitInBytes 100MiB is the limit
	AzureBlobBlockChunkLimitInBytes = 104857600
)

// NewABSWriter creates a abs writer.
func NewABSWriter(abs *storage.BlobStorageClient) Writer {
	return &absWriter{abs}
}

func (absw *absWriter) getContainer(container string) (*storage.Container, error) {
	containerRef := absw.abs.GetContainerReference(container)
	containerExists, err := containerRef.Exists()
	if err != nil {
		return nil, err
	}

	if !containerExists {
		return nil, fmt.Errorf("container %v does not exist", container)
	}
	return containerRef, nil
}

// Write writes the backup file to the given abs path, "<abs-container-name>/<key>".
func (absw *absWriter) Write(path string, r io.Reader) (int64, error) {
	container, key, err := util.ParseBucketAndKey(path)
	if err != nil {
		return 0, err
	}
	containerRef, err := absw.getContainer(container)
	if err != nil {
		return 0, err
	}

	blob := containerRef.GetBlobReference(key)
	putBlobOpts := storage.PutBlobOptions{}

	err = blob.CreateBlockBlob(&putBlobOpts)
	if err != nil {
		return 0, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	len := len(buf.Bytes())
	chunckCount := len/AzureBlobBlockChunkLimitInBytes + 1
	blocks := make([]storage.Block, 0, chunckCount)
	for i := 0; i < chunckCount; i++ {
		blockID := base64.StdEncoding.EncodeToString([]byte(uuid.New()))
		blocks = append(blocks, storage.Block{ID: blockID, Status: storage.BlockStatusLatest})
		start := i * AzureBlobBlockChunkLimitInBytes
		end := (i + 1) * AzureBlobBlockChunkLimitInBytes
		if len < end {
			end = len
		}

		chunk := buf.Bytes()[start:end]
		err = blob.PutBlock(blockID, chunk, &storage.PutBlockOptions{})
		if err != nil {
			return 0, err
		}
	}

	err = blob.PutBlockList(blocks, &storage.PutBlockListOptions{})
	if err != nil {
		return 0, err
	}

	getBlobOpts := &storage.GetBlobOptions{}
	_, err = blob.Get(getBlobOpts)
	if err != nil {
		return 0, err
	}

	return blob.Properties.ContentLength, nil
}

func (absw *absWriter) Purge(path string, maxBackups int) error {
	container, key, err := util.ParseBucketAndKey(path)
	if err != nil {
		return err
	}

	containerRef, err := absw.getContainer(container)
	if err != nil {
		return err
	}

	params := storage.ListBlobsParameters{Prefix: fmt.Sprintf("%s_", key)}
	resp, err := containerRef.ListBlobs(params)
	if err != nil {
		return err
	}

	blobNames := []string{}
	for _, blob := range resp.Blobs {
		blobNames = append(blobNames, (blob.Name))
	}

	// we can just use string comparison
	sort.Strings(blobNames)
	for i := 0; i < len(blobNames)-maxBackups; i++ {
		blob := containerRef.GetBlobReference(blobNames[i])
		err = blob.Delete(&storage.DeleteBlobOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
