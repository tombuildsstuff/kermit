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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// ManagerDeploymentStatusClient is the network Client
type ManagerDeploymentStatusClient struct {
	BaseClient
}

// NewManagerDeploymentStatusClient creates an instance of the ManagerDeploymentStatusClient client.
func NewManagerDeploymentStatusClient(subscriptionID string) ManagerDeploymentStatusClient {
	return NewManagerDeploymentStatusClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagerDeploymentStatusClientWithBaseURI creates an instance of the ManagerDeploymentStatusClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewManagerDeploymentStatusClientWithBaseURI(baseURI string, subscriptionID string) ManagerDeploymentStatusClient {
	return ManagerDeploymentStatusClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List post to List of Network Manager Deployment Status.
// Parameters:
// parameters - parameters supplied to specify which Managed Network deployment status is.
// resourceGroupName - the name of the resource group.
// networkManagerName - the name of the network manager.
// top - an optional query parameter which specifies the maximum number of records to be returned by the
// server.
func (client ManagerDeploymentStatusClient) List(ctx context.Context, parameters ManagerDeploymentStatusParameter, resourceGroupName string, networkManagerName string, top *int32) (result ManagerDeploymentStatusListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagerDeploymentStatusClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMaximum, Rule: int64(20), Chain: nil},
					{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("network.ManagerDeploymentStatusClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, parameters, resourceGroupName, networkManagerName, top)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ManagerDeploymentStatusClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "network.ManagerDeploymentStatusClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.ManagerDeploymentStatusClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ManagerDeploymentStatusClient) ListPreparer(ctx context.Context, parameters ManagerDeploymentStatusParameter, resourceGroupName string, networkManagerName string, top *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"networkManagerName": autorest.Encode("path", networkManagerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/listDeploymentStatus", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ManagerDeploymentStatusClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ManagerDeploymentStatusClient) ListResponder(resp *http.Response) (result ManagerDeploymentStatusListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
