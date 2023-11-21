package compute

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

// CapacityReservationsClient is the compute Client
type CapacityReservationsClient struct {
    BaseClient
}
// NewCapacityReservationsClient creates an instance of the CapacityReservationsClient client.
func NewCapacityReservationsClient(subscriptionID string) CapacityReservationsClient {
    return NewCapacityReservationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCapacityReservationsClientWithBaseURI creates an instance of the CapacityReservationsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
    func NewCapacityReservationsClientWithBaseURI(baseURI string, subscriptionID string) CapacityReservationsClient {
        return CapacityReservationsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// CreateOrUpdate the operation to create or update a capacity reservation. Please note some properties can be set only
// during capacity reservation creation. Please refer to https://aka.ms/CapacityReservation for more details.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // capacityReservationGroupName - the name of the capacity reservation group.
        // capacityReservationName - the name of the capacity reservation.
        // parameters - parameters supplied to the Create capacity reservation.
func (client CapacityReservationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, parameters CapacityReservation) (result CapacityReservationsCreateOrUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CapacityReservationsClient.CreateOrUpdate")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        if err := validation.Validate([]validation.Validation{
        { TargetValue: parameters,
         Constraints: []validation.Constraint{	{Target: "parameters.Sku", Name: validation.Null, Rule: true, Chain: nil }}}}); err != nil {
        return result, validation.NewError("compute.CapacityReservationsClient", "CreateOrUpdate", err.Error())
        }

        req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, capacityReservationGroupName, capacityReservationName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "CreateOrUpdate", nil , "Failure preparing request")
    return
    }

        result, err = client.CreateOrUpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
        return
        }

    return
}

    // CreateOrUpdatePreparer prepares the CreateOrUpdate request.
    func (client CapacityReservationsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, parameters CapacityReservation) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "capacityReservationGroupName": autorest.Encode("path",capacityReservationGroupName),
        "capacityReservationName": autorest.Encode("path",capacityReservationName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-03-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPut(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/capacityReservations/{capacityReservationName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
    // http.Response Body if it receives an error.
    func (client CapacityReservationsClient) CreateOrUpdateSender(req *http.Request) (future CapacityReservationsCreateOrUpdateFuture, err error) {
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
    func (client CapacityReservationsClient) CreateOrUpdateResponder(resp *http.Response) (result CapacityReservation, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusCreated),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// Delete the operation to delete a capacity reservation. This operation is allowed only when all the associated
// resources are disassociated from the capacity reservation. Please refer to https://aka.ms/CapacityReservation for
// more details.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // capacityReservationGroupName - the name of the capacity reservation group.
        // capacityReservationName - the name of the capacity reservation.
func (client CapacityReservationsClient) Delete(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string) (result CapacityReservationsDeleteFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CapacityReservationsClient.Delete")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.DeletePreparer(ctx, resourceGroupName, capacityReservationGroupName, capacityReservationName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "Delete", nil , "Failure preparing request")
    return
    }

        result, err = client.DeleteSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "Delete", result.Response(), "Failure sending request")
        return
        }

    return
}

    // DeletePreparer prepares the Delete request.
    func (client CapacityReservationsClient) DeletePreparer(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "capacityReservationGroupName": autorest.Encode("path",capacityReservationGroupName),
        "capacityReservationName": autorest.Encode("path",capacityReservationName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-03-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsDelete(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/capacityReservations/{capacityReservationName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // DeleteSender sends the Delete request. The method will close the
    // http.Response Body if it receives an error.
    func (client CapacityReservationsClient) DeleteSender(req *http.Request) (future CapacityReservationsDeleteFuture, err error) {
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
    func (client CapacityReservationsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted,http.StatusNoContent),
            autorest.ByClosing())
            result.Response = resp
            return
    }

// Get the operation that retrieves information about the capacity reservation.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // capacityReservationGroupName - the name of the capacity reservation group.
        // capacityReservationName - the name of the capacity reservation.
        // expand - the expand expression to apply on the operation. 'InstanceView' retrieves a snapshot of the runtime
        // properties of the capacity reservation that is managed by the platform and can change outside of control
        // plane operations.
func (client CapacityReservationsClient) Get(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, expand CapacityReservationInstanceViewTypes) (result CapacityReservation, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CapacityReservationsClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, resourceGroupName, capacityReservationGroupName, capacityReservationName, expand)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client CapacityReservationsClient) GetPreparer(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, expand CapacityReservationInstanceViewTypes) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "capacityReservationGroupName": autorest.Encode("path",capacityReservationGroupName),
        "capacityReservationName": autorest.Encode("path",capacityReservationName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-03-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
        if len(string(expand)) > 0 {
        queryParameters["$expand"] = autorest.Encode("query",expand)
        }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/capacityReservations/{capacityReservationName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client CapacityReservationsClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client CapacityReservationsClient) GetResponder(resp *http.Response) (result CapacityReservation, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// ListByCapacityReservationGroup lists all of the capacity reservations in the specified capacity reservation group.
// Use the nextLink property in the response to get the next page of capacity reservations.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // capacityReservationGroupName - the name of the capacity reservation group.
func (client CapacityReservationsClient) ListByCapacityReservationGroup(ctx context.Context, resourceGroupName string, capacityReservationGroupName string) (result CapacityReservationListResultPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CapacityReservationsClient.ListByCapacityReservationGroup")
        defer func() {
            sc := -1
        if result.crlr.Response.Response != nil {
        sc = result.crlr.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.listByCapacityReservationGroupNextResults
    req, err := client.ListByCapacityReservationGroupPreparer(ctx, resourceGroupName, capacityReservationGroupName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "ListByCapacityReservationGroup", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListByCapacityReservationGroupSender(req)
        if err != nil {
        result.crlr.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "ListByCapacityReservationGroup", resp, "Failure sending request")
        return
        }

        result.crlr, err = client.ListByCapacityReservationGroupResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "ListByCapacityReservationGroup", resp, "Failure responding to request")
        return
        }
            if result.crlr.hasNextLink() && result.crlr.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListByCapacityReservationGroupPreparer prepares the ListByCapacityReservationGroup request.
    func (client CapacityReservationsClient) ListByCapacityReservationGroupPreparer(ctx context.Context, resourceGroupName string, capacityReservationGroupName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "capacityReservationGroupName": autorest.Encode("path",capacityReservationGroupName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-03-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/capacityReservations",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByCapacityReservationGroupSender sends the ListByCapacityReservationGroup request. The method will close the
    // http.Response Body if it receives an error.
    func (client CapacityReservationsClient) ListByCapacityReservationGroupSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListByCapacityReservationGroupResponder handles the response to the ListByCapacityReservationGroup request. The method always
    // closes the http.Response Body.
    func (client CapacityReservationsClient) ListByCapacityReservationGroupResponder(resp *http.Response) (result CapacityReservationListResult, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listByCapacityReservationGroupNextResults retrieves the next set of results, if any.
            func (client CapacityReservationsClient) listByCapacityReservationGroupNextResults(ctx context.Context, lastResults CapacityReservationListResult) (result CapacityReservationListResult, err error) {
            req, err := lastResults.capacityReservationListResultPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "listByCapacityReservationGroupNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListByCapacityReservationGroupSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "listByCapacityReservationGroupNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListByCapacityReservationGroupResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "listByCapacityReservationGroupNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListByCapacityReservationGroupComplete enumerates all values, automatically crossing page boundaries as required.
            func (client CapacityReservationsClient) ListByCapacityReservationGroupComplete(ctx context.Context, resourceGroupName string, capacityReservationGroupName string) (result CapacityReservationListResultIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/CapacityReservationsClient.ListByCapacityReservationGroup")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.ListByCapacityReservationGroup(ctx, resourceGroupName, capacityReservationGroupName)
                            return
            }

// Update the operation to update a capacity reservation.
    // Parameters:
        // resourceGroupName - the name of the resource group.
        // capacityReservationGroupName - the name of the capacity reservation group.
        // capacityReservationName - the name of the capacity reservation.
        // parameters - parameters supplied to the Update capacity reservation operation.
func (client CapacityReservationsClient) Update(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, parameters CapacityReservationUpdate) (result CapacityReservationsUpdateFuture, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/CapacityReservationsClient.Update")
        defer func() {
            sc := -1
        if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
        sc = result.FutureAPI.Response().StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.UpdatePreparer(ctx, resourceGroupName, capacityReservationGroupName, capacityReservationName, parameters)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "Update", nil , "Failure preparing request")
    return
    }

        result, err = client.UpdateSender(req)
        if err != nil {
        err = autorest.NewErrorWithError(err, "compute.CapacityReservationsClient", "Update", result.Response(), "Failure sending request")
        return
        }

    return
}

    // UpdatePreparer prepares the Update request.
    func (client CapacityReservationsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, capacityReservationGroupName string, capacityReservationName string, parameters CapacityReservationUpdate) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "capacityReservationGroupName": autorest.Encode("path",capacityReservationGroupName),
        "capacityReservationName": autorest.Encode("path",capacityReservationName),
        "resourceGroupName": autorest.Encode("path",resourceGroupName),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2023-03-01"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsContentType("application/json; charset=utf-8"),
autorest.AsPatch(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/capacityReservationGroups/{capacityReservationGroupName}/capacityReservations/{capacityReservationName}",pathParameters),
autorest.WithJSON(parameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // UpdateSender sends the Update request. The method will close the
    // http.Response Body if it receives an error.
    func (client CapacityReservationsClient) UpdateSender(req *http.Request) (future CapacityReservationsUpdateFuture, err error) {
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

    // UpdateResponder handles the response to the Update request. The method always
    // closes the http.Response Body.
    func (client CapacityReservationsClient) UpdateResponder(resp *http.Response) (result CapacityReservation, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK,http.StatusAccepted),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

