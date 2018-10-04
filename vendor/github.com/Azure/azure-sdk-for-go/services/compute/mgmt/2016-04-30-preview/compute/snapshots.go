package compute

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// SnapshotsClient is the compute Client
type SnapshotsClient struct {
	BaseClient
}

// NewSnapshotsClient creates an instance of the SnapshotsClient client.
func NewSnapshotsClient(subscriptionID string) SnapshotsClient {
	return NewSnapshotsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSnapshotsClientWithBaseURI creates an instance of the SnapshotsClient client.
func NewSnapshotsClientWithBaseURI(baseURI string, subscriptionID string) SnapshotsClient {
	return SnapshotsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a snapshot.
//
// resourceGroupName is the name of the resource group. snapshotName is the name of the snapshot within the given
// subscription and resource group. snapshot is snapshot object supplied in the body of the Put disk operation.
func (client SnapshotsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, snapshotName string, snapshot Snapshot) (result SnapshotsCreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: snapshot,
			Constraints: []validation.Constraint{{Target: "snapshot.DiskProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "snapshot.DiskProperties.CreationData", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "snapshot.DiskProperties.CreationData.ImageReference", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "snapshot.DiskProperties.CreationData.ImageReference.ID", Name: validation.Null, Rule: true, Chain: nil}}},
					}},
					{Target: "snapshot.DiskProperties.EncryptionSettings", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "snapshot.DiskProperties.EncryptionSettings.DiskEncryptionKey", Name: validation.Null, Rule: false,
							Chain: []validation.Constraint{{Target: "snapshot.DiskProperties.EncryptionSettings.DiskEncryptionKey.SourceVault", Name: validation.Null, Rule: true, Chain: nil},
								{Target: "snapshot.DiskProperties.EncryptionSettings.DiskEncryptionKey.SecretURL", Name: validation.Null, Rule: true, Chain: nil},
							}},
							{Target: "snapshot.DiskProperties.EncryptionSettings.KeyEncryptionKey", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "snapshot.DiskProperties.EncryptionSettings.KeyEncryptionKey.SourceVault", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "snapshot.DiskProperties.EncryptionSettings.KeyEncryptionKey.KeyURL", Name: validation.Null, Rule: true, Chain: nil},
								}},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("compute.SnapshotsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, snapshotName, snapshot)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SnapshotsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, snapshotName string, snapshot Snapshot) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"snapshotName":      autorest.Encode("path", snapshotName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}", pathParameters),
		autorest.WithJSON(snapshot),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotsClient) CreateOrUpdateSender(req *http.Request) (future SnapshotsCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SnapshotsClient) CreateOrUpdateResponder(resp *http.Response) (result Snapshot, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a snapshot.
//
// resourceGroupName is the name of the resource group. snapshotName is the name of the snapshot within the given
// subscription and resource group.
func (client SnapshotsClient) Delete(ctx context.Context, resourceGroupName string, snapshotName string) (result SnapshotsDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, snapshotName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SnapshotsClient) DeletePreparer(ctx context.Context, resourceGroupName string, snapshotName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"snapshotName":      autorest.Encode("path", snapshotName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotsClient) DeleteSender(req *http.Request) (future SnapshotsDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SnapshotsClient) DeleteResponder(resp *http.Response) (result OperationStatusResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets information about a snapshot.
//
// resourceGroupName is the name of the resource group. snapshotName is the name of the snapshot within the given
// subscription and resource group.
func (client SnapshotsClient) Get(ctx context.Context, resourceGroupName string, snapshotName string) (result Snapshot, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, snapshotName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client SnapshotsClient) GetPreparer(ctx context.Context, resourceGroupName string, snapshotName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"snapshotName":      autorest.Encode("path", snapshotName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SnapshotsClient) GetResponder(resp *http.Response) (result Snapshot, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GrantAccess grants access to a snapshot.
//
// resourceGroupName is the name of the resource group. snapshotName is the name of the snapshot within the given
// subscription and resource group. grantAccessData is access data object supplied in the body of the get snapshot
// access operation.
func (client SnapshotsClient) GrantAccess(ctx context.Context, resourceGroupName string, snapshotName string, grantAccessData GrantAccessData) (result SnapshotsGrantAccessFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: grantAccessData,
			Constraints: []validation.Constraint{{Target: "grantAccessData.DurationInSeconds", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("compute.SnapshotsClient", "GrantAccess", err.Error())
	}

	req, err := client.GrantAccessPreparer(ctx, resourceGroupName, snapshotName, grantAccessData)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "GrantAccess", nil, "Failure preparing request")
		return
	}

	result, err = client.GrantAccessSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "GrantAccess", result.Response(), "Failure sending request")
		return
	}

	return
}

// GrantAccessPreparer prepares the GrantAccess request.
func (client SnapshotsClient) GrantAccessPreparer(ctx context.Context, resourceGroupName string, snapshotName string, grantAccessData GrantAccessData) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"snapshotName":      autorest.Encode("path", snapshotName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}/beginGetAccess", pathParameters),
		autorest.WithJSON(grantAccessData),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GrantAccessSender sends the GrantAccess request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotsClient) GrantAccessSender(req *http.Request) (future SnapshotsGrantAccessFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// GrantAccessResponder handles the response to the GrantAccess request. The method always
// closes the http.Response Body.
func (client SnapshotsClient) GrantAccessResponder(resp *http.Response) (result AccessURI, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists snapshots under a subscription.
func (client SnapshotsClient) List(ctx context.Context) (result SnapshotListPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.sl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "List", resp, "Failure sending request")
		return
	}

	result.sl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client SnapshotsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/snapshots", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SnapshotsClient) ListResponder(resp *http.Response) (result SnapshotList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SnapshotsClient) listNextResults(lastResults SnapshotList) (result SnapshotList, err error) {
	req, err := lastResults.snapshotListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.SnapshotsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.SnapshotsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SnapshotsClient) ListComplete(ctx context.Context) (result SnapshotListIterator, err error) {
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup lists snapshots under a resource group.
//
// resourceGroupName is the name of the resource group.
func (client SnapshotsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result SnapshotListPage, err error) {
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.sl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.sl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client SnapshotsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client SnapshotsClient) ListByResourceGroupResponder(resp *http.Response) (result SnapshotList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client SnapshotsClient) listByResourceGroupNextResults(lastResults SnapshotList) (result SnapshotList, err error) {
	req, err := lastResults.snapshotListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "compute.SnapshotsClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "compute.SnapshotsClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SnapshotsClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result SnapshotListIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// RevokeAccess revokes access to a snapshot.
//
// resourceGroupName is the name of the resource group. snapshotName is the name of the snapshot within the given
// subscription and resource group.
func (client SnapshotsClient) RevokeAccess(ctx context.Context, resourceGroupName string, snapshotName string) (result SnapshotsRevokeAccessFuture, err error) {
	req, err := client.RevokeAccessPreparer(ctx, resourceGroupName, snapshotName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "RevokeAccess", nil, "Failure preparing request")
		return
	}

	result, err = client.RevokeAccessSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "RevokeAccess", result.Response(), "Failure sending request")
		return
	}

	return
}

