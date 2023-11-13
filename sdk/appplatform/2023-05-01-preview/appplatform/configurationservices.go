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
    "github.com/Azure/go-autorest/autorest/validation"
)

// ConfigurationServicesClient is the REST API for Azure Spring Apps
type ConfigurationServicesClient struct {
    BaseClient
}
// NewConfigurationServicesClient creates an instance of the ConfigurationServicesClient client.
func NewConfigurationServicesClient(subscriptionID string) ConfigurationServicesClient {
    return NewConfigurationServicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigurationServicesClientWithBaseURI creates an instance of the ConfigurationServicesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
    func NewConfigurationServicesClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationServicesClient {
        return ConfigurationServicesClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate create the default Application Configuration Service or update the existing Application Configuration
// Service.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // configurationServiceName - the name of Application Configuration Service.
        // configurationServiceResource - parameters for the update operation
func (client ConfigurationServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string, configurationServiceResource ConfigurationServiceResource) (result ConfigurationServicesCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConfigurationServicesClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: serviceName,
         Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9-]*[a-z0-9]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("appplatform.ConfigurationServicesClient", "CreateOrUpdate", err.Error())
        }

        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, configurationServiceName, configurationServiceResource)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        result, err = client.CreateOrUpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client ConfigurationServicesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string, configurationServiceResource ConfigurationServiceResource) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationServiceName": autorest.Encode("path",configurationServiceName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configurationServices/{configurationServiceName}",pathParameters),
autorest.WithJSON(configurationServiceResource),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConfigurationServicesClient) CreateOrUpdateSender(req *http.Request) (future ConfigurationServicesCreateOrUpdateFuture, err error) {
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
    func (client ConfigurationServicesClient) CreateOrUpdateResponder(resp *http.Response) (result ConfigurationServiceResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete disable the default Application Configuration Service.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // configurationServiceName - the name of Application Configuration Service.
func (client ConfigurationServicesClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string) (result ConfigurationServicesDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConfigurationServicesClient.Delete")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: serviceName,
         Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9-]*[a-z0-9]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("appplatform.ConfigurationServicesClient", "Delete", err.Error())
        }

        req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, configurationServiceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client ConfigurationServicesClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationServiceName": autorest.Encode("path",configurationServiceName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configurationServices/{configurationServiceName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConfigurationServicesClient) DeleteSender(req *http.Request) (future ConfigurationServicesDeleteFuture, err error) {
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
    func (client ConfigurationServicesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get get the Application Configuration Service and its properties.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // configurationServiceName - the name of Application Configuration Service.
func (client ConfigurationServicesClient) Get(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string) (result ConfigurationServiceResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConfigurationServicesClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: serviceName,
         Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9-]*[a-z0-9]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("appplatform.ConfigurationServicesClient", "Get", err.Error())
        }

        req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, configurationServiceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client ConfigurationServicesClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationServiceName": autorest.Encode("path",configurationServiceName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configurationServices/{configurationServiceName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConfigurationServicesClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client ConfigurationServicesClient) GetResponder(resp *http.Response) (result ConfigurationServiceResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List handles requests to list all resources in a Service.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
func (client ConfigurationServicesClient) List(ctx context.Context, resourceGroupName string, serviceName string) (result ConfigurationServiceResourceCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConfigurationServicesClient.List")
        defer func() {
            sc := -1
        if result.csrc.Response.Response != nil {
        sc = result.csrc.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: serviceName,
         Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9-]*[a-z0-9]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("appplatform.ConfigurationServicesClient", "List", err.Error())
        }

            result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, serviceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.csrc.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "List", resp, "Failure sending request")
        return
        }

        result.csrc, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "List", resp, "Failure responding to request")
        return
        }
            if result.csrc.hasNextLink() && result.csrc.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client ConfigurationServicesClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configurationServices",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConfigurationServicesClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client ConfigurationServicesClient) ListResponder(resp *http.Response) (result ConfigurationServiceResourceCollection, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client ConfigurationServicesClient) listNextResults(ctx context.Context, lastResults ConfigurationServiceResourceCollection) (result ConfigurationServiceResourceCollection, err error) {
            req, err := lastResults.configurationServiceResourceCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client ConfigurationServicesClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string) (result ConfigurationServiceResourceCollectionIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/ConfigurationServicesClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName, serviceName)
                            return
            }

// Validate check if the Application Configuration Service settings are valid.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // configurationServiceName - the name of Application Configuration Service.
        // settings - application Configuration Service settings to be validated
func (client ConfigurationServicesClient) Validate(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string, settings ConfigurationServiceSettings) (result ConfigurationServicesValidateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConfigurationServicesClient.Validate")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: serviceName,
         Constraints: []validation.Constraint{	{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-z][a-z0-9-]*[a-z0-9]$`, Chain: nil }}}}); err != nil {
        return result, validation.NewError("appplatform.ConfigurationServicesClient", "Validate", err.Error())
        }

        req, err := client.ValidatePreparer(ctx, resourceGroupName, serviceName, configurationServiceName, settings)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "Validate", nil , "Failure preparing request")
    return
    }

        result, err = client.ValidateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.ConfigurationServicesClient", "Validate", result.Response(), "Failure sending request")
        return
        }

    return
}

    // ValidatePreparer prepares the Validate request.
    func (client ConfigurationServicesClient) ValidatePreparer(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string, settings ConfigurationServiceSettings) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationServiceName": autorest.Encode("path",configurationServiceName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-05-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPost(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configurationServices/{configurationServiceName}/validate",pathParameters),
autorest.WithJSON(settings),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ValidateSender sends the Validate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConfigurationServicesClient) ValidateSender(req *http.Request) (future ConfigurationServicesValidateFuture, err error) {
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

    // ValidateResponder handles the response to the Validate request. The method always
    // closes the http.Response Body.
    func (client ConfigurationServicesClient) ValidateResponder(resp *http.Response) (result ConfigurationServiceSettingsValidateResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

