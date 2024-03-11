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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// CertificateOrdersDiagnosticsClient is the webSite Management Client
type CertificateOrdersDiagnosticsClient struct {
	BaseClient
}

// NewCertificateOrdersDiagnosticsClient creates an instance of the CertificateOrdersDiagnosticsClient client.
func NewCertificateOrdersDiagnosticsClient(subscriptionID string) CertificateOrdersDiagnosticsClient {
	return NewCertificateOrdersDiagnosticsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCertificateOrdersDiagnosticsClientWithBaseURI creates an instance of the CertificateOrdersDiagnosticsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewCertificateOrdersDiagnosticsClientWithBaseURI(baseURI string, subscriptionID string) CertificateOrdersDiagnosticsClient {
	return CertificateOrdersDiagnosticsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetAppServiceCertificateOrderDetectorResponse description for Microsoft.CertificateRegistration call to get a
// detector response from App Lens.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// certificateOrderName - the certificate order name for which the response is needed.
// detectorName - the detector name which needs to be run.
// startTime - the start time for detector response.
// endTime - the end time for the detector response.
// timeGrain - the time grain for the detector response.
func (client CertificateOrdersDiagnosticsClient) GetAppServiceCertificateOrderDetectorResponse(ctx context.Context, resourceGroupName string, certificateOrderName string, detectorName string, startTime *date.Time, endTime *date.Time, timeGrain string) (result DetectorResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateOrdersDiagnosticsClient.GetAppServiceCertificateOrderDetectorResponse")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}},
		{TargetValue: timeGrain,
			Constraints: []validation.Constraint{{Target: "timeGrain", Name: validation.Empty, Rule: false,
				Chain: []validation.Constraint{{Target: "timeGrain", Name: validation.Pattern, Rule: `PT[1-9][0-9]+[SMH]`, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("web.CertificateOrdersDiagnosticsClient", "GetAppServiceCertificateOrderDetectorResponse", err.Error())
	}

	req, err := client.GetAppServiceCertificateOrderDetectorResponsePreparer(ctx, resourceGroupName, certificateOrderName, detectorName, startTime, endTime, timeGrain)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificateOrdersDiagnosticsClient", "GetAppServiceCertificateOrderDetectorResponse", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAppServiceCertificateOrderDetectorResponseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.CertificateOrdersDiagnosticsClient", "GetAppServiceCertificateOrderDetectorResponse", resp, "Failure sending request")
		return
	}

	result, err = client.GetAppServiceCertificateOrderDetectorResponseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificateOrdersDiagnosticsClient", "GetAppServiceCertificateOrderDetectorResponse", resp, "Failure responding to request")
		return
	}

	return
}

// GetAppServiceCertificateOrderDetectorResponsePreparer prepares the GetAppServiceCertificateOrderDetectorResponse request.
func (client CertificateOrdersDiagnosticsClient) GetAppServiceCertificateOrderDetectorResponsePreparer(ctx context.Context, resourceGroupName string, certificateOrderName string, detectorName string, startTime *date.Time, endTime *date.Time, timeGrain string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateOrderName": autorest.Encode("path", certificateOrderName),
		"detectorName":         autorest.Encode("path", detectorName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startTime != nil {
		queryParameters["startTime"] = autorest.Encode("query", *startTime)
	}
	if endTime != nil {
		queryParameters["endTime"] = autorest.Encode("query", *endTime)
	}
	if len(timeGrain) > 0 {
		queryParameters["timeGrain"] = autorest.Encode("query", timeGrain)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/detectors/{detectorName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAppServiceCertificateOrderDetectorResponseSender sends the GetAppServiceCertificateOrderDetectorResponse request. The method will close the
// http.Response Body if it receives an error.
func (client CertificateOrdersDiagnosticsClient) GetAppServiceCertificateOrderDetectorResponseSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetAppServiceCertificateOrderDetectorResponseResponder handles the response to the GetAppServiceCertificateOrderDetectorResponse request. The method always
// closes the http.Response Body.
func (client CertificateOrdersDiagnosticsClient) GetAppServiceCertificateOrderDetectorResponseResponder(resp *http.Response) (result DetectorResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAppServiceCertificateOrderDetectorResponse description for Microsoft.CertificateRegistration to get the list of
// detectors for this RP.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// certificateOrderName - the certificate order name for which the response is needed.
func (client CertificateOrdersDiagnosticsClient) ListAppServiceCertificateOrderDetectorResponse(ctx context.Context, resourceGroupName string, certificateOrderName string) (result DetectorResponseCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateOrdersDiagnosticsClient.ListAppServiceCertificateOrderDetectorResponse")
		defer func() {
			sc := -1
			if result.drc.Response.Response != nil {
				sc = result.drc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.CertificateOrdersDiagnosticsClient", "ListAppServiceCertificateOrderDetectorResponse", err.Error())
	}

	result.fn = client.listAppServiceCertificateOrderDetectorResponseNextResults
	req, err := client.ListAppServiceCertificateOrderDetectorResponsePreparer(ctx, resourceGroupName, certificateOrderName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificateOrdersDiagnosticsClient", "ListAppServiceCertificateOrderDetectorResponse", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAppServiceCertificateOrderDetectorResponseSender(req)
	if err != nil {
		result.drc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "web.CertificateOrdersDiagnosticsClient", "ListAppServiceCertificateOrderDetectorResponse", resp, "Failure sending request")
		return
	}

	result.drc, err = client.ListAppServiceCertificateOrderDetectorResponseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificateOrdersDiagnosticsClient", "ListAppServiceCertificateOrderDetectorResponse", resp, "Failure responding to request")
		return
	}
	if result.drc.hasNextLink() && result.drc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListAppServiceCertificateOrderDetectorResponsePreparer prepares the ListAppServiceCertificateOrderDetectorResponse request.
func (client CertificateOrdersDiagnosticsClient) ListAppServiceCertificateOrderDetectorResponsePreparer(ctx context.Context, resourceGroupName string, certificateOrderName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateOrderName": autorest.Encode("path", certificateOrderName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CertificateRegistration/certificateOrders/{certificateOrderName}/detectors", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAppServiceCertificateOrderDetectorResponseSender sends the ListAppServiceCertificateOrderDetectorResponse request. The method will close the
// http.Response Body if it receives an error.
func (client CertificateOrdersDiagnosticsClient) ListAppServiceCertificateOrderDetectorResponseSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAppServiceCertificateOrderDetectorResponseResponder handles the response to the ListAppServiceCertificateOrderDetectorResponse request. The method always
// closes the http.Response Body.
func (client CertificateOrdersDiagnosticsClient) ListAppServiceCertificateOrderDetectorResponseResponder(resp *http.Response) (result DetectorResponseCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAppServiceCertificateOrderDetectorResponseNextResults retrieves the next set of results, if any.
func (client CertificateOrdersDiagnosticsClient) listAppServiceCertificateOrderDetectorResponseNextResults(ctx context.Context, lastResults DetectorResponseCollection) (result DetectorResponseCollection, err error) {
	req, err := lastResults.detectorResponseCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "web.CertificateOrdersDiagnosticsClient", "listAppServiceCertificateOrderDetectorResponseNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAppServiceCertificateOrderDetectorResponseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "web.CertificateOrdersDiagnosticsClient", "listAppServiceCertificateOrderDetectorResponseNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAppServiceCertificateOrderDetectorResponseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.CertificateOrdersDiagnosticsClient", "listAppServiceCertificateOrderDetectorResponseNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAppServiceCertificateOrderDetectorResponseComplete enumerates all values, automatically crossing page boundaries as required.
func (client CertificateOrdersDiagnosticsClient) ListAppServiceCertificateOrderDetectorResponseComplete(ctx context.Context, resourceGroupName string, certificateOrderName string) (result DetectorResponseCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CertificateOrdersDiagnosticsClient.ListAppServiceCertificateOrderDetectorResponse")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAppServiceCertificateOrderDetectorResponse(ctx, resourceGroupName, certificateOrderName)
	return
}
