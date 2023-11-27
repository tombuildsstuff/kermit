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

// SecurityAdminConfigurationsClient is the network Client
type SecurityAdminConfigurationsClient struct {
    BaseClient
}
// NewSecurityAdminConfigurationsClient creates an instance of the SecurityAdminConfigurationsClient client.
func NewSecurityAdminConfigurationsClient(subscriptionID string) SecurityAdminConfigurationsClient {
    return NewSecurityAdminConfigurationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSecurityAdminConfigurationsClientWithBaseURI creates an instance of the SecurityAdminConfigurationsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
    func NewSecurityAdminConfigurationsClientWithBaseURI(baseURI string, subscriptionID string) SecurityAdminConfigurationsClient {
        return SecurityAdminConfigurationsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates or updates a network manager security admin configuration.
    // Parameters:
        // securityAdminConfiguration - the security admin configuration to create or update
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
        // configurationName - the name of the network manager Security Configuration.
func (client SecurityAdminConfigurationsClient) CreateOrUpdate(ctx context.Context, securityAdminConfiguration SecurityAdminConfiguration, resourceGroupName string, networkManagerName string, configurationName string) (result SecurityAdminConfiguration, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SecurityAdminConfigurationsClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.CreateOrUpdatePreparer(ctx, securityAdminConfiguration, resourceGroupName, networkManagerName, configurationName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.SecurityAdminConfigurationsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        resp, err := client.CreateOrUpdateSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.SecurityAdminConfigurationsClient", "CreateOrUpdate", resp, "Failure sending request")
        return
        }

        result, err = client.CreateOrUpdateResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.SecurityAdminConfigurationsClient", "CreateOrUpdate", resp, "Failure responding to request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client SecurityAdminConfigurationsClient) CreateOrUpdatePreparer(ctx context.Context, securityAdminConfiguration SecurityAdminConfiguration, resourceGroupName string, networkManagerName string, configurationName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationName": autorest.Encode("path",configurationName),
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

            securityAdminConfiguration.SystemData = nil
    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations/{configurationName}",pathParameters),
autorest.WithJSON(securityAdminConfiguration),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client SecurityAdminConfigurationsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
    // closes the http.Response Body.
    func (client SecurityAdminConfigurationsClient) CreateOrUpdateResponder(resp *http.Response) (result SecurityAdminConfiguration, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete deletes a network manager security admin configuration.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
        // configurationName - the name of the network manager Security Configuration.
        // force - deletes the resource even if it is part of a deployed configuration. If the configuration has been
        // deployed, the service will do a cleanup deployment in the background, prior to the delete.
func (client SecurityAdminConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, force *bool) (result SecurityAdminConfigurationsDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SecurityAdminConfigurationsClient.Delete")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeletePreparer(ctx, resourceGroupName, networkManagerName, configurationName, force)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.SecurityAdminConfigurationsClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.SecurityAdminConfigurationsClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client SecurityAdminConfigurationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, force *bool) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationName": autorest.Encode("path",configurationName),
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
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
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations/{configurationName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client SecurityAdminConfigurationsClient) DeleteSender(req *http.Request) (future SecurityAdminConfigurationsDeleteFuture, err error) {
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
    func (client SecurityAdminConfigurationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get retrieves a network manager security admin configuration.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
        // configurationName - the name of the network manager Security Configuration.
func (client SecurityAdminConfigurationsClient) Get(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string) (result SecurityAdminConfiguration, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SecurityAdminConfigurationsClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, resourceGroupName, networkManagerName, configurationName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.SecurityAdminConfigurationsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.SecurityAdminConfigurationsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.SecurityAdminConfigurationsClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client SecurityAdminConfigurationsClient) GetPreparer(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationName": autorest.Encode("path",configurationName),
        "networkManagerName": autorest.Encode("path",networkManagerName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations/{configurationName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client SecurityAdminConfigurationsClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client SecurityAdminConfigurationsClient) GetResponder(resp *http.Response) (result SecurityAdminConfiguration, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List lists all the network manager security admin configurations in a network manager, in a paginated format.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // networkManagerName - the name of the network manager.
        // top - an optional query parameter which specifies the maximum number of records to be returned by the
        // server.
        // skipToken - skipToken is only used if a previous operation returned a partial result. If a previous response
        // contains a nextLink element, the value of the nextLink element will include a skipToken parameter that
        // specifies a starting point to use for subsequent calls.
func (client SecurityAdminConfigurationsClient) List(ctx context.Context, resourceGroupName string, networkManagerName string, top *int32, skipToken string) (result SecurityAdminConfigurationListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SecurityAdminConfigurationsClient.List")
        defer func() {
            sc := -1
        if result.saclr.Response.Response != nil {
        sc = result.saclr.Response.Response.StatusCode
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
        return result, validation.NewError("network.SecurityAdminConfigurationsClient", "List", err.Error())
        }

            result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, networkManagerName, top, skipToken)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.SecurityAdminConfigurationsClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.saclr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.SecurityAdminConfigurationsClient", "List", resp, "Failure sending request")
        return
        }

        result.saclr, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.SecurityAdminConfigurationsClient", "List", resp, "Failure responding to request")
        return
        }
            if result.saclr.hasNextLink() && result.saclr.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client SecurityAdminConfigurationsClient) ListPreparer(ctx context.Context, resourceGroupName string, networkManagerName string, top *int32, skipToken string) (*http.Request, error) {
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
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/securityAdminConfigurations",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client SecurityAdminConfigurationsClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client SecurityAdminConfigurationsClient) ListResponder(resp *http.Response) (result SecurityAdminConfigurationListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client SecurityAdminConfigurationsClient) listNextResults(ctx context.Context, lastResults SecurityAdminConfigurationListResult) (result SecurityAdminConfigurationListResult, err error) {
            req, err := lastResults.securityAdminConfigurationListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "network.SecurityAdminConfigurationsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "network.SecurityAdminConfigurationsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.SecurityAdminConfigurationsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client SecurityAdminConfigurationsClient) ListComplete(ctx context.Context, resourceGroupName string, networkManagerName string, top *int32, skipToken string) (result SecurityAdminConfigurationListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/SecurityAdminConfigurationsClient.List")
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

