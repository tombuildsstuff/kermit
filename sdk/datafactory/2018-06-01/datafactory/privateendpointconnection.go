package datafactory

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

// PrivateEndpointConnectionClient is the the Azure Data Factory V2 management API provides a RESTful set of web
// services that interact with Azure Data Factory V2 services.
type PrivateEndpointConnectionClient struct {
    BaseClient
}
// NewPrivateEndpointConnectionClient creates an instance of the PrivateEndpointConnectionClient client.
func NewPrivateEndpointConnectionClient(subscriptionID string) PrivateEndpointConnectionClient {
    return NewPrivateEndpointConnectionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPrivateEndpointConnectionClientWithBaseURI creates an instance of the PrivateEndpointConnectionClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
    func NewPrivateEndpointConnectionClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionClient {
        return PrivateEndpointConnectionClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate approves or rejects a private endpoint connection
    // Parameters:
        // resourceGroupName - the resource group name.
        // factoryName - the factory name.
        // privateEndpointConnectionName - the private endpoint connection name.
        // ifMatch - eTag of the private endpoint connection entity.  Should only be specified for update, for which it
        // should match existing entity or can be * for unconditional update.
func (client PrivateEndpointConnectionClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, privateEndpointConnectionName string, privateEndpointWrapper PrivateLinkConnectionApprovalRequestResource, ifMatch string) (result PrivateEndpointConnectionResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PrivateEndpointConnectionClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
        { TargetValue: factoryName,
         Constraints: []validation.Constraint{	{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil },
        	{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil },
        	{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("datafactory.PrivateEndpointConnectionClient", "CreateOrUpdate", err.Error())
        }

        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, factoryName, privateEndpointConnectionName, privateEndpointWrapper, ifMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datafactory.PrivateEndpointConnectionClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        resp, err := client.CreateOrUpdateSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "datafactory.PrivateEndpointConnectionClient", "CreateOrUpdate", resp, "Failure sending request")
        return
        }

        result, err = client.CreateOrUpdateResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "datafactory.PrivateEndpointConnectionClient", "CreateOrUpdate", resp, "Failure responding to request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client PrivateEndpointConnectionClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, factoryName string, privateEndpointConnectionName string, privateEndpointWrapper PrivateLinkConnectionApprovalRequestResource, ifMatch string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "factoryName": autorest.Encode("path",factoryName),
        "privateEndpointConnectionName": autorest.Encode("path",privateEndpointConnectionName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2018-06-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/privateEndpointConnections/{privateEndpointConnectionName}",pathParameters),
autorest.WithJSON(privateEndpointWrapper),
autorest.WithQueryParameters(queryParameters))
        if len(ifMatch) > 0 {
        preparer = autorest.DecoratePreparer(preparer,
        autorest.WithHeader("If-Match",autorest.String(ifMatch)))
        }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client PrivateEndpointConnectionClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
    // closes the http.Response Body.
    func (client PrivateEndpointConnectionClient) CreateOrUpdateResponder(resp *http.Response) (result PrivateEndpointConnectionResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete deletes a private endpoint connection
    // Parameters:
        // resourceGroupName - the resource group name.
        // factoryName - the factory name.
        // privateEndpointConnectionName - the private endpoint connection name.
func (client PrivateEndpointConnectionClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, privateEndpointConnectionName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PrivateEndpointConnectionClient.Delete")
        defer func() {
            sc := -1
        if result.Response != nil {
        sc = result.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
        { TargetValue: factoryName,
         Constraints: []validation.Constraint{	{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil },
        	{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil },
        	{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("datafactory.PrivateEndpointConnectionClient", "Delete", err.Error())
        }

        req, err := client.DeletePreparer(ctx, resourceGroupName, factoryName, privateEndpointConnectionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datafactory.PrivateEndpointConnectionClient", "Delete", nil , "Failure preparing request")
    return
    }

        resp, err := client.DeleteSender(req)
        if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, "datafactory.PrivateEndpointConnectionClient", "Delete", resp, "Failure sending request")
        return
        }

        result, err = client.DeleteResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "datafactory.PrivateEndpointConnectionClient", "Delete", resp, "Failure responding to request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client PrivateEndpointConnectionClient) DeletePreparer(ctx context.Context, resourceGroupName string, factoryName string, privateEndpointConnectionName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "factoryName": autorest.Encode("path",factoryName),
        "privateEndpointConnectionName": autorest.Encode("path",privateEndpointConnectionName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2018-06-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/privateEndpointConnections/{privateEndpointConnectionName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client PrivateEndpointConnectionClient) DeleteSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // DeleteResponder handles the response to the Delete request. The method always
    // closes the http.Response Body.
    func (client PrivateEndpointConnectionClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get gets a private endpoint connection
    // Parameters:
        // resourceGroupName - the resource group name.
        // factoryName - the factory name.
        // privateEndpointConnectionName - the private endpoint connection name.
        // ifNoneMatch - eTag of the private endpoint connection entity. Should only be specified for get. If the ETag
        // matches the existing entity tag, or if * was provided, then no content will be returned.
func (client PrivateEndpointConnectionClient) Get(ctx context.Context, resourceGroupName string, factoryName string, privateEndpointConnectionName string, ifNoneMatch string) (result PrivateEndpointConnectionResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/PrivateEndpointConnectionClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil }}},
        { TargetValue: factoryName,
         Constraints: []validation.Constraint{	{Target: "factoryName", Name: validation.MaxLength, Rule: 63, Chain: nil },
        	{Target: "factoryName", Name: validation.MinLength, Rule: 3, Chain: nil },
        	{Target: "factoryName", Name: validation.Pattern, Rule: `^[A-Za-z0-9]+(?:-[A-Za-z0-9]+)*$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("datafactory.PrivateEndpointConnectionClient", "Get", err.Error())
        }

        req, err := client.GetPreparer(ctx, resourceGroupName, factoryName, privateEndpointConnectionName, ifNoneMatch)
    if err != nil {
    err = autorest.NewErrorWithError(err, "datafactory.PrivateEndpointConnectionClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "datafactory.PrivateEndpointConnectionClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "datafactory.PrivateEndpointConnectionClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client PrivateEndpointConnectionClient) GetPreparer(ctx context.Context, resourceGroupName string, factoryName string, privateEndpointConnectionName string, ifNoneMatch string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "factoryName": autorest.Encode("path",factoryName),
        "privateEndpointConnectionName": autorest.Encode("path",privateEndpointConnectionName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2018-06-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/privateEndpointConnections/{privateEndpointConnectionName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
        if len(ifNoneMatch) > 0 {
        preparer = autorest.DecoratePreparer(preparer,
        autorest.WithHeader("If-None-Match",autorest.String(ifNoneMatch)))
        }
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client PrivateEndpointConnectionClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client PrivateEndpointConnectionClient) GetResponder(resp *http.Response) (result PrivateEndpointConnectionResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

