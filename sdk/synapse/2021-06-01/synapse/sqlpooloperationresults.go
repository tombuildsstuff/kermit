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

// SQLPoolOperationResultsClient is the azure Synapse Analytics Management Client
type SQLPoolOperationResultsClient struct {
	BaseClient
}

// NewSQLPoolOperationResultsClient creates an instance of the SQLPoolOperationResultsClient client.
func NewSQLPoolOperationResultsClient(subscriptionID string) SQLPoolOperationResultsClient {
	return NewSQLPoolOperationResultsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSQLPoolOperationResultsClientWithBaseURI creates an instance of the SQLPoolOperationResultsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewSQLPoolOperationResultsClientWithBaseURI(baseURI string, subscriptionID string) SQLPoolOperationResultsClient {
	return SQLPoolOperationResultsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetLocationHeaderResult get the status of a SQL pool operation
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// SQLPoolName - SQL pool name
// operationID - operation ID
func (client SQLPoolOperationResultsClient) GetLocationHeaderResult(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, operationID string) (result SQLPoolOperationResultsGetLocationHeaderResultFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SQLPoolOperationResultsClient.GetLocationHeaderResult")
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
		return result, validation.NewError("synapse.SQLPoolOperationResultsClient", "GetLocationHeaderResult", err.Error())
	}

	req, err := client.GetLocationHeaderResultPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolOperationResultsClient", "GetLocationHeaderResult", nil, "Failure preparing request")
		return
	}

	result, err = client.GetLocationHeaderResultSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.SQLPoolOperationResultsClient", "GetLocationHeaderResult", result.Response(), "Failure sending request")
		return
	}

	return
}

// GetLocationHeaderResultPreparer prepares the GetLocationHeaderResult request.
func (client SQLPoolOperationResultsClient) GetLocationHeaderResultPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sqlPoolName":       autorest.Encode("path", SQLPoolName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/operationResults/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetLocationHeaderResultSender sends the GetLocationHeaderResult request. The method will close the
// http.Response Body if it receives an error.
func (client SQLPoolOperationResultsClient) GetLocationHeaderResultSender(req *http.Request) (future SQLPoolOperationResultsGetLocationHeaderResultFuture, err error) {
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

// GetLocationHeaderResultResponder handles the response to the GetLocationHeaderResult request. The method always
// closes the http.Response Body.
func (client SQLPoolOperationResultsClient) GetLocationHeaderResultResponder(resp *http.Response) (result SQLPool, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
