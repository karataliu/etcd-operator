package cdn

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

// ProfilesClient is the use these APIs to manage Azure CDN resources through the Azure Resource Manager. You must make
// sure that requests made to these resources are secure. For more information, see
// https://msdn.microsoft.com/en-us/library/azure/dn790557.aspx.
type ProfilesClient struct {
	BaseClient
}

// NewProfilesClient creates an instance of the ProfilesClient client.
func NewProfilesClient(subscriptionID string) ProfilesClient {
	return NewProfilesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProfilesClientWithBaseURI creates an instance of the ProfilesClient client.
func NewProfilesClientWithBaseURI(baseURI string, subscriptionID string) ProfilesClient {
	return ProfilesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create sends the create request.
//
// profileName is name of the CDN profile within the resource group. profileProperties is profile properties needed
// for creation. resourceGroupName is name of the resource group within the Azure subscription.
func (client ProfilesClient) Create(ctx context.Context, profileName string, profileProperties ProfileCreateParameters, resourceGroupName string) (result ProfilesCreateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: profileProperties,
			Constraints: []validation.Constraint{{Target: "profileProperties.Location", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "profileProperties.ProfilePropertiesCreateParameters", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "profileProperties.ProfilePropertiesCreateParameters.Sku", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("cdn.ProfilesClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, profileName, profileProperties, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ProfilesClient) CreatePreparer(ctx context.Context, profileName string, profileProperties ProfileCreateParameters, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", pathParameters),
		autorest.WithJSON(profileProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) CreateSender(req *http.Request) (future ProfilesCreateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted))
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ProfilesClient) CreateResponder(resp *http.Response) (result Profile, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteIfExists sends the delete if exists request.
//
// profileName is name of the CDN profile within the resource group. resourceGroupName is name of the resource
// group within the Azure subscription.
func (client ProfilesClient) DeleteIfExists(ctx context.Context, profileName string, resourceGroupName string) (result ProfilesDeleteIfExistsFuture, err error) {
	req, err := client.DeleteIfExistsPreparer(ctx, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "DeleteIfExists", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteIfExistsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "DeleteIfExists", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeleteIfExistsPreparer prepares the DeleteIfExists request.
func (client ProfilesClient) DeleteIfExistsPreparer(ctx context.Context, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteIfExistsSender sends the DeleteIfExists request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) DeleteIfExistsSender(req *http.Request) (future ProfilesDeleteIfExistsFuture, err error) {
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

// DeleteIfExistsResponder handles the response to the DeleteIfExists request. The method always
// closes the http.Response Body.
func (client ProfilesClient) DeleteIfExistsResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GenerateSsoURI sends the generate sso uri request.
//
// profileName is name of the CDN profile within the resource group. resourceGroupName is name of the resource
// group within the Azure subscription.
func (client ProfilesClient) GenerateSsoURI(ctx context.Context, profileName string, resourceGroupName string) (result SsoURI, err error) {
	req, err := client.GenerateSsoURIPreparer(ctx, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "GenerateSsoURI", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateSsoURISender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "GenerateSsoURI", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateSsoURIResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "GenerateSsoURI", resp, "Failure responding to request")
	}

	return
}

// GenerateSsoURIPreparer prepares the GenerateSsoURI request.
func (client ProfilesClient) GenerateSsoURIPreparer(ctx context.Context, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/generateSsoUri", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GenerateSsoURISender sends the GenerateSsoURI request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) GenerateSsoURISender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GenerateSsoURIResponder handles the response to the GenerateSsoURI request. The method always
// closes the http.Response Body.
func (client ProfilesClient) GenerateSsoURIResponder(resp *http.Response) (result SsoURI, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get sends the get request.
//
// profileName is name of the CDN profile within the resource group. resourceGroupName is name of the resource
// group within the Azure subscription.
func (client ProfilesClient) Get(ctx context.Context, profileName string, resourceGroupName string) (result Profile, err error) {
	req, err := client.GetPreparer(ctx, profileName, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProfilesClient) GetPreparer(ctx context.Context, profileName string, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProfilesClient) GetResponder(resp *http.Response) (result Profile, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup sends the list by resource group request.
//
// resourceGroupName is name of the resource group within the Azure subscription.
func (client ProfilesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ProfileListResult, err error) {
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ProfilesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ProfilesClient) ListByResourceGroupResponder(resp *http.Response) (result ProfileListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscriptionID sends the list by subscription id request.
func (client ProfilesClient) ListBySubscriptionID(ctx context.Context) (result ProfileListResult, err error) {
	req, err := client.ListBySubscriptionIDPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListBySubscriptionID", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListBySubscriptionID", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "ListBySubscriptionID", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionIDPreparer prepares the ListBySubscriptionID request.
func (client ProfilesClient) ListBySubscriptionIDPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/profiles", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionIDSender sends the ListBySubscriptionID request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) ListBySubscriptionIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionIDResponder handles the response to the ListBySubscriptionID request. The method always
// closes the http.Response Body.
func (client ProfilesClient) ListBySubscriptionIDResponder(resp *http.Response) (result ProfileListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update sends the update request.
//
// profileName is name of the CDN profile within the resource group. profileProperties is profile properties needed
// for update. resourceGroupName is name of the resource group within the Azure subscription.
func (client ProfilesClient) Update(ctx context.Context, profileName string, profileProperties ProfileUpdateParameters, resourceGroupName string) (result ProfilesUpdateFuture, err error) {
	req, err := client.UpdatePreparer(ctx, profileName, profileProperties, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.ProfilesClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ProfilesClient) UpdatePreparer(ctx context.Context, profileName string, profileProperties ProfileUpdateParameters, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}", pathParameters),
		autorest.WithJSON(profileProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ProfilesClient) UpdateSender(req *http.Request) (future ProfilesUpdateFuture, err error) {
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
func (client ProfilesClient) UpdateResponder(resp *http.Response) (result Profile, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
