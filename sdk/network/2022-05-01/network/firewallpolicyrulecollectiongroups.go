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

// FirewallPolicyRuleCollectionGroupsClient is the network Client
type FirewallPolicyRuleCollectionGroupsClient struct {
    BaseClient
}
// NewFirewallPolicyRuleCollectionGroupsClient creates an instance of the FirewallPolicyRuleCollectionGroupsClient
// client.
func NewFirewallPolicyRuleCollectionGroupsClient(subscriptionID string) FirewallPolicyRuleCollectionGroupsClient {
    return NewFirewallPolicyRuleCollectionGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFirewallPolicyRuleCollectionGroupsClientWithBaseURI creates an instance of the
// FirewallPolicyRuleCollectionGroupsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewFirewallPolicyRuleCollectionGroupsClientWithBaseURI(baseURI string, subscriptionID string) FirewallPolicyRuleCollectionGroupsClient {
        return FirewallPolicyRuleCollectionGroupsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates the specified FirewallPolicyRuleCollectionGroup.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // firewallPolicyName - the name of the Firewall Policy.
        // ruleCollectionGroupName - the name of the FirewallPolicyRuleCollectionGroup.
        // parameters - parameters supplied to the create or update FirewallPolicyRuleCollectionGroup operation.
func (client FirewallPolicyRuleCollectionGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup) (result FirewallPolicyRuleCollectionGroupsCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FirewallPolicyRuleCollectionGroupsClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: parameters,
         Constraints: []validation.Constraint{	{Target: "parameters.FirewallPolicyRuleCollectionGroupProperties", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "parameters.FirewallPolicyRuleCollectionGroupProperties.Priority", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "parameters.FirewallPolicyRuleCollectionGroupProperties.Priority", Name: validation.InclusiveMaximum, Rule: int64(65000), Chain: nil },
        	{Target: "parameters.FirewallPolicyRuleCollectionGroupProperties.Priority", Name: validation.InclusiveMinimum, Rule: int64(100), Chain: nil },
        }},
        }}}}}); err != nil {
        return result, validation.NewError("network.FirewallPolicyRuleCollectionGroupsClient", "CreateOrUpdate", err.Error())
        }

        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleCollectionGroupsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        result, err = client.CreateOrUpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleCollectionGroupsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client FirewallPolicyRuleCollectionGroupsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters FirewallPolicyRuleCollectionGroup) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "firewallPolicyName": autorest.Encode("path",firewallPolicyName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "ruleCollectionGroupName": autorest.Encode("path",ruleCollectionGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-05-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

            parameters.Etag = nil
            parameters.Type = nil
    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client FirewallPolicyRuleCollectionGroupsClient) CreateOrUpdateSender(req *http.Request) (future FirewallPolicyRuleCollectionGroupsCreateOrUpdateFuture, err error) {
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
    func (client FirewallPolicyRuleCollectionGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result FirewallPolicyRuleCollectionGroup, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete deletes the specified FirewallPolicyRuleCollectionGroup.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // firewallPolicyName - the name of the Firewall Policy.
        // ruleCollectionGroupName - the name of the FirewallPolicyRuleCollectionGroup.
func (client FirewallPolicyRuleCollectionGroupsClient) Delete(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string) (result FirewallPolicyRuleCollectionGroupsDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FirewallPolicyRuleCollectionGroupsClient.Delete")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeletePreparer(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleCollectionGroupsClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleCollectionGroupsClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client FirewallPolicyRuleCollectionGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "firewallPolicyName": autorest.Encode("path",firewallPolicyName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "ruleCollectionGroupName": autorest.Encode("path",ruleCollectionGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-05-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client FirewallPolicyRuleCollectionGroupsClient) DeleteSender(req *http.Request) (future FirewallPolicyRuleCollectionGroupsDeleteFuture, err error) {
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
    func (client FirewallPolicyRuleCollectionGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get gets the specified FirewallPolicyRuleCollectionGroup.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // firewallPolicyName - the name of the Firewall Policy.
        // ruleCollectionGroupName - the name of the FirewallPolicyRuleCollectionGroup.
func (client FirewallPolicyRuleCollectionGroupsClient) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string) (result FirewallPolicyRuleCollectionGroup, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FirewallPolicyRuleCollectionGroupsClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, resourceGroupName, firewallPolicyName, ruleCollectionGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleCollectionGroupsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleCollectionGroupsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleCollectionGroupsClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client FirewallPolicyRuleCollectionGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "firewallPolicyName": autorest.Encode("path",firewallPolicyName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "ruleCollectionGroupName": autorest.Encode("path",ruleCollectionGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-05-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups/{ruleCollectionGroupName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client FirewallPolicyRuleCollectionGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client FirewallPolicyRuleCollectionGroupsClient) GetResponder(resp *http.Response) (result FirewallPolicyRuleCollectionGroup, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List lists all FirewallPolicyRuleCollectionGroups in a FirewallPolicy resource.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // firewallPolicyName - the name of the Firewall Policy.
func (client FirewallPolicyRuleCollectionGroupsClient) List(ctx context.Context, resourceGroupName string, firewallPolicyName string) (result FirewallPolicyRuleCollectionGroupListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FirewallPolicyRuleCollectionGroupsClient.List")
        defer func() {
            sc := -1
        if result.fprcglr.Response.Response != nil {
        sc = result.fprcglr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, firewallPolicyName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleCollectionGroupsClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.fprcglr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleCollectionGroupsClient", "List", resp, "Failure sending request")
        return
        }

        result.fprcglr, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleCollectionGroupsClient", "List", resp, "Failure responding to request")
        return
        }
            if result.fprcglr.hasNextLink() && result.fprcglr.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client FirewallPolicyRuleCollectionGroupsClient) ListPreparer(ctx context.Context, resourceGroupName string, firewallPolicyName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "firewallPolicyName": autorest.Encode("path",firewallPolicyName),
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
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/ruleCollectionGroups",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client FirewallPolicyRuleCollectionGroupsClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client FirewallPolicyRuleCollectionGroupsClient) ListResponder(resp *http.Response) (result FirewallPolicyRuleCollectionGroupListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client FirewallPolicyRuleCollectionGroupsClient) listNextResults(ctx context.Context, lastResults FirewallPolicyRuleCollectionGroupListResult) (result FirewallPolicyRuleCollectionGroupListResult, err error) {
            req, err := lastResults.firewallPolicyRuleCollectionGroupListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "network.FirewallPolicyRuleCollectionGroupsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "network.FirewallPolicyRuleCollectionGroupsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.FirewallPolicyRuleCollectionGroupsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client FirewallPolicyRuleCollectionGroupsClient) ListComplete(ctx context.Context, resourceGroupName string, firewallPolicyName string) (result FirewallPolicyRuleCollectionGroupListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/FirewallPolicyRuleCollectionGroupsClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName, firewallPolicyName)
                            return
            }

