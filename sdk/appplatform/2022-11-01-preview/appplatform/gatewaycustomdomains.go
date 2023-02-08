package appplatform

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

// GatewayCustomDomainsClient is the REST API for Azure Spring Apps
type GatewayCustomDomainsClient struct {
    BaseClient
}
// NewGatewayCustomDomainsClient creates an instance of the GatewayCustomDomainsClient client.
func NewGatewayCustomDomainsClient(subscriptionID string) GatewayCustomDomainsClient {
    return NewGatewayCustomDomainsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGatewayCustomDomainsClientWithBaseURI creates an instance of the GatewayCustomDomainsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
    func NewGatewayCustomDomainsClientWithBaseURI(baseURI string, subscriptionID string) GatewayCustomDomainsClient {
        return GatewayCustomDomainsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate create or update the Spring Cloud Gateway custom domain.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // gatewayName - the name of Spring Cloud Gateway.
        // domainName - the name of the Spring Cloud Gateway custom domain.
        // gatewayCustomDomainResource - the gateway custom domain resource for the create or update operation
func (client GatewayCustomDomainsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string, gatewayCustomDomainResource GatewayCustomDomainResource) (result GatewayCustomDomainsCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/GatewayCustomDomainsClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, gatewayName, domainName, gatewayCustomDomainResource)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.GatewayCustomDomainsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        result, err = client.CreateOrUpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.GatewayCustomDomainsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client GatewayCustomDomainsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string, gatewayCustomDomainResource GatewayCustomDomainResource) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "domainName": autorest.Encode("path",domainName),
        "gatewayName": autorest.Encode("path",gatewayName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/domains/{domainName}",pathParameters),
autorest.WithJSON(gatewayCustomDomainResource),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client GatewayCustomDomainsClient) CreateOrUpdateSender(req *http.Request) (future GatewayCustomDomainsCreateOrUpdateFuture, err error) {
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
    func (client GatewayCustomDomainsClient) CreateOrUpdateResponder(resp *http.Response) (result GatewayCustomDomainResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete delete the Spring Cloud Gateway custom domain.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // gatewayName - the name of Spring Cloud Gateway.
        // domainName - the name of the Spring Cloud Gateway custom domain.
func (client GatewayCustomDomainsClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string) (result GatewayCustomDomainsDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/GatewayCustomDomainsClient.Delete")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, gatewayName, domainName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.GatewayCustomDomainsClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.GatewayCustomDomainsClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client GatewayCustomDomainsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "domainName": autorest.Encode("path",domainName),
        "gatewayName": autorest.Encode("path",gatewayName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/domains/{domainName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client GatewayCustomDomainsClient) DeleteSender(req *http.Request) (future GatewayCustomDomainsDeleteFuture, err error) {
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
    func (client GatewayCustomDomainsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get get the Spring Cloud Gateway custom domain.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // gatewayName - the name of Spring Cloud Gateway.
        // domainName - the name of the Spring Cloud Gateway custom domain.
func (client GatewayCustomDomainsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string) (result GatewayCustomDomainResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/GatewayCustomDomainsClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, gatewayName, domainName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.GatewayCustomDomainsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "appplatform.GatewayCustomDomainsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.GatewayCustomDomainsClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client GatewayCustomDomainsClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "domainName": autorest.Encode("path",domainName),
        "gatewayName": autorest.Encode("path",gatewayName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/domains/{domainName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client GatewayCustomDomainsClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client GatewayCustomDomainsClient) GetResponder(resp *http.Response) (result GatewayCustomDomainResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List handle requests to list all Spring Cloud Gateway custom domains.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // gatewayName - the name of Spring Cloud Gateway.
func (client GatewayCustomDomainsClient) List(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string) (result GatewayCustomDomainResourceCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/GatewayCustomDomainsClient.List")
        defer func() {
            sc := -1
        if result.gcdrc.Response.Response != nil {
        sc = result.gcdrc.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, serviceName, gatewayName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.GatewayCustomDomainsClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.gcdrc.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "appplatform.GatewayCustomDomainsClient", "List", resp, "Failure sending request")
        return
        }

        result.gcdrc, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.GatewayCustomDomainsClient", "List", resp, "Failure responding to request")
        return
        }
            if result.gcdrc.hasNextLink() && result.gcdrc.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client GatewayCustomDomainsClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "gatewayName": autorest.Encode("path",gatewayName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/gateways/{gatewayName}/domains",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client GatewayCustomDomainsClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client GatewayCustomDomainsClient) ListResponder(resp *http.Response) (result GatewayCustomDomainResourceCollection, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client GatewayCustomDomainsClient) listNextResults(ctx context.Context, lastResults GatewayCustomDomainResourceCollection) (result GatewayCustomDomainResourceCollection, err error) {
            req, err := lastResults.gatewayCustomDomainResourceCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "appplatform.GatewayCustomDomainsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "appplatform.GatewayCustomDomainsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "appplatform.GatewayCustomDomainsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client GatewayCustomDomainsClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string) (result GatewayCustomDomainResourceCollectionIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/GatewayCustomDomainsClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName, serviceName, gatewayName)
                            return
            }

