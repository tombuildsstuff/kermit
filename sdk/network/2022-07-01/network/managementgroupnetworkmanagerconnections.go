package network

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

// ManagementGroupNetworkManagerConnectionsClient is the network Client
type ManagementGroupNetworkManagerConnectionsClient struct {
	BaseClient
}

// NewManagementGroupNetworkManagerConnectionsClient creates an instance of the
// ManagementGroupNetworkManagerConnectionsClient client.
func NewManagementGroupNetworkManagerConnectionsClient(subscriptionID string) ManagementGroupNetworkManagerConnectionsClient {
	return NewManagementGroupNetworkManagerConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagementGroupNetworkManagerConnectionsClientWithBaseURI creates an instance of the
// ManagementGroupNetworkManagerConnectionsClient client using a custom endpoint.  Use this when interacting with an
// Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewManagementGroupNetworkManagerConnectionsClientWithBaseURI(baseURI string, subscriptionID string) ManagementGroupNetworkManagerConnectionsClient {
	return ManagementGroupNetworkManagerConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create a network manager connection on this management group.
// Parameters:
// parameters - network manager connection to be created/updated.
// managementGroupID - the management group Id which uniquely identify the Microsoft Azure management group.
// networkManagerConnectionName - name for the network manager connection.
func (client ManagementGroupNetworkManagerConnectionsClient) CreateOrUpdate(ctx context.Context, parameters ManagerConnection, managementGroupID string, networkManagerConnectionName string) (result ManagerConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagementGroupNetworkManagerConnectionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, parameters, managementGroupID, networkManagerConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ManagementGroupNetworkManagerConnectionsClient) CreateOrUpdatePreparer(ctx context.Context, parameters ManagerConnection, managementGroupID string, networkManagerConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":            autorest.Encode("path", managementGroupID),
		"networkManagerConnectionName": autorest.Encode("path", networkManagerConnectionName),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Network/networkManagerConnections/{networkManagerConnectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementGroupNetworkManagerConnectionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ManagementGroupNetworkManagerConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result ManagerConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete specified pending connection created by this management group.
// Parameters:
// managementGroupID - the management group Id which uniquely identify the Microsoft Azure management group.
// networkManagerConnectionName - name for the network manager connection.
func (client ManagementGroupNetworkManagerConnectionsClient) Delete(ctx context.Context, managementGroupID string, networkManagerConnectionName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagementGroupNetworkManagerConnectionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, managementGroupID, networkManagerConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ManagementGroupNetworkManagerConnectionsClient) DeletePreparer(ctx context.Context, managementGroupID string, networkManagerConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":            autorest.Encode("path", managementGroupID),
		"networkManagerConnectionName": autorest.Encode("path", networkManagerConnectionName),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Network/networkManagerConnections/{networkManagerConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementGroupNetworkManagerConnectionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ManagementGroupNetworkManagerConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a specified connection created by this management group.
// Parameters:
// managementGroupID - the management group Id which uniquely identify the Microsoft Azure management group.
// networkManagerConnectionName - name for the network manager connection.
func (client ManagementGroupNetworkManagerConnectionsClient) Get(ctx context.Context, managementGroupID string, networkManagerConnectionName string) (result ManagerConnection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagementGroupNetworkManagerConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, managementGroupID, networkManagerConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagementGroupNetworkManagerConnectionsClient) GetPreparer(ctx context.Context, managementGroupID string, networkManagerConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":            autorest.Encode("path", managementGroupID),
		"networkManagerConnectionName": autorest.Encode("path", networkManagerConnectionName),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Network/networkManagerConnections/{networkManagerConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementGroupNetworkManagerConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagementGroupNetworkManagerConnectionsClient) GetResponder(resp *http.Response) (result ManagerConnection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all network manager connections created by this management group.
// Parameters:
// managementGroupID - the management group Id which uniquely identify the Microsoft Azure management group.
// top - an optional query parameter which specifies the maximum number of records to be returned by the
// server.
// skipToken - skipToken is only used if a previous operation returned a partial result. If a previous response
// contains a nextLink element, the value of the nextLink element will include a skipToken parameter that
// specifies a starting point to use for subsequent calls.
func (client ManagementGroupNetworkManagerConnectionsClient) List(ctx context.Context, managementGroupID string, top *int32, skipToken string) (result ManagerConnectionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagementGroupNetworkManagerConnectionsClient.List")
		defer func() {
			sc := -1
			if result.mclr.Response.Response != nil {
				sc = result.mclr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(20), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("network.ManagementGroupNetworkManagerConnectionsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, managementGroupID, top, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.mclr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "List", resp, "Failure sending request")
		return
	}

	result.mclr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.mclr.hasNextLink() && result.mclr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ManagementGroupNetworkManagerConnectionsClient) ListPreparer(ctx context.Context, managementGroupID string, top *int32, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId": autorest.Encode("path", managementGroupID),
	}

	const APIVersion = "2022-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Network/networkManagerConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ManagementGroupNetworkManagerConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ManagementGroupNetworkManagerConnectionsClient) ListResponder(resp *http.Response) (result ManagerConnectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ManagementGroupNetworkManagerConnectionsClient) listNextResults(ctx context.Context, lastResults ManagerConnectionListResult) (result ManagerConnectionListResult, err error) {
	req, err := lastResults.managerConnectionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ManagementGroupNetworkManagerConnectionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ManagementGroupNetworkManagerConnectionsClient) ListComplete(ctx context.Context, managementGroupID string, top *int32, skipToken string) (result ManagerConnectionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagementGroupNetworkManagerConnectionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, managementGroupID, top, skipToken)
	return
}