// RevokeAccessPreparer prepares the RevokeAccess request.
func (client SnapshotsClient) RevokeAccessPreparer(ctx context.Context, resourceGroupName string, snapshotName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"snapshotName":      autorest.Encode("path", snapshotName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}/endGetAccess", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RevokeAccessSender sends the RevokeAccess request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotsClient) RevokeAccessSender(req *http.Request) (future SnapshotsRevokeAccessFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// RevokeAccessResponder handles the response to the RevokeAccess request. The method always
// closes the http.Response Body.
func (client SnapshotsClient) RevokeAccessResponder(resp *http.Response) (result OperationStatusResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates (patches) a snapshot.
//
// resourceGroupName is the name of the resource group. snapshotName is the name of the snapshot within the given
// subscription and resource group. snapshot is snapshot object supplied in the body of the Patch snapshot
// operation.
func (client SnapshotsClient) Update(ctx context.Context, resourceGroupName string, snapshotName string, snapshot SnapshotUpdate) (result SnapshotsUpdateFuture, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, snapshotName, snapshot)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.SnapshotsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client SnapshotsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, snapshotName string, snapshot SnapshotUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"snapshotName":      autorest.Encode("path", snapshotName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-30-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}", pathParameters),
		autorest.WithJSON(snapshot),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SnapshotsClient) UpdateSender(req *http.Request) (future SnapshotsUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SnapshotsClient) UpdateResponder(resp *http.Response) (result Snapshot, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
