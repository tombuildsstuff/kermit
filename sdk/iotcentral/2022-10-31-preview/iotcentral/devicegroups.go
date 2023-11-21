package iotcentral

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

// DeviceGroupsClient is the azure IoT Central is a service that makes it easy to connect, monitor, and manage your IoT
// devices at scale.
type DeviceGroupsClient struct {
	BaseClient
}

// NewDeviceGroupsClient creates an instance of the DeviceGroupsClient client.
func NewDeviceGroupsClient(subdomain string) DeviceGroupsClient {
	return DeviceGroupsClient{New(subdomain)}
}

// Create create or update a device group.
// Parameters:
// deviceGroupID - unique ID for the device group.
// body - device group body.
func (client DeviceGroupsClient) Create(ctx context.Context, deviceGroupID string, body DeviceGroup) (result DeviceGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceGroupsClient.Create")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: deviceGroupID,
			Constraints: []validation.Constraint{{Target: "deviceGroupID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "deviceGroupID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.DisplayName", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.Filter", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.DeviceGroupsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, deviceGroupID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "Create", resp, "Failure responding to request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client DeviceGroupsClient) CreatePreparer(ctx context.Context, deviceGroupID string, body DeviceGroup) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"deviceGroupId": autorest.Encode("path", deviceGroupID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/deviceGroups/{deviceGroupId}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceGroupsClient) CreateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client DeviceGroupsClient) CreateResponder(resp *http.Response) (result DeviceGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get sends the get request.
// Parameters:
// deviceGroupID - unique ID for the device group.
func (client DeviceGroupsClient) Get(ctx context.Context, deviceGroupID string) (result DeviceGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: deviceGroupID,
			Constraints: []validation.Constraint{{Target: "deviceGroupID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "deviceGroupID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.DeviceGroupsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, deviceGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DeviceGroupsClient) GetPreparer(ctx context.Context, deviceGroupID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"deviceGroupId": autorest.Encode("path", deviceGroupID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/deviceGroups/{deviceGroupId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeviceGroupsClient) GetResponder(resp *http.Response) (result DeviceGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetDevices get the list of devices by device group ID.
// Parameters:
// deviceGroupID - unique ID for the device group.
// maxpagesize - the maximum number of resources to return from one response.
func (client DeviceGroupsClient) GetDevices(ctx context.Context, deviceGroupID string, maxpagesize *int32) (result DeviceGroupDeviceCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceGroupsClient.GetDevices")
		defer func() {
			sc := -1
			if result.dgdc.Response.Response != nil {
				sc = result.dgdc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: deviceGroupID,
			Constraints: []validation.Constraint{{Target: "deviceGroupID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "deviceGroupID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}},
		{TargetValue: maxpagesize,
			Constraints: []validation.Constraint{{Target: "maxpagesize", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxpagesize", Name: validation.InclusiveMaximum, Rule: int64(100), Chain: nil},
					{Target: "maxpagesize", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("iotcentral.DeviceGroupsClient", "GetDevices", err.Error())
	}

	result.fn = client.getDevicesNextResults
	req, err := client.GetDevicesPreparer(ctx, deviceGroupID, maxpagesize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "GetDevices", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDevicesSender(req)
	if err != nil {
		result.dgdc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "GetDevices", resp, "Failure sending request")
		return
	}

	result.dgdc, err = client.GetDevicesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "GetDevices", resp, "Failure responding to request")
		return
	}
	if result.dgdc.hasNextLink() && result.dgdc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetDevicesPreparer prepares the GetDevices request.
func (client DeviceGroupsClient) GetDevicesPreparer(ctx context.Context, deviceGroupID string, maxpagesize *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"deviceGroupId": autorest.Encode("path", deviceGroupID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if maxpagesize != nil {
		queryParameters["maxpagesize"] = autorest.Encode("query", *maxpagesize)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/deviceGroups/{deviceGroupId}/devices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDevicesSender sends the GetDevices request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceGroupsClient) GetDevicesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetDevicesResponder handles the response to the GetDevices request. The method always
// closes the http.Response Body.
func (client DeviceGroupsClient) GetDevicesResponder(resp *http.Response) (result DeviceGroupDeviceCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getDevicesNextResults retrieves the next set of results, if any.
func (client DeviceGroupsClient) getDevicesNextResults(ctx context.Context, lastResults DeviceGroupDeviceCollection) (result DeviceGroupDeviceCollection, err error) {
	req, err := lastResults.deviceGroupDeviceCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "getDevicesNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetDevicesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "getDevicesNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetDevicesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "getDevicesNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetDevicesComplete enumerates all values, automatically crossing page boundaries as required.
func (client DeviceGroupsClient) GetDevicesComplete(ctx context.Context, deviceGroupID string, maxpagesize *int32) (result DeviceGroupDeviceCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceGroupsClient.GetDevices")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetDevices(ctx, deviceGroupID, maxpagesize)
	return
}

// List sends the list request.
// Parameters:
// filter - an expression on the resource type that selects the resources to be returned.
// maxpagesize - the maximum number of resources to return from one response.
// orderby - an expression that specify the order of the returned resources.
func (client DeviceGroupsClient) List(ctx context.Context, filter string, maxpagesize *int32, orderby string) (result DeviceGroupCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceGroupsClient.List")
		defer func() {
			sc := -1
			if result.dgc.Response.Response != nil {
				sc = result.dgc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: maxpagesize,
			Constraints: []validation.Constraint{{Target: "maxpagesize", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "maxpagesize", Name: validation.InclusiveMaximum, Rule: int64(100), Chain: nil},
					{Target: "maxpagesize", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("iotcentral.DeviceGroupsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter, maxpagesize, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dgc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "List", resp, "Failure sending request")
		return
	}

	result.dgc, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.dgc.hasNextLink() && result.dgc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DeviceGroupsClient) ListPreparer(ctx context.Context, filter string, maxpagesize *int32, orderby string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["filter"] = autorest.Encode("query", filter)
	}
	if maxpagesize != nil {
		queryParameters["maxpagesize"] = autorest.Encode("query", *maxpagesize)
	}
	if len(orderby) > 0 {
		queryParameters["orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPath("/deviceGroups"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceGroupsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DeviceGroupsClient) ListResponder(resp *http.Response) (result DeviceGroupCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DeviceGroupsClient) listNextResults(ctx context.Context, lastResults DeviceGroupCollection) (result DeviceGroupCollection, err error) {
	req, err := lastResults.deviceGroupCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DeviceGroupsClient) ListComplete(ctx context.Context, filter string, maxpagesize *int32, orderby string) (result DeviceGroupCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceGroupsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter, maxpagesize, orderby)
	return
}

// Remove sends the remove request.
// Parameters:
// deviceGroupID - unique ID for the device group.
func (client DeviceGroupsClient) Remove(ctx context.Context, deviceGroupID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceGroupsClient.Remove")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: deviceGroupID,
			Constraints: []validation.Constraint{{Target: "deviceGroupID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "deviceGroupID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.DeviceGroupsClient", "Remove", err.Error())
	}

	req, err := client.RemovePreparer(ctx, deviceGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "Remove", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "Remove", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "Remove", resp, "Failure responding to request")
		return
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client DeviceGroupsClient) RemovePreparer(ctx context.Context, deviceGroupID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"deviceGroupId": autorest.Encode("path", deviceGroupID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/deviceGroups/{deviceGroupId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceGroupsClient) RemoveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client DeviceGroupsClient) RemoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update update an existing device group by ID.
// Parameters:
// deviceGroupID - unique ID for the device group.
// body - device group patch body.
// ifMatch - only perform the operation if the entity's etag matches one of the etags provided or * is
// provided.
func (client DeviceGroupsClient) Update(ctx context.Context, deviceGroupID string, body interface{}, ifMatch string) (result DeviceGroup, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DeviceGroupsClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: deviceGroupID,
			Constraints: []validation.Constraint{{Target: "deviceGroupID", Name: validation.MaxLength, Rule: 255, Chain: nil},
				{Target: "deviceGroupID", Name: validation.Pattern, Rule: `^[a-zA-Z0-9_-]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iotcentral.DeviceGroupsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, deviceGroupID, body, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotcentral.DeviceGroupsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DeviceGroupsClient) UpdatePreparer(ctx context.Context, deviceGroupID string, body interface{}, ifMatch string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"baseDomain": client.BaseDomain,
		"subdomain":  autorest.Encode("path", client.Subdomain),
	}

	pathParameters := map[string]interface{}{
		"deviceGroupId": autorest.Encode("path", deviceGroupID),
	}

	const APIVersion = "2022-10-31-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("https://{subdomain}.{baseDomain}/api", urlParameters),
		autorest.WithPathParameters("/deviceGroups/{deviceGroupId}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DeviceGroupsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DeviceGroupsClient) UpdateResponder(resp *http.Response) (result DeviceGroup, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
