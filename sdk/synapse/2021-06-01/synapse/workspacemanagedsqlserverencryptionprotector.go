package synapse

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

// WorkspaceManagedSQLServerEncryptionProtectorClient is the azure Synapse Analytics Management Client
type WorkspaceManagedSQLServerEncryptionProtectorClient struct {
	BaseClient
}

// NewWorkspaceManagedSQLServerEncryptionProtectorClient creates an instance of the
// WorkspaceManagedSQLServerEncryptionProtectorClient client.
func NewWorkspaceManagedSQLServerEncryptionProtectorClient(subscriptionID string) WorkspaceManagedSQLServerEncryptionProtectorClient {
	return NewWorkspaceManagedSQLServerEncryptionProtectorClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkspaceManagedSQLServerEncryptionProtectorClientWithBaseURI creates an instance of the
// WorkspaceManagedSQLServerEncryptionProtectorClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWorkspaceManagedSQLServerEncryptionProtectorClientWithBaseURI(baseURI string, subscriptionID string) WorkspaceManagedSQLServerEncryptionProtectorClient {
	return WorkspaceManagedSQLServerEncryptionProtectorClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate updates workspace managed sql server's encryption protector.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// parameters - the requested encryption protector resource state.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, parameters EncryptionProtector) (result WorkspaceManagedSQLServerEncryptionProtectorCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerEncryptionProtectorClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, parameters EncryptionProtector) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"encryptionProtectorName": autorest.Encode("path", "current"),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceName":           autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Kind = nil
	parameters.Location = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/encryptionProtector/{encryptionProtectorName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) CreateOrUpdateSender(req *http.Request) (future WorkspaceManagedSQLServerEncryptionProtectorCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) CreateOrUpdateResponder(resp *http.Response) (result EncryptionProtector, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get workspace managed sql server's encryption protector.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) Get(ctx context.Context, resourceGroupName string, workspaceName string) (result EncryptionProtector, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerEncryptionProtectorClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"encryptionProtectorName": autorest.Encode("path", "current"),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceName":           autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/encryptionProtector/{encryptionProtectorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) GetResponder(resp *http.Response) (result EncryptionProtector, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get list of encryption protectors for workspace managed sql server.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) List(ctx context.Context, resourceGroupName string, workspaceName string) (result EncryptionProtectorListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerEncryptionProtectorClient.List")
		defer func() {
			sc := -1
			if result.eplr.Response.Response != nil {
				sc = result.eplr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.eplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "List", resp, "Failure sending request")
		return
	}

	result.eplr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "List", resp, "Failure responding to request")
		return
	}
	if result.eplr.hasNextLink() && result.eplr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) ListPreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/encryptionProtector", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) ListResponder(resp *http.Response) (result EncryptionProtectorListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) listNextResults(ctx context.Context, lastResults EncryptionProtectorListResult) (result EncryptionProtectorListResult, err error) {
	req, err := lastResults.encryptionProtectorListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) ListComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result EncryptionProtectorListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerEncryptionProtectorClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, workspaceName)
	return
}

// Revalidate revalidates workspace managed sql server's existing encryption protector.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) Revalidate(ctx context.Context, resourceGroupName string, workspaceName string) (result WorkspaceManagedSQLServerEncryptionProtectorRevalidateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceManagedSQLServerEncryptionProtectorClient.Revalidate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "Revalidate", err.Error())
	}

	req, err := client.RevalidatePreparer(ctx, resourceGroupName, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "Revalidate", nil, "Failure preparing request")
		return
	}

	result, err = client.RevalidateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.WorkspaceManagedSQLServerEncryptionProtectorClient", "Revalidate", result.Response(), "Failure sending request")
		return
	}

	return
}

// RevalidatePreparer prepares the Revalidate request.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) RevalidatePreparer(ctx context.Context, resourceGroupName string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"encryptionProtectorName": autorest.Encode("path", "current"),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
		"workspaceName":           autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/encryptionProtector/{encryptionProtectorName}/revalidate", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RevalidateSender sends the Revalidate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) RevalidateSender(req *http.Request) (future WorkspaceManagedSQLServerEncryptionProtectorRevalidateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// RevalidateResponder handles the response to the Revalidate request. The method always
// closes the http.Response Body.
func (client WorkspaceManagedSQLServerEncryptionProtectorClient) RevalidateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
