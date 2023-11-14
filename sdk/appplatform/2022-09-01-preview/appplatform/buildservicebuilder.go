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

// BuildServiceBuilderClient is the REST API for Azure Spring Apps
type BuildServiceBuilderClient struct {
    BaseClient
}
// NewBuildServiceBuilderClient creates an instance of the BuildServiceBuilderClient client.
func NewBuildServiceBuilderClient(subscriptionID string) BuildServiceBuilderClient {
    return NewBuildServiceBuilderClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBuildServiceBuilderClientWithBaseURI creates an instance of the BuildServiceBuilderClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
    func NewBuildServiceBuilderClientWithBaseURI(baseURI string, subscriptionID string) BuildServiceBuilderClient {
        return BuildServiceBuilderClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate create or update a KPack builder.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // buildServiceName - the name of the build service resource.
        // builderName - the name of the builder resource.
        // builderResource - the target builder for the create or update operation
func (client BuildServiceBuilderClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, builderResource BuilderResource) (result BuildServiceBuilderCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BuildServiceBuilderClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, buildServiceName, builderName, builderResource)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        result, err = client.CreateOrUpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "CreateOrUpdate", result.Response(), "Failure sending request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client BuildServiceBuilderClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, builderResource BuilderResource) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "buildServiceName": autorest.Encode("path",buildServiceName),
        "builderName": autorest.Encode("path",builderName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-09-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders/{builderName}",pathParameters),
autorest.WithJSON(builderResource),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client BuildServiceBuilderClient) CreateOrUpdateSender(req *http.Request) (future BuildServiceBuilderCreateOrUpdateFuture, err error) {
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
    func (client BuildServiceBuilderClient) CreateOrUpdateResponder(resp *http.Response) (result BuilderResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete delete a KPack builder.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // buildServiceName - the name of the build service resource.
        // builderName - the name of the builder resource.
func (client BuildServiceBuilderClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string) (result BuildServiceBuilderDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BuildServiceBuilderClient.Delete")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, buildServiceName, builderName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client BuildServiceBuilderClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "buildServiceName": autorest.Encode("path",buildServiceName),
        "builderName": autorest.Encode("path",builderName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-09-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders/{builderName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client BuildServiceBuilderClient) DeleteSender(req *http.Request) (future BuildServiceBuilderDeleteFuture, err error) {
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
    func (client BuildServiceBuilderClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get get a KPack builder.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // buildServiceName - the name of the build service resource.
        // builderName - the name of the builder resource.
func (client BuildServiceBuilderClient) Get(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string) (result BuilderResource, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BuildServiceBuilderClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, buildServiceName, builderName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client BuildServiceBuilderClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "buildServiceName": autorest.Encode("path",buildServiceName),
        "builderName": autorest.Encode("path",builderName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-09-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders/{builderName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client BuildServiceBuilderClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client BuildServiceBuilderClient) GetResponder(resp *http.Response) (result BuilderResource, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List list KPack builders result.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // buildServiceName - the name of the build service resource.
func (client BuildServiceBuilderClient) List(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (result BuilderResourceCollectionPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BuildServiceBuilderClient.List")
        defer func() {
            sc := -1
        if result.brc.Response.Response != nil {
        sc = result.brc.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, resourceGroupName, serviceName, buildServiceName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.brc.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "List", resp, "Failure sending request")
        return
        }

        result.brc, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "List", resp, "Failure responding to request")
        return
        }
            if result.brc.hasNextLink() && result.brc.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client BuildServiceBuilderClient) ListPreparer(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "buildServiceName": autorest.Encode("path",buildServiceName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-09-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client BuildServiceBuilderClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client BuildServiceBuilderClient) ListResponder(resp *http.Response) (result BuilderResourceCollection, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client BuildServiceBuilderClient) listNextResults(ctx context.Context, lastResults BuilderResourceCollection) (result BuilderResourceCollection, err error) {
            req, err := lastResults.builderResourceCollectionPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client BuildServiceBuilderClient) ListComplete(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (result BuilderResourceCollectionIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/BuildServiceBuilderClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, resourceGroupName, serviceName, buildServiceName)
                            return
            }

// ListDeployments list deployments that are using the builder.
    // Parameters:
        // resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
        // from the Azure Resource Manager API or the portal.
        // serviceName - the name of the Service resource.
        // buildServiceName - the name of the build service resource.
        // builderName - the name of the builder resource.
func (client BuildServiceBuilderClient) ListDeployments(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string) (result DeploymentList, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BuildServiceBuilderClient.ListDeployments")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.ListDeploymentsPreparer(ctx, resourceGroupName, serviceName, buildServiceName, builderName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "ListDeployments", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListDeploymentsSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "ListDeployments", resp, "Failure sending request")
        return
        }

        result, err = client.ListDeploymentsResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "appplatform.BuildServiceBuilderClient", "ListDeployments", resp, "Failure responding to request")
        return
        }

    return
}

    // ListDeploymentsPreparer prepares the ListDeployments request.
    func (client BuildServiceBuilderClient) ListDeploymentsPreparer(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "buildServiceName": autorest.Encode("path",buildServiceName),
        "builderName": autorest.Encode("path",builderName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "serviceName": autorest.Encode("path",serviceName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-09-01-preview"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsPost(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/builders/{builderName}/listUsingDeployments",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListDeploymentsSender sends the ListDeployments request. The method will close the
    // http.Response Body if it receives an error.
    func (client BuildServiceBuilderClient) ListDeploymentsSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListDeploymentsResponder handles the response to the ListDeployments request. The method always
    // closes the http.Response Body.
    func (client BuildServiceBuilderClient) ListDeploymentsResponder(resp *http.Response) (result DeploymentList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

