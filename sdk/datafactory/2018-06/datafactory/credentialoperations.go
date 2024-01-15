package datafactory

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// CredentialOperationsClient is the the Azure Data Factory V2 management API provides a RESTful set of web services
// that interact with Azure Data Factory V2 services.
type CredentialOperationsClient struct {
	BaseClient
}

// NewCredentialOperationsClient creates an instance of the CredentialOperationsClient client.
func NewCredentialOperationsClient(subscriptionID string) CredentialOperationsClient {
	return NewCredentialOperationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCredentialOperationsClientWithBaseURI creates an instance of the CredentialOperationsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewCredentialOperationsClientWithBaseURI(baseURI string, subscriptionID string) CredentialOperationsClient {
	return CredentialOperationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a credential.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// credentialName - credential name
// credential - credential resource definition.
// ifMatch - eTag of the credential entity. Should only be specified for update, for which it should match
// existing entity or can be * for unconditional update.
func (client CredentialOperationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, credentialName string, credential ManagedIdentityCredentialResource, ifMatch string) (result ManagedIdentityCredentialResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CredentialOperationsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: credentialName,
			Constraints: []validation.Constraint{{Target: "credentialName", Name: validation.MaxLength, Rule: 127, Chain: nil},
				{Target: "credentialName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "credentialName", Name: validation.Pattern, Rule: `^([_A-Za-z0-9]|([_A-Za-z0-9][-_A-Za-z0-9]{0,125}[_A-Za-z0-9]))$`, Chain: nil}}},
		{TargetValue: credential,
			Constraints: []validation.Constraint{{Target: "credential.Properties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.CredentialOperationsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, factoryName, credentialName, credential, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client CredentialOperationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, factoryName string, credentialName string, credential ManagedIdentityCredentialResource, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"credentialName":    autorest.Encode("path", credentialName),
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/credentials/{credentialName}", pathParameters),
		autorest.WithJSON(credential),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client CredentialOperationsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client CredentialOperationsClient) CreateOrUpdateResponder(resp *http.Response) (result ManagedIdentityCredentialResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a credential.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// credentialName - credential name
func (client CredentialOperationsClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, credentialName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CredentialOperationsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: credentialName,
			Constraints: []validation.Constraint{{Target: "credentialName", Name: validation.MaxLength, Rule: 127, Chain: nil},
				{Target: "credentialName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "credentialName", Name: validation.Pattern, Rule: `^([_A-Za-z0-9]|([_A-Za-z0-9][-_A-Za-z0-9]{0,125}[_A-Za-z0-9]))$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.CredentialOperationsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, factoryName, credentialName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CredentialOperationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, factoryName string, credentialName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"credentialName":    autorest.Encode("path", credentialName),
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/credentials/{credentialName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CredentialOperationsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CredentialOperationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a credential.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
// credentialName - credential name
// ifNoneMatch - eTag of the credential entity. Should only be specified for get. If the ETag matches the
// existing entity tag, or if * was provided, then no content will be returned.
func (client CredentialOperationsClient) Get(ctx context.Context, resourceGroupName string, factoryName string, credentialName string, ifNoneMatch string) (result ManagedIdentityCredentialResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CredentialOperationsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}},
		{TargetValue: credentialName,
			Constraints: []validation.Constraint{{Target: "credentialName", Name: validation.MaxLength, Rule: 127, Chain: nil},
				{Target: "credentialName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "credentialName", Name: validation.Pattern, Rule: `^([_A-Za-z0-9]|([_A-Za-z0-9][-_A-Za-z0-9]{0,125}[_A-Za-z0-9]))$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.CredentialOperationsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, factoryName, credentialName, ifNoneMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CredentialOperationsClient) GetPreparer(ctx context.Context, resourceGroupName string, factoryName string, credentialName string, ifNoneMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"credentialName":    autorest.Encode("path", credentialName),
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/credentials/{credentialName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifNoneMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-None-Match", autorest.String(ifNoneMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CredentialOperationsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CredentialOperationsClient) GetResponder(resp *http.Response) (result ManagedIdentityCredentialResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNotModified),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByFactory list credentials.
// Parameters:
// resourceGroupName - the resource group name.
// factoryName - the factory name.
func (client CredentialOperationsClient) ListByFactory(ctx context.Context, resourceGroupName string, factoryName string) (result CredentialListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CredentialOperationsClient.ListByFactory")
		defer func() {
			sc := -1
			if result.clr.Response.Response != nil {
				sc = result.clr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: factoryName,
			Constraints: []validation.Constraint{{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("datafactory.CredentialOperationsClient", "ListByFactory", err.Error())
	}

	result.fn = client.listByFactoryNextResults
	req, err := client.ListByFactoryPreparer(ctx, resourceGroupName, factoryName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "ListByFactory", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByFactorySender(req)
	if err != nil {
		result.clr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "ListByFactory", resp, "Failure sending request")
		return
	}

	result.clr, err = client.ListByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "ListByFactory", resp, "Failure responding to request")
		return
	}
	if result.clr.hasNextLink() && result.clr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByFactoryPreparer prepares the ListByFactory request.
func (client CredentialOperationsClient) ListByFactoryPreparer(ctx context.Context, resourceGroupName string, factoryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"factoryName":       autorest.Encode("path", factoryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/credentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByFactorySender sends the ListByFactory request. The method will close the
// http.Response Body if it receives an error.
func (client CredentialOperationsClient) ListByFactorySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByFactoryResponder handles the response to the ListByFactory request. The method always
// closes the http.Response Body.
func (client CredentialOperationsClient) ListByFactoryResponder(resp *http.Response) (result CredentialListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByFactoryNextResults retrieves the next set of results, if any.
func (client CredentialOperationsClient) listByFactoryNextResults(ctx context.Context, lastResults CredentialListResponse) (result CredentialListResponse, err error) {
	req, err := lastResults.credentialListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "listByFactoryNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByFactorySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "listByFactoryNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByFactoryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datafactory.CredentialOperationsClient", "listByFactoryNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByFactoryComplete enumerates all values, automatically crossing page boundaries as required.
func (client CredentialOperationsClient) ListByFactoryComplete(ctx context.Context, resourceGroupName string, factoryName string) (result CredentialListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CredentialOperationsClient.ListByFactory")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByFactory(ctx, resourceGroupName, factoryName)
	return
}