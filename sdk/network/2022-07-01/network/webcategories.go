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
)

// WebCategoriesClient is the network Client
type WebCategoriesClient struct {
    BaseClient
}
// NewWebCategoriesClient creates an instance of the WebCategoriesClient client.
func NewWebCategoriesClient(subscriptionID string) WebCategoriesClient {
    return NewWebCategoriesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWebCategoriesClientWithBaseURI creates an instance of the WebCategoriesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
    func NewWebCategoriesClientWithBaseURI(baseURI string, subscriptionID string) WebCategoriesClient {
        return WebCategoriesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get gets the specified Azure Web Category.
    // Parameters:
        // name - the name of the azureWebCategory.
        // expand - expands resourceIds back referenced by the azureWebCategory resource.
func (client WebCategoriesClient) Get(ctx context.Context, name string, expand string) (result AzureWebCategory, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WebCategoriesClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, name, expand)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.WebCategoriesClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.WebCategoriesClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.WebCategoriesClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client WebCategoriesClient) GetPreparer(ctx context.Context, name string, expand string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "name": autorest.Encode("path",name),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-07-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
        if len(expand) > 0 {
        queryParameters["$expand"] = autorest.Encode("query",expand)
        }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureWebCategories/{name}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client WebCategoriesClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client WebCategoriesClient) GetResponder(resp *http.Response) (result AzureWebCategory, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// ListBySubscription gets all the Azure Web Categories in a subscription.
func (client WebCategoriesClient) ListBySubscription(ctx context.Context) (result AzureWebCategoryListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/WebCategoriesClient.ListBySubscription")
        defer func() {
            sc := -1
        if result.awclr.Response.Response != nil {
        sc = result.awclr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.listBySubscriptionNextResults
    req, err := client.ListBySubscriptionPreparer(ctx)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.WebCategoriesClient", "ListBySubscription", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListBySubscriptionSender(req)
        if err != nil {
        result.awclr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.WebCategoriesClient", "ListBySubscription", resp, "Failure sending request")
        return
        }

        result.awclr, err = client.ListBySubscriptionResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.WebCategoriesClient", "ListBySubscription", resp, "Failure responding to request")
        return
        }
            if result.awclr.hasNextLink() && result.awclr.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListBySubscriptionPreparer prepares the ListBySubscription request.
    func (client WebCategoriesClient) ListBySubscriptionPreparer(ctx context.Context) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-07-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureWebCategories",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListBySubscriptionSender sends the ListBySubscription request. The method will close the
    // http.Response Body if it receives an error.
    func (client WebCategoriesClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
    // closes the http.Response Body.
    func (client WebCategoriesClient) ListBySubscriptionResponder(resp *http.Response) (result AzureWebCategoryListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listBySubscriptionNextResults retrieves the next set of results, if any.
            func (client WebCategoriesClient) listBySubscriptionNextResults(ctx context.Context, lastResults AzureWebCategoryListResult) (result AzureWebCategoryListResult, err error) {
            req, err := lastResults.azureWebCategoryListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "network.WebCategoriesClient", "listBySubscriptionNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListBySubscriptionSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "network.WebCategoriesClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListBySubscriptionResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "network.WebCategoriesClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
            func (client WebCategoriesClient) ListBySubscriptionComplete(ctx context.Context) (result AzureWebCategoryListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/WebCategoriesClient.ListBySubscription")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.ListBySubscription(ctx)
                            return
            }

