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

// AdminRulesClient is the network Client
type AdminRulesClient struct {
    BaseClient
}
// NewAdminRulesClient creates an instance of the AdminRulesClient client.
func NewAdminRulesClient(subscriptionID string) AdminRulesClient {
    return NewAdminRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAdminRulesClientWithBaseURI creates an instance of the AdminRulesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewAdminRulesClientWithBaseURI(baseURI string, subscriptionID string) AdminRulesClient {
        return AdminRulesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates an admin rule.
    // Parameters:
        // adminRule - the admin rule to create or update
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
        // configurationName - the name of the network manager Security Configuration.
        // ruleCollectionName - the name of the network manager security Configuration rule collection.
        // ruleName - the name of the rule.
func (client AdminRulesClient) CreateOrUpdate(ctx context.Context, adminRule BasicBaseAdminRule, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string) (result BaseAdminRuleModel, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AdminRulesClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.CreateOrUpdatePreparer(ctx, adminRule, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, ruleName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.AdminRulesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        resp, err := client.CreateOrUpdateSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.AdminRulesClient", "CreateOrUpdate", resp, "Failure sending request")
        return
        }

        result, err = client.CreateOrUpdateResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.AdminRulesClient", "CreateOrUpdate", resp, "Failure responding to request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client AdminRulesClient) CreateOrUpdatePreparer(ctx context.Context, adminRule BasicBaseAdminRule, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationName": autorest.Encode("path",configurationName),
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "ruleCollectionName": autorest.Encode("path",ruleCollectionName),
        "ruleName": autorest.Encode("path",ruleName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations/{configurationName}/ruleCollections/{ruleCollectionName}/rules/{ruleName}",pathParameters),
autorest.WithJSON(adminRule),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client AdminRulesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
    // closes the http.Response Body.
    func (client AdminRulesClient) CreateOrUpdateResponder(resp *http.Response) (result BaseAdminRuleModel, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete deletes an admin rule.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
        // configurationName - the name of the network manager Security Configuration.
        // ruleCollectionName - the name of the network manager security Configuration rule collection.
        // ruleName - the name of the rule.
        // force - deletes the resource even if it is part of a deployed configuration. If the configuration has been
        // deployed, the service will do a cleanup deployment in the background, prior to the delete.
func (client AdminRulesClient) Delete(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, force *bool) (result AdminRulesDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AdminRulesClient.Delete")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeletePreparer(ctx, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, ruleName, force)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.AdminRulesClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.AdminRulesClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client AdminRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, force *bool) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationName": autorest.Encode("path",configurationName),
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "ruleCollectionName": autorest.Encode("path",ruleCollectionName),
        "ruleName": autorest.Encode("path",ruleName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
        if force != nil {
        queryParameters["force"] = autorest.Encode("query",*force)
        }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations/{configurationName}/ruleCollections/{ruleCollectionName}/rules/{ruleName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client AdminRulesClient) DeleteSender(req *http.Request) (future AdminRulesDeleteFuture, err error) {
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
    func (client AdminRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get gets a network manager security configuration admin rule.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
        // configurationName - the name of the network manager Security Configuration.
        // ruleCollectionName - the name of the network manager security Configuration rule collection.
        // ruleName - the name of the rule.
func (client AdminRulesClient) Get(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string) (result BaseAdminRuleModel, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AdminRulesClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, ruleName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.AdminRulesClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.AdminRulesClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.AdminRulesClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client AdminRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationName": autorest.Encode("path",configurationName),
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "ruleCollectionName": autorest.Encode("path",ruleCollectionName),
        "ruleName": autorest.Encode("path",ruleName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations/{configurationName}/ruleCollections/{ruleCollectionName}/rules/{ruleName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client AdminRulesClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client AdminRulesClient) GetResponder(resp *http.Response) (result BaseAdminRuleModel, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List list all network manager security configuration admin rules.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
        // configurationName - the name of the network manager Security Configuration.
        // ruleCollectionName - the name of the network manager security Configuration rule collection.
        // top - an optional query parameter which specifies the maximum number of records to be returned by the
        // server.
        // skipToken - skipToken is only used if a previous operation returned a partial result. If a previous response
        // contains a nextLink element, the value of the nextLink element will include a skipToken parameter that
        // specifies a starting point to use for subsequent calls.
func (client AdminRulesClient) List(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, top *int32, skipToken string) (result AdminRuleListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/AdminRulesClient.List")
        defer func() {
            sc := -1
        if result.arlr.Response.Response != nil {
        sc = result.arlr.Response.Response.StatusCode
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
        return result, validation.NewError("network.AdminRulesClient", "List", err.Error())
        }

            result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, top, skipToken)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.AdminRulesClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.arlr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.AdminRulesClient", "List", resp, "Failure sending request")
        return
        }

        result.arlr, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.AdminRulesClient", "List", resp, "Failure responding to request")
        return
        }
            if result.arlr.hasNextLink() && result.arlr.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client AdminRulesClient) ListPreparer(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, top *int32, skipToken string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationName": autorest.Encode("path",configurationName),
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "ruleCollectionName": autorest.Encode("path",ruleCollectionName),
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
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations/{configurationName}/ruleCollections/{ruleCollectionName}/rules",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client AdminRulesClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client AdminRulesClient) ListResponder(resp *http.Response) (result AdminRuleListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client AdminRulesClient) listNextResults(ctx context.Context, lastResults AdminRuleListResult) (result AdminRuleListResult, err error) {
            req, err := lastResults.adminRuleListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "network.AdminRulesClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "network.AdminRulesClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.AdminRulesClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client AdminRulesClient) ListComplete(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, top *int32, skipToken string) (result AdminRuleListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/AdminRulesClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName, networkManagerName, configurationName, ruleCollectionName, top, skipToken)
                            return
            }

