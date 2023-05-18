package web

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

// DeletedWebAppsClient is the webSite Management Client
type DeletedWebAppsClient struct {
	BaseClient
}

// NewDeletedWebAppsClient creates an instance of the DeletedWebAppsClient client.
func NewDeletedWebAppsClient(subscriptionID string) DeletedWebAppsClient {
	return NewDeletedWebAppsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDeletedWebAppsClientWithBaseURI creates an instance of the DeletedWebAppsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDeletedWebAppsClientWithBaseURI(baseURI string, subscriptionID string) DeletedWebAppsClient {
	return DeletedWebAppsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetDeletedWebAppByLocation description for Get deleted app for a subscription at location.
// Parameters:
// deletedSiteID - the numeric ID of the deleted app, e.g. 12345
func (client DeletedWebAppsClient) GetDeletedWebAppByLocation(ctx context.Context, location string, deletedSiteID string) (result DeletedSite, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedWebAppsClient.GetDeletedWebAppByLocation")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDeletedWebAppByLocationPreparer(ctx, location, deletedSiteID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "GetDeletedWebAppByLocation", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDeletedWebAppByLocationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "GetDeletedWebAppByLocation", resp, "Failure sending request")
		return
	}

	result, err = client.GetDeletedWebAppByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "GetDeletedWebAppByLocation", resp, "Failure responding to request")
		return
	}

	return
}

// GetDeletedWebAppByLocationPreparer prepares the GetDeletedWebAppByLocation request.
func (client DeletedWebAppsClient) GetDeletedWebAppByLocationPreparer(ctx context.Context, location string, deletedSiteID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"deletedSiteId":  autorest.Encode("path", deletedSiteID),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/locations/{location}/deletedSites/{deletedSiteId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDeletedWebAppByLocationSender sends the GetDeletedWebAppByLocation request. The method will close the
// http.Response Body if it receives an error.
func (client DeletedWebAppsClient) GetDeletedWebAppByLocationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetDeletedWebAppByLocationResponder handles the response to the GetDeletedWebAppByLocation request. The method always
// closes the http.Response Body.
func (client DeletedWebAppsClient) GetDeletedWebAppByLocationResponder(resp *http.Response) (result DeletedSite, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List description for Get all deleted apps for a subscription.
func (client DeletedWebAppsClient) List(ctx context.Context) (result DeletedWebAppCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedWebAppsClient.List")
		defer func() {
			sc := -1
			if result.dwac.Response.Response != nil {
				sc = result.dwac.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dwac.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "List", resp, "Failure sending request")
		return
	}

	result.dwac, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.dwac.hasNextLink() && result.dwac.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DeletedWebAppsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/deletedSites", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DeletedWebAppsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DeletedWebAppsClient) ListResponder(resp *http.Response) (result DeletedWebAppCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DeletedWebAppsClient) listNextResults(ctx context.Context, lastResults DeletedWebAppCollection) (result DeletedWebAppCollection, err error) {
	req, err := lastResults.deletedWebAppCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DeletedWebAppsClient) ListComplete(ctx context.Context) (result DeletedWebAppCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedWebAppsClient.List")
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

// ListByLocation description for Get all deleted apps for a subscription at location
func (client DeletedWebAppsClient) ListByLocation(ctx context.Context, location string) (result DeletedWebAppCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedWebAppsClient.ListByLocation")
		defer func() {
			sc := -1
			if result.dwac.Response.Response != nil {
				sc = result.dwac.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByLocationNextResults
	req, err := client.ListByLocationPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "ListByLocation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByLocationSender(req)
	if err != nil {
		result.dwac.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "ListByLocation", resp, "Failure sending request")
		return
	}

	result.dwac, err = client.ListByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "ListByLocation", resp, "Failure responding to request")
		return
	}
	if result.dwac.hasNextLink() && result.dwac.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByLocationPreparer prepares the ListByLocation request.
func (client DeletedWebAppsClient) ListByLocationPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/locations/{location}/deletedSites", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByLocationSender sends the ListByLocation request. The method will close the
// http.Response Body if it receives an error.
func (client DeletedWebAppsClient) ListByLocationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByLocationResponder handles the response to the ListByLocation request. The method always
// closes the http.Response Body.
func (client DeletedWebAppsClient) ListByLocationResponder(resp *http.Response) (result DeletedWebAppCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByLocationNextResults retrieves the next set of results, if any.
func (client DeletedWebAppsClient) listByLocationNextResults(ctx context.Context, lastResults DeletedWebAppCollection) (result DeletedWebAppCollection, err error) {
	req, err := lastResults.deletedWebAppCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "listByLocationNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByLocationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "listByLocationNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.DeletedWebAppsClient", "listByLocationNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByLocationComplete enumerates all values, automatically crossing page boundaries as required.
func (client DeletedWebAppsClient) ListByLocationComplete(ctx context.Context, location string) (result DeletedWebAppCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeletedWebAppsClient.ListByLocation")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByLocation(ctx, location)
	return
}
