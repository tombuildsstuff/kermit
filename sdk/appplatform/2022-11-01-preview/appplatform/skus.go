package appplatform

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
	"github.com/Azure/go-autorest/tracing"
)

// SkusClient is the REST API for Azure Spring Apps
type SkusClient struct {
	BaseClient
}

// NewSkusClient creates an instance of the SkusClient client.
func NewSkusClient(subscriptionID string) SkusClient {
	return NewSkusClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSkusClientWithBaseURI creates an instance of the SkusClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSkusClientWithBaseURI(baseURI string, subscriptionID string) SkusClient {
	return SkusClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all of the available skus of the Microsoft.AppPlatform provider.
func (client SkusClient) List(ctx context.Context) (result ResourceSkuCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SkusClient.List")
		defer func() {
			sc := -1
			if result.rsc.Response.Response != nil {
				sc = result.rsc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.SkusClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rsc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "appplatform.SkusClient", "List", resp, "Failure sending request")
		return
	}

	result.rsc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.SkusClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rsc.hasNextLink() && result.rsc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client SkusClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.AppPlatform/skus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SkusClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SkusClient) ListResponder(resp *http.Response) (result ResourceSkuCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SkusClient) listNextResults(ctx context.Context, lastResults ResourceSkuCollection) (result ResourceSkuCollection, err error) {
	req, err := lastResults.resourceSkuCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "appplatform.SkusClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "appplatform.SkusClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.SkusClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SkusClient) ListComplete(ctx context.Context) (result ResourceSkuCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SkusClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}
