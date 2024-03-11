package botservice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "github.com/Azure/go-autorest/autorest/azure"
    "net/http"
    "context"
    "github.com/Azure/go-autorest/tracing"
)

// OperationResultsClient is the azure Bot Service is a platform for creating smart conversational agents.
type OperationResultsClient struct {
    BaseClient
}
// NewOperationResultsClient creates an instance of the OperationResultsClient client.
func NewOperationResultsClient(subscriptionID string) OperationResultsClient {
    return NewOperationResultsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOperationResultsClientWithBaseURI creates an instance of the OperationResultsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
    func NewOperationResultsClientWithBaseURI(baseURI string, subscriptionID string) OperationResultsClient {
        return OperationResultsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get get the operation result for a long running operation.
    // Parameters:
        // operationResultID - the ID of the operation result to get.
func (client OperationResultsClient) Get(ctx context.Context, operationResultID string) (result OperationResultsGetFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/OperationResultsClient.Get")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, operationResultID)
    if err != nil {
    err = autorest.NewErrorWithError(err, "botservice.OperationResultsClient", "Get", nil , "Failure preparing request")
    return
    }

        result, err = client.GetSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "botservice.OperationResultsClient", "Get", result.Response(), "Failure sending request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client OperationResultsClient) GetPreparer(ctx context.Context, operationResultID string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "operationResultId": autorest.Encode("path",operationResultID),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2021-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.BotService/operationresults/{operationResultId}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client OperationResultsClient) GetSender(req *http.Request) (future OperationResultsGetFuture, err error) {
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

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client OperationResultsClient) GetResponder(resp *http.Response) (result OperationResultsDescription, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

