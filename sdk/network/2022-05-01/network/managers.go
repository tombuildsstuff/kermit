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

// ManagersClient is the network Client
type ManagersClient struct {
    BaseClient
}
// NewManagersClient creates an instance of the ManagersClient client.
func NewManagersClient(subscriptionID string) ManagersClient {
    return NewManagersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagersClientWithBaseURI creates an instance of the ManagersClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewManagersClientWithBaseURI(baseURI string, subscriptionID string) ManagersClient {
        return ManagersClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates a Network Manager.
    // Parameters:
        // parameters - parameters supplied to specify which network manager is.
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
func (client ManagersClient) CreateOrUpdate(ctx context.Context, parameters Manager, resourceGroupName string, networkManagerName string) (result Manager, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ManagersClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: parameters,
         Constraints: []validation.Constraint{	{Target: "parameters.ManagerProperties", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "parameters.ManagerProperties.NetworkManagerScopes", Name: validation.Null, Rule: true, Chain: nil },
        	{Target: "parameters.ManagerProperties.NetworkManagerScopeAccesses", Name: validation.Null, Rule: true, Chain: nil },
        }}}}}); err != nil {
        return result, validation.NewError("network.ManagersClient", "CreateOrUpdate", err.Error())
        }

        req, err := client.CreateOrUpdatePreparer(ctx, parameters, resourceGroupName, networkManagerName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ManagersClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        resp, err := client.CreateOrUpdateSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.ManagersClient", "CreateOrUpdate", resp, "Failure sending request")
        return
        }

        result, err = client.CreateOrUpdateResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.ManagersClient", "CreateOrUpdate", resp, "Failure responding to request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client ManagersClient) CreateOrUpdatePreparer(ctx context.Context, parameters Manager, resourceGroupName string, networkManagerName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-05-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

            parameters.Etag = nil
            parameters.SystemData = nil
    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ManagersClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
    // closes the http.Response Body.
    func (client ManagersClient) CreateOrUpdateResponder(resp *http.Response) (result Manager, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete deletes a network manager.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
        // force - deletes the resource even if it is part of a deployed configuration. If the configuration has been
        // deployed, the service will do a cleanup deployment in the background, prior to the delete.
func (client ManagersClient) Delete(ctx context.Context, resourceGroupName string, networkManagerName string, force *bool) (result ManagersDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ManagersClient.Delete")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeletePreparer(ctx, resourceGroupName, networkManagerName, force)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ManagersClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.ManagersClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client ManagersClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkManagerName string, force *bool) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-05-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
        if force != nil {
        queryParameters["force"] = autorest.Encode("query",*force)
        }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ManagersClient) DeleteSender(req *http.Request) (future ManagersDeleteFuture, err error) {
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

    // DeleteResponder handles the response to the Delete request. The method always
    // closes the http.Response Body.
    func (client ManagersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get gets the specified Network Manager.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
func (client ManagersClient) Get(ctx context.Context, resourceGroupName string, networkManagerName string) (result Manager, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ManagersClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, resourceGroupName, networkManagerName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ManagersClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.ManagersClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.ManagersClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client ManagersClient) GetPreparer(ctx context.Context, resourceGroupName string, networkManagerName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-05-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ManagersClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client ManagersClient) GetResponder(resp *http.Response) (result Manager, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List list network managers in a resource group.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // top - an optional query parameter which specifies the maximum number of records to be returned by the
        // server.
        // skipToken - skipToken is only used if a previous operation returned a partial result. If a previous response
        // contains a nextLink element, the value of the nextLink element will include a skipToken parameter that
        // specifies a starting point to use for subsequent calls.
func (client ManagersClient) List(ctx context.Context, resourceGroupName string, top *int32, skipToken string) (result ManagerListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ManagersClient.List")
        defer func() {
            sc := -1
        if result.mlr.Response.Response != nil {
        sc = result.mlr.Response.Response.StatusCode
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
        return result, validation.NewError("network.ManagersClient", "List", err.Error())
        }

            result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, top, skipToken)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ManagersClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.mlr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.ManagersClient", "List", resp, "Failure sending request")
        return
        }

        result.mlr, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.ManagersClient", "List", resp, "Failure responding to request")
        return
        }
            if result.mlr.hasNextLink() && result.mlr.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client ManagersClient) ListPreparer(ctx context.Context, resourceGroupName string, top *int32, skipToken string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-05-01"
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
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client ManagersClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client ManagersClient) ListResponder(resp *http.Response) (result ManagerListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client ManagersClient) listNextResults(ctx context.Context, lastResults ManagerListResult) (result ManagerListResult, err error) {
            req, err := lastResults.managerListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "network.ManagersClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "network.ManagersClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.ManagersClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client ManagersClient) ListComplete(ctx context.Context, resourceGroupName string, top *int32, skipToken string) (result ManagerListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ManagersClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName, top, skipToken)
                            return
            }

// ListBySubscription list all network managers in a subscription.
    // Parameters:
        // top - an optional query parameter which specifies the maximum number of records to be returned by the
        // server.
        // skipToken - skipToken is only used if a previous operation returned a partial result. If a previous response
        // contains a nextLink element, the value of the nextLink element will include a skipToken parameter that
        // specifies a starting point to use for subsequent calls.
func (client ManagersClient) ListBySubscription(ctx context.Context, top *int32, skipToken string) (result ManagerListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ManagersClient.ListBySubscription")
        defer func() {
            sc := -1
        if result.mlr.Response.Response != nil {
        sc = result.mlr.Response.Response.StatusCode
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
        return result, validation.NewError("network.ManagersClient", "ListBySubscription", err.Error())
        }

            result.fn = client.listBySubscriptionNextResults
    req, err := client.ListBySubscriptionPreparer(ctx, top, skipToken)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ManagersClient", "ListBySubscription", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListBySubscriptionSender(req)
        if err != nil {
        result.mlr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.ManagersClient", "ListBySubscription", resp, "Failure sending request")
        return
        }

        result.mlr, err = client.ListBySubscriptionResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.ManagersClient", "ListBySubscription", resp, "Failure responding to request")
        return
        }
            if result.mlr.hasNextLink() && result.mlr.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListBySubscriptionPreparer prepares the ListBySubscription request.
    func (client ManagersClient) ListBySubscriptionPreparer(ctx context.Context, top *int32, skipToken string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-05-01"
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
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkManagers",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListBySubscriptionSender sends the ListBySubscription request. The method will close the
    // http.Response Body if it receives an error.
    func (client ManagersClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
    // closes the http.Response Body.
    func (client ManagersClient) ListBySubscriptionResponder(resp *http.Response) (result ManagerListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listBySubscriptionNextResults retrieves the next set of results, if any.
            func (client ManagersClient) listBySubscriptionNextResults(ctx context.Context, lastResults ManagerListResult) (result ManagerListResult, err error) {
            req, err := lastResults.managerListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "network.ManagersClient", "listBySubscriptionNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListBySubscriptionSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "network.ManagersClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListBySubscriptionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.ManagersClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
            func (client ManagersClient) ListBySubscriptionComplete(ctx context.Context, top *int32, skipToken string) (result ManagerListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ManagersClient.ListBySubscription")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.ListBySubscription(ctx, top, skipToken)
                            return
            }

// Patch patch NetworkManager.
    // Parameters:
        // parameters - parameters supplied to specify which network manager is.
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
func (client ManagersClient) Patch(ctx context.Context, parameters PatchObject, resourceGroupName string, networkManagerName string) (result Manager, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ManagersClient.Patch")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.PatchPreparer(ctx, parameters, resourceGroupName, networkManagerName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.ManagersClient", "Patch", nil , "Failure preparing request")
    return
    }

        resp, err := client.PatchSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.ManagersClient", "Patch", resp, "Failure sending request")
        return
        }

        result, err = client.PatchResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.ManagersClient", "Patch", resp, "Failure responding to request")
        return
        }

    return
}

    // PatchPreparer prepares the Patch request.
    func (client ManagersClient) PatchPreparer(ctx context.Context, parameters PatchObject, resourceGroupName string, networkManagerName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-05-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPatch(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // PatchSender sends the Patch request. The method will close the
    // http.Response Body if it receives an error.
    func (client ManagersClient) PatchSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // PatchResponder handles the response to the Patch request. The method always
    // closes the http.Response Body.
    func (client ManagersClient) PatchResponder(resp *http.Response) (result Manager, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

