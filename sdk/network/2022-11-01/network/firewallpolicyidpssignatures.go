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

// FirewallPolicyIdpsSignaturesClient is the network Client
type FirewallPolicyIdpsSignaturesClient struct {
    BaseClient
}
// NewFirewallPolicyIdpsSignaturesClient creates an instance of the FirewallPolicyIdpsSignaturesClient client.
func NewFirewallPolicyIdpsSignaturesClient(subscriptionID string) FirewallPolicyIdpsSignaturesClient {
    return NewFirewallPolicyIdpsSignaturesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFirewallPolicyIdpsSignaturesClientWithBaseURI creates an instance of the FirewallPolicyIdpsSignaturesClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
    func NewFirewallPolicyIdpsSignaturesClientWithBaseURI(baseURI string, subscriptionID string) FirewallPolicyIdpsSignaturesClient {
        return FirewallPolicyIdpsSignaturesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// List retrieves the current status of IDPS signatures for the relevant policy
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // firewallPolicyName - the name of the Firewall Policy.
func (client FirewallPolicyIdpsSignaturesClient) List(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters IDPSQueryObject) (result QueryResults, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/FirewallPolicyIdpsSignaturesClient.List")
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
         Constraints: []validation.Constraint{	{Target: "parameters.ResultsPerPage", Name: validation.Null, Rule: false ,
        Chain: []validation.Constraint{	{Target: "parameters.ResultsPerPage", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil },
        	{Target: "parameters.ResultsPerPage", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil },
        }}}}}); err != nil {
        return result, validation.NewError("network.FirewallPolicyIdpsSignaturesClient", "List", err.Error())
        }

        req, err := client.ListPreparer(ctx, resourceGroupName, firewallPolicyName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesClient", "List", resp, "Failure sending request")
        return
        }

        result, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "network.FirewallPolicyIdpsSignaturesClient", "List", resp, "Failure responding to request")
        return
        }

    return
}

    // ListPreparer prepares the List request.
    func (client FirewallPolicyIdpsSignaturesClient) ListPreparer(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters IDPSQueryObject) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "firewallPolicyName": autorest.Encode("path",firewallPolicyName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-11-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPost(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}/listIdpsSignatures",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client FirewallPolicyIdpsSignaturesClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client FirewallPolicyIdpsSignaturesClient) ListResponder(resp *http.Response) (result QueryResults, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

