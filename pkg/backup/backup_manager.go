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

package backup

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/coreos/etcd-operator/pkg/backup/writer"
	"github.com/coreos/etcd-operator/pkg/util/constants"

	"github.com/coreos/etcd/clientv3"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

// BackupManager backups an etcd cluster.
type BackupManager struct {
	kubecli kubernetes.Interface

	endpoints     []string
	namespace     string
	etcdTLSConfig *tls.Config

	bw writer.Writer
}

// NewBackupManagerFromWriter creates a BackupManager with backup writer.
func NewBackupManagerFromWriter(kubecli kubernetes.Interface, bw writer.Writer, tc *tls.Config, endpoints []string, namespace string) *BackupManager {
	return &BackupManager{
		kubecli:       kubecli,
		endpoints:     endpoints,
		namespace:     namespace,
		etcdTLSConfig: tc,
		bw:            bw,
	}
}

// PurgeBackup used the s3Path as prefix, to purge stale backups more than maxBackups count
func (bm *BackupManager) PurgeBackup(s3Path string, maxBackups int) error {
	return bm.bw.Purge(s3Path, maxBackups)
}

// SaveSnap uses backup writer to save etcd snapshot to a specified S3 path
// and returns backup etcd server's kv store revision and its version.
// appendRev specify whether we want to append Rev to the s3Path
func (bm *BackupManager) SaveSnap(s3Path string, appendRev bool) (int64, string, error) {
	etcdcli, rev, err := bm.etcdClientWithMaxRevision()
	if err != nil {
		return 0, "", fmt.Errorf("create etcd client failed: %v", err)
	}
	defer etcdcli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultRequestTimeout)
	resp, err := etcdcli.Status(ctx, etcdcli.Endpoints()[0])
	cancel()
	if err != nil {
		return 0, "", fmt.Errorf("failed to retrieve etcd version from the status call: %v", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), constants.DefaultSnapshotTimeout)
	defer cancel() // Can't cancel() after Snapshot() because that will close the reader.
	rc, err := etcdcli.Snapshot(ctx)
	if err != nil {
		return 0, "", fmt.Errorf("failed to receive snapshot (%v)", err)
	}
	defer rc.Close()

	_, err = bm.bw.Write(
		appendRevToPath(appendRev, rev, s3Path),
		rc)
	if err != nil {
		return 0, "", fmt.Errorf("failed to write snapshot (%v)", err)
	}
	return rev, resp.Version, nil
}

func appendRevToPath(appendRev bool, rev int64, path string) string {
	if !appendRev {
		return path
	}
	return fmt.Sprintf("%s_%016x", path, rev)
}

// etcdClientWithMaxRevision gets the etcd endpoint with the maximum kv store revision
// and returns the etcd client of that member.
func (bm *BackupManager) etcdClientWithMaxRevision() (*clientv3.Client, int64, error) {
	etcdcli, rev, err := getClientWithMaxRev(bm.endpoints, bm.etcdTLSConfig)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get etcd client with maximum kv store revision: %v", err)
	}
	return etcdcli, rev, nil
}

func getClientWithMaxRev(endpoints []string, tc *tls.Config) (*clientv3.Client, int64, error) {
	mapEps := make(map[string]*clientv3.Client)
	var maxClient *clientv3.Client
	maxRev := int64(0)
	errors := make([]string, 0)
	for _, endpoint := range endpoints {
		cfg := clientv3.Config{
			Endpoints:   []string{endpoint},
			DialTimeout: constants.DefaultDialTimeout,
			TLS:         tc,
		}
		etcdcli, err := clientv3.New(cfg)
		if err != nil {
			errors = append(errors, fmt.Sprintf("failed to create etcd client for endpoint (%v): %v", endpoint, err))
			continue
		}
		mapEps[endpoint] = etcdcli

		ctx, cancel := context.WithTimeout(context.Background(), constants.DefaultRequestTimeout)
		resp, err := etcdcli.Get(ctx, "/", clientv3.WithSerializable())
		cancel()
		if err != nil {
			errors = append(errors, fmt.Sprintf("failed to get revision from endpoint (%s)", endpoint))
			continue
		}

		logrus.Infof("getMaxRev: endpoint %s revision (%d)", endpoint, resp.Header.Revision)
		if resp.Header.Revision > maxRev {
			maxRev = resp.Header.Revision
			maxClient = etcdcli
		}
	}

	// close all open clients that are not maxClient.
	for _, cli := range mapEps {
		if cli == maxClient {
			continue
		}
		cli.Close()
	}

	var err error
	if len(errors) > 0 {
		errorStr := ""
		for _, errStr := range errors {
			errorStr += errStr + "\n"
		}
		err = fmt.Errorf(errorStr)
	}

	return maxClient, maxRev, err
}
