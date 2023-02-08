package network

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

// ApplicationGatewayWafDynamicManifestsClient is the network Client
type ApplicationGatewayWafDynamicManifestsClient struct {
	BaseClient
}

// NewApplicationGatewayWafDynamicManifestsClient creates an instance of the
// ApplicationGatewayWafDynamicManifestsClient client.
func NewApplicationGatewayWafDynamicManifestsClient(subscriptionID string) ApplicationGatewayWafDynamicManifestsClient {
	return NewApplicationGatewayWafDynamicManifestsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationGatewayWafDynamicManifestsClientWithBaseURI creates an instance of the
// ApplicationGatewayWafDynamicManifestsClient client using a custom endpoint.  Use this when interacting with an Azure
// cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewApplicationGatewayWafDynamicManifestsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationGatewayWafDynamicManifestsClient {
	return ApplicationGatewayWafDynamicManifestsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the regional application gateway waf manifest.
// Parameters:
// location - the region where the nrp are located at.
func (client ApplicationGatewayWafDynamicManifestsClient) Get(ctx context.Context, location string) (result ApplicationGatewayWafDynamicManifestResultListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewayWafDynamicManifestsClient.Get")
		defer func() {
			sc := -1
			if result.agwdmrl.Response.Response != nil {
				sc = result.agwdmrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getNextResults
	req, err := client.GetPreparer(ctx, location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayWafDynamicManifestsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.agwdmrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayWafDynamicManifestsClient", "Get", resp, "Failure sending request")
		return
	}

	result.agwdmrl, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayWafDynamicManifestsClient", "Get", resp, "Failure responding to request")
		return
	}
	if result.agwdmrl.hasNextLink() && result.agwdmrl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationGatewayWafDynamicManifestsClient) GetPreparer(ctx context.Context, location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/applicationGatewayWafDynamicManifests", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationGatewayWafDynamicManifestsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationGatewayWafDynamicManifestsClient) GetResponder(resp *http.Response) (result ApplicationGatewayWafDynamicManifestResultList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getNextResults retrieves the next set of results, if any.
func (client ApplicationGatewayWafDynamicManifestsClient) getNextResults(ctx context.Context, lastResults ApplicationGatewayWafDynamicManifestResultList) (result ApplicationGatewayWafDynamicManifestResultList, err error) {
	req, err := lastResults.applicationGatewayWafDynamicManifestResultListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.ApplicationGatewayWafDynamicManifestsClient", "getNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.ApplicationGatewayWafDynamicManifestsClient", "getNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ApplicationGatewayWafDynamicManifestsClient", "getNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetComplete enumerates all values, automatically crossing page boundaries as required.
func (client ApplicationGatewayWafDynamicManifestsClient) GetComplete(ctx context.Context, location string) (result ApplicationGatewayWafDynamicManifestResultListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ApplicationGatewayWafDynamicManifestsClient.Get")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.Get(ctx, location)
	return
}
