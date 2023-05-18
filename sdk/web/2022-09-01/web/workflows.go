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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
)

// WorkflowsClient is the webSite Management Client
type WorkflowsClient struct {
	BaseClient
}

// NewWorkflowsClient creates an instance of the WorkflowsClient client.
func NewWorkflowsClient(subscriptionID string) WorkflowsClient {
	return NewWorkflowsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkflowsClientWithBaseURI creates an instance of the WorkflowsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWorkflowsClientWithBaseURI(baseURI string, subscriptionID string) WorkflowsClient {
	return WorkflowsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// RegenerateAccessKey regenerates the callback URL access key for request triggers.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// name - site name.
// workflowName - the workflow name.
// keyType - the access key type.
func (client WorkflowsClient) RegenerateAccessKey(ctx context.Context, resourceGroupName string, name string, workflowName string, keyType RegenerateActionParameter) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowsClient.RegenerateAccessKey")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.WorkflowsClient", "RegenerateAccessKey", err.Error())
	}

	req, err := client.RegenerateAccessKeyPreparer(ctx, resourceGroupName, name, workflowName, keyType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.WorkflowsClient", "RegenerateAccessKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegenerateAccessKeySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "web.WorkflowsClient", "RegenerateAccessKey", resp, "Failure sending request")
		return
	}

	result, err = client.RegenerateAccessKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.WorkflowsClient", "RegenerateAccessKey", resp, "Failure responding to request")
		return
	}

	return
}

// RegenerateAccessKeyPreparer prepares the RegenerateAccessKey request.
func (client WorkflowsClient) RegenerateAccessKeyPreparer(ctx context.Context, resourceGroupName string, name string, workflowName string, keyType RegenerateActionParameter) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/regenerateAccessKey", pathParameters),
		autorest.WithJSON(keyType),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RegenerateAccessKeySender sends the RegenerateAccessKey request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) RegenerateAccessKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RegenerateAccessKeyResponder handles the response to the RegenerateAccessKey request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) RegenerateAccessKeyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Validate validates the workflow definition.
// Parameters:
// resourceGroupName - name of the resource group to which the resource belongs.
// name - site name.
// workflowName - the workflow name.
// validate - the workflow.
func (client WorkflowsClient) Validate(ctx context.Context, resourceGroupName string, name string, workflowName string, validate Workflow) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkflowsClient.Validate")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+[^\.]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("web.WorkflowsClient", "Validate", err.Error())
	}

	req, err := client.ValidatePreparer(ctx, resourceGroupName, name, workflowName, validate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.WorkflowsClient", "Validate", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "web.WorkflowsClient", "Validate", resp, "Failure sending request")
		return
	}

	result, err = client.ValidateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "web.WorkflowsClient", "Validate", resp, "Failure responding to request")
		return
	}

	return
}

// ValidatePreparer prepares the Validate request.
func (client WorkflowsClient) ValidatePreparer(ctx context.Context, resourceGroupName string, name string, workflowName string, validate Workflow) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workflowName":      autorest.Encode("path", workflowName),
	}

	const APIVersion = "2022-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/hostruntime/runtime/webhooks/workflow/api/management/workflows/{workflowName}/validate", pathParameters),
		autorest.WithJSON(validate),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidateSender sends the Validate request. The method will close the
// http.Response Body if it receives an error.
func (client WorkflowsClient) ValidateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ValidateResponder handles the response to the Validate request. The method always
// closes the http.Response Body.
func (client WorkflowsClient) ValidateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
