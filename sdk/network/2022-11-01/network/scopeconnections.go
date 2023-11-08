package network

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
    "github.com/Azure/go-autorest/autorest/validation"
)

// ScopeConnectionsClient is the network Client
type ScopeConnectionsClient struct {
    BaseClient
}
// NewScopeConnectionsClient creates an instance of the ScopeConnectionsClient client.
func NewScopeConnectionsClient(subscriptionID string) ScopeConnectionsClient {
    return NewScopeConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewScopeConnectionsClientWithBaseURI creates an instance of the ScopeConnectionsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
    func NewScopeConnectionsClientWithBaseURI(baseURI string, subscriptionID string) ScopeConnectionsClient {
        return ScopeConnectionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates scope connection from Network Manager
    // Parameters:
        // parameters - scope connection to be created/updated.
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
        // scopeConnectionName - name for the cross-tenant connection.
func (client ScopeConnectionsClient) CreateOrUpdate(ctx context.Context, parameters ScopeConnection, resourceGroupName string, networkManagerName string, scopeConnectionName string) (result ScopeConnection, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ScopeConnectionsClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.CreateOrUpdatePreparer(ctx, parameters, resourceGroupName, networkManagerName, scopeConnectionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        resp, err := client.CreateOrUpdateSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "CreateOrUpdate", resp, "Failure sending request")
        return
        }

        result, err = client.CreateOrUpdateResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "CreateOrUpdate", resp, "Failure responding to request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client ScopeConnectionsClient) CreateOrUpdatePreparer(ctx context.Context, parameters ScopeConnection, resourceGroupName string, networkManagerName string, scopeConnectionName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "scopeConnectionName": autorest.Encode("path",scopeConnectionName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

            parameters.SystemData = nil
    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/scopeConnections/{scopeConnectionName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ScopeConnectionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
    // closes the http.Response Body.
    func (client ScopeConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result ScopeConnection, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete delete the pending scope connection created by this network manager.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
        // scopeConnectionName - name for the cross-tenant connection.
func (client ScopeConnectionsClient) Delete(ctx context.Context, resourceGroupName string, networkManagerName string, scopeConnectionName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ScopeConnectionsClient.Delete")
        defer func() {
            sc := -1
        if result.Response != nil {
        sc = result.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeletePreparer(ctx, resourceGroupName, networkManagerName, scopeConnectionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "Delete", nil , "Failure preparing request")
    return
    }

        resp, err := client.DeleteSender(req)
        if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "Delete", resp, "Failure sending request")
        return
        }

        result, err = client.DeleteResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "Delete", resp, "Failure responding to request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client ScopeConnectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkManagerName string, scopeConnectionName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "scopeConnectionName": autorest.Encode("path",scopeConnectionName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/scopeConnections/{scopeConnectionName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ScopeConnectionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // DeleteResponder handles the response to the Delete request. The method always
    // closes the http.Response Body.
    func (client ScopeConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get get specified scope connection created by this Network Manager.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
        // scopeConnectionName - name for the cross-tenant connection.
func (client ScopeConnectionsClient) Get(ctx context.Context, resourceGroupName string, networkManagerName string, scopeConnectionName string) (result ScopeConnection, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ScopeConnectionsClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, resourceGroupName, networkManagerName, scopeConnectionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client ScopeConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, networkManagerName string, scopeConnectionName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "scopeConnectionName": autorest.Encode("path",scopeConnectionName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/scopeConnections/{scopeConnectionName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ScopeConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client ScopeConnectionsClient) GetResponder(resp *http.Response) (result ScopeConnection, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List list all scope connections created by this network manager.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
        // top - an optional query parameter which specifies the maximum number of records to be returned by the
        // server.
        // skipToken - skipToken is only used if a previous operation returned a partial result. If a previous response
        // contains a nextLink element, the value of the nextLink element will include a skipToken parameter that
        // specifies a starting point to use for subsequent calls.
func (client ScopeConnectionsClient) List(ctx context.Context, resourceGroupName string, networkManagerName string, top *int32, skipToken string) (result ScopeConnectionListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ScopeConnectionsClient.List")
        defer func() {
            sc := -1
        if result.sclr.Response.Response != nil {
        sc = result.sclr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: top,
         Constraints: []validation.Constraint{	{Target: "top", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(20), Chain: nil },
        	{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil },
        }}}}}); err != nil {
        return result, validation.NewError("network.ScopeConnectionsClient", "List", err.Error())
        }

            result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, networkManagerName, top, skipToken)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.sclr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "List", resp, "Failure sending request")
        return
        }

        result.sclr, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "List", resp, "Failure responding to request")
        return
        }
            if result.sclr.hasNextLink() && result.sclr.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client ScopeConnectionsClient) ListPreparer(ctx context.Context, resourceGroupName string, networkManagerName string, top *int32, skipToken string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
        if top != nil {
        queryParameters["$top"] = autorest.Encode("query",*top)
        }
        if len(skipToken) > 0 {
        queryParameters["$skipToken"] = autorest.Encode("query",skipToken)
        }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/scopeConnections",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client ScopeConnectionsClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client ScopeConnectionsClient) ListResponder(resp *http.Response) (result ScopeConnectionListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client ScopeConnectionsClient) listNextResults(ctx context.Context, lastResults ScopeConnectionListResult) (result ScopeConnectionListResult, err error) {
            req, err := lastResults.scopeConnectionListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.ScopeConnectionsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client ScopeConnectionsClient) ListComplete(ctx context.Context, resourceGroupName string, networkManagerName string, top *int32, skipToken string) (result ScopeConnectionListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ScopeConnectionsClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName, networkManagerName, top, skipToken)
                            return
            }

