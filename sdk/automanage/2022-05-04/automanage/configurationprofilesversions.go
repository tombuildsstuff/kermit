package automanage

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

// ConfigurationProfilesVersionsClient is the automanage Client
type ConfigurationProfilesVersionsClient struct {
    BaseClient
}
// NewConfigurationProfilesVersionsClient creates an instance of the ConfigurationProfilesVersionsClient client.
func NewConfigurationProfilesVersionsClient(subscriptionID string) ConfigurationProfilesVersionsClient {
    return NewConfigurationProfilesVersionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewConfigurationProfilesVersionsClientWithBaseURI creates an instance of the ConfigurationProfilesVersionsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
    func NewConfigurationProfilesVersionsClientWithBaseURI(baseURI string, subscriptionID string) ConfigurationProfilesVersionsClient {
        return ConfigurationProfilesVersionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate creates a configuration profile version
    // Parameters:
        // configurationProfileName - name of the configuration profile.
        // versionName - the configuration profile version name.
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // parameters - parameters supplied to create or update configuration profile.
func (client ConfigurationProfilesVersionsClient) CreateOrUpdate(ctx context.Context, configurationProfileName string, versionName string, resourceGroupName string, parameters ConfigurationProfile) (result ConfigurationProfile, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConfigurationProfilesVersionsClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
        return result, validation.NewError("automanage.ConfigurationProfilesVersionsClient", "CreateOrUpdate", err.Error())
        }

        req, err := client.CreateOrUpdatePreparer(ctx, configurationProfileName, versionName, resourceGroupName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilesVersionsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        resp, err := client.CreateOrUpdateSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilesVersionsClient", "CreateOrUpdate", resp, "Failure sending request")
        return
        }

        result, err = client.CreateOrUpdateResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilesVersionsClient", "CreateOrUpdate", resp, "Failure responding to request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client ConfigurationProfilesVersionsClient) CreateOrUpdatePreparer(ctx context.Context, configurationProfileName string, versionName string, resourceGroupName string, parameters ConfigurationProfile) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationProfileName": autorest.Encode("path",configurationProfileName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "versionName": autorest.Encode("path",versionName),
        }

            const APIVersion = "2022-05-04"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

            parameters.SystemData = nil
    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfiles/{configurationProfileName}/versions/{versionName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConfigurationProfilesVersionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
    // closes the http.Response Body.
    func (client ConfigurationProfilesVersionsClient) CreateOrUpdateResponder(resp *http.Response) (result ConfigurationProfile, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete delete a configuration profile version
    // Parameters:
        // resourceGroupName - the name of the resource group. The name is case insensitive.
        // configurationProfileName - name of the configuration profile
        // versionName - the configuration profile version name.
func (client ConfigurationProfilesVersionsClient) Delete(ctx context.Context, resourceGroupName string, configurationProfileName string, versionName string) (result autorest.Response, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConfigurationProfilesVersionsClient.Delete")
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
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
        return result, validation.NewError("automanage.ConfigurationProfilesVersionsClient", "Delete", err.Error())
        }

        req, err := client.DeletePreparer(ctx, resourceGroupName, configurationProfileName, versionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilesVersionsClient", "Delete", nil , "Failure preparing request")
    return
    }

        resp, err := client.DeleteSender(req)
        if err != nil {
        result.Response = resp
        err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilesVersionsClient", "Delete", resp, "Failure sending request")
        return
        }

        result, err = client.DeleteResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilesVersionsClient", "Delete", resp, "Failure responding to request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client ConfigurationProfilesVersionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, configurationProfileName string, versionName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationProfileName": autorest.Encode("path",configurationProfileName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "versionName": autorest.Encode("path",versionName),
        }

            const APIVersion = "2022-05-04"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfiles/{configurationProfileName}/versions/{versionName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConfigurationProfilesVersionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // DeleteResponder handles the response to the Delete request. The method always
    // closes the http.Response Body.
    func (client ConfigurationProfilesVersionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get get information about a configuration profile version
    // Parameters:
        // configurationProfileName - the configuration profile name.
        // versionName - the configuration profile version name.
        // resourceGroupName - the name of the resource group. The name is case insensitive.
func (client ConfigurationProfilesVersionsClient) Get(ctx context.Context, configurationProfileName string, versionName string, resourceGroupName string) (result ConfigurationProfile, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConfigurationProfilesVersionsClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
        return result, validation.NewError("automanage.ConfigurationProfilesVersionsClient", "Get", err.Error())
        }

        req, err := client.GetPreparer(ctx, configurationProfileName, versionName, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilesVersionsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilesVersionsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilesVersionsClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client ConfigurationProfilesVersionsClient) GetPreparer(ctx context.Context, configurationProfileName string, versionName string, resourceGroupName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationProfileName": autorest.Encode("path",configurationProfileName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        "versionName": autorest.Encode("path",versionName),
        }

            const APIVersion = "2022-05-04"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfiles/{configurationProfileName}/versions/{versionName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConfigurationProfilesVersionsClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client ConfigurationProfilesVersionsClient) GetResponder(resp *http.Response) (result ConfigurationProfile, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// ListChildResources retrieve a list of configuration profile version for a configuration profile
    // Parameters:
        // configurationProfileName - name of the configuration profile.
        // resourceGroupName - the name of the resource group. The name is case insensitive.
func (client ConfigurationProfilesVersionsClient) ListChildResources(ctx context.Context, configurationProfileName string, resourceGroupName string) (result ConfigurationProfileList, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/ConfigurationProfilesVersionsClient.ListChildResources")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: client.SubscriptionID,
         Constraints: []validation.Constraint{	{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil }}},
        { TargetValue: resourceGroupName,
         Constraints: []validation.Constraint{	{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil },
        	{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil }}}}); err != nil {
        return result, validation.NewError("automanage.ConfigurationProfilesVersionsClient", "ListChildResources", err.Error())
        }

        req, err := client.ListChildResourcesPreparer(ctx, configurationProfileName, resourceGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilesVersionsClient", "ListChildResources", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListChildResourcesSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilesVersionsClient", "ListChildResources", resp, "Failure sending request")
        return
        }

        result, err = client.ListChildResourcesResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "automanage.ConfigurationProfilesVersionsClient", "ListChildResources", resp, "Failure responding to request")
        return
        }

    return
}

    // ListChildResourcesPreparer prepares the ListChildResources request.
    func (client ConfigurationProfilesVersionsClient) ListChildResourcesPreparer(ctx context.Context, configurationProfileName string, resourceGroupName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "configurationProfileName": autorest.Encode("path",configurationProfileName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-05-04"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automanage/configurationProfiles/{configurationProfileName}/versions",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListChildResourcesSender sends the ListChildResources request. The method will close the
    // http.Response Body if it receives an error.
    func (client ConfigurationProfilesVersionsClient) ListChildResourcesSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListChildResourcesResponder handles the response to the ListChildResources request. The method always
    // closes the http.Response Body.
    func (client ConfigurationProfilesVersionsClient) ListChildResourcesResponder(resp *http.Response) (result ConfigurationProfileList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

