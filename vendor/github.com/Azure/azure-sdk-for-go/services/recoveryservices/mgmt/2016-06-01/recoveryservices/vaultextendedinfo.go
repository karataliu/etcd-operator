package recoveryservices

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
	"net/http"
)

// VaultExtendedInfoClient is the recovery Services Client
type VaultExtendedInfoClient struct {
	BaseClient
}

// NewVaultExtendedInfoClient creates an instance of the VaultExtendedInfoClient client.
func NewVaultExtendedInfoClient(subscriptionID string) VaultExtendedInfoClient {
	return NewVaultExtendedInfoClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVaultExtendedInfoClientWithBaseURI creates an instance of the VaultExtendedInfoClient client.
func NewVaultExtendedInfoClientWithBaseURI(baseURI string, subscriptionID string) VaultExtendedInfoClient {
	return VaultExtendedInfoClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create vault extended info.
//
// resourceGroupName is the name of the resource group where the recovery services vault is present. vaultName is
// the name of the recovery services vault. resourceResourceExtendedInfoDetails is details of ResourceExtendedInfo
func (client VaultExtendedInfoClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, resourceResourceExtendedInfoDetails VaultExtendedInfoResource) (result VaultExtendedInfoResource, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, vaultName, resourceResourceExtendedInfoDetails)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.VaultExtendedInfoClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservices.VaultExtendedInfoClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.VaultExtendedInfoClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VaultExtendedInfoClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, vaultName string, resourceResourceExtendedInfoDetails VaultExtendedInfoResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo", pathParameters),
		autorest.WithJSON(resourceResourceExtendedInfoDetails),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VaultExtendedInfoClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VaultExtendedInfoClient) CreateOrUpdateResponder(resp *http.Response) (result VaultExtendedInfoResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get the vault extended info.
//
// resourceGroupName is the name of the resource group where the recovery services vault is present. vaultName is
// the name of the recovery services vault.
func (client VaultExtendedInfoClient) Get(ctx context.Context, resourceGroupName string, vaultName string) (result VaultExtendedInfoResource, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, vaultName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.VaultExtendedInfoClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservices.VaultExtendedInfoClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.VaultExtendedInfoClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VaultExtendedInfoClient) GetPreparer(ctx context.Context, resourceGroupName string, vaultName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VaultExtendedInfoClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VaultExtendedInfoClient) GetResponder(resp *http.Response) (result VaultExtendedInfoResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update vault extended info.
//
// resourceGroupName is the name of the resource group where the recovery services vault is present. vaultName is
// the name of the recovery services vault. resourceResourceExtendedInfoDetails is details of ResourceExtendedInfo
func (client VaultExtendedInfoClient) Update(ctx context.Context, resourceGroupName string, vaultName string, resourceResourceExtendedInfoDetails VaultExtendedInfoResource) (result VaultExtendedInfoResource, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, vaultName, resourceResourceExtendedInfoDetails)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.VaultExtendedInfoClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservices.VaultExtendedInfoClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservices.VaultExtendedInfoClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client VaultExtendedInfoClient) UpdatePreparer(ctx context.Context, resourceGroupName string, vaultName string, resourceResourceExtendedInfoDetails VaultExtendedInfoResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/extendedInformation/vaultExtendedInfo", pathParameters),
		autorest.WithJSON(resourceResourceExtendedInfoDetails),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client VaultExtendedInfoClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client VaultExtendedInfoClient) UpdateResponder(resp *http.Response) (result VaultExtendedInfoResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
