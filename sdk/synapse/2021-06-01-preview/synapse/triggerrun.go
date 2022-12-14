package synapse

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

// TriggerRunClient is the client for the TriggerRun methods of the Synapse service.
type TriggerRunClient struct {
	BaseClient
}

// NewTriggerRunClient creates an instance of the TriggerRunClient client.
func NewTriggerRunClient(endpoint string) TriggerRunClient {
	return TriggerRunClient{New(endpoint)}
}

// CancelTriggerInstance cancel single trigger instance by runId.
// Parameters:
// triggerName - the trigger name.
// runID - the pipeline run identifier.
func (client TriggerRunClient) CancelTriggerInstance(ctx context.Context, triggerName string, runID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerRunClient.CancelTriggerInstance")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.TriggerRunClient", "CancelTriggerInstance", err.Error())
	}

	req, err := client.CancelTriggerInstancePreparer(ctx, triggerName, runID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.TriggerRunClient", "CancelTriggerInstance", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelTriggerInstanceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "synapse.TriggerRunClient", "CancelTriggerInstance", resp, "Failure sending request")
		return
	}

	result, err = client.CancelTriggerInstanceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.TriggerRunClient", "CancelTriggerInstance", resp, "Failure responding to request")
		return
	}

	return
}

// CancelTriggerInstancePreparer prepares the CancelTriggerInstance request.
func (client TriggerRunClient) CancelTriggerInstancePreparer(ctx context.Context, triggerName string, runID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"runId":       autorest.Encode("path", runID),
		"triggerName": autorest.Encode("path", triggerName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/triggers/{triggerName}/triggerRuns/{runId}/cancel", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelTriggerInstanceSender sends the CancelTriggerInstance request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerRunClient) CancelTriggerInstanceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CancelTriggerInstanceResponder handles the response to the CancelTriggerInstance request. The method always
// closes the http.Response Body.
func (client TriggerRunClient) CancelTriggerInstanceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// QueryTriggerRunsByWorkspace query trigger runs.
// Parameters:
// filterParameters - parameters to filter the pipeline run.
func (client TriggerRunClient) QueryTriggerRunsByWorkspace(ctx context.Context, filterParameters RunFilterParameters) (result TriggerRunsQueryResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerRunClient.QueryTriggerRunsByWorkspace")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: filterParameters,
			Constraints: []validation.Constraint{{Target: "filterParameters.LastUpdatedAfter", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "filterParameters.LastUpdatedBefore", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.TriggerRunClient", "QueryTriggerRunsByWorkspace", err.Error())
	}

	req, err := client.QueryTriggerRunsByWorkspacePreparer(ctx, filterParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.TriggerRunClient", "QueryTriggerRunsByWorkspace", nil, "Failure preparing request")
		return
	}

	resp, err := client.QueryTriggerRunsByWorkspaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.TriggerRunClient", "QueryTriggerRunsByWorkspace", resp, "Failure sending request")
		return
	}

	result, err = client.QueryTriggerRunsByWorkspaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.TriggerRunClient", "QueryTriggerRunsByWorkspace", resp, "Failure responding to request")
		return
	}

	return
}

// QueryTriggerRunsByWorkspacePreparer prepares the QueryTriggerRunsByWorkspace request.
func (client TriggerRunClient) QueryTriggerRunsByWorkspacePreparer(ctx context.Context, filterParameters RunFilterParameters) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPath("/queryTriggerRuns"),
		autorest.WithJSON(filterParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// QueryTriggerRunsByWorkspaceSender sends the QueryTriggerRunsByWorkspace request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerRunClient) QueryTriggerRunsByWorkspaceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// QueryTriggerRunsByWorkspaceResponder handles the response to the QueryTriggerRunsByWorkspace request. The method always
// closes the http.Response Body.
func (client TriggerRunClient) QueryTriggerRunsByWorkspaceResponder(resp *http.Response) (result TriggerRunsQueryResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RerunTriggerInstance rerun single trigger instance by runId.
// Parameters:
// triggerName - the trigger name.
// runID - the pipeline run identifier.
func (client TriggerRunClient) RerunTriggerInstance(ctx context.Context, triggerName string, runID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TriggerRunClient.RerunTriggerInstance")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: triggerName,
			Constraints: []validation.Constraint{{Target: "triggerName", Name: validation.MaxLength, Rule: 260, Chain: nil},
				{Target: "triggerName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "triggerName", Name: validation.Pattern, Rule: `^[A-Za-z0-9_][^<>*#.%&:\\+?/]*$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.TriggerRunClient", "RerunTriggerInstance", err.Error())
	}

	req, err := client.RerunTriggerInstancePreparer(ctx, triggerName, runID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.TriggerRunClient", "RerunTriggerInstance", nil, "Failure preparing request")
		return
	}

	resp, err := client.RerunTriggerInstanceSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "synapse.TriggerRunClient", "RerunTriggerInstance", resp, "Failure sending request")
		return
	}

	result, err = client.RerunTriggerInstanceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.TriggerRunClient", "RerunTriggerInstance", resp, "Failure responding to request")
		return
	}

	return
}

// RerunTriggerInstancePreparer prepares the RerunTriggerInstance request.
func (client TriggerRunClient) RerunTriggerInstancePreparer(ctx context.Context, triggerName string, runID string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint": client.Endpoint,
	}

	pathParameters := map[string]interface{}{
		"runId":       autorest.Encode("path", runID),
		"triggerName": autorest.Encode("path", triggerName),
	}

	const APIVersion = "2021-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}", urlParameters),
		autorest.WithPathParameters("/triggers/{triggerName}/triggerRuns/{runId}/rerun", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RerunTriggerInstanceSender sends the RerunTriggerInstance request. The method will close the
// http.Response Body if it receives an error.
func (client TriggerRunClient) RerunTriggerInstanceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RerunTriggerInstanceResponder handles the response to the RerunTriggerInstance request. The method always
// closes the http.Response Body.
func (client TriggerRunClient) RerunTriggerInstanceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
