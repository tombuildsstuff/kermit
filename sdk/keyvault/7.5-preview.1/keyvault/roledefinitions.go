package keyvault

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

// RoleDefinitionsClient is the the key vault client performs cryptographic key operations and vault operations against
// the Key Vault service.
type RoleDefinitionsClient struct {
	BaseClient
}

// NewRoleDefinitionsClient creates an instance of the RoleDefinitionsClient client.
func NewRoleDefinitionsClient() RoleDefinitionsClient {
	return RoleDefinitionsClient{New()}
}

// CreateOrUpdate creates or updates a custom role definition.
// Parameters:
// vaultBaseURL - the vault name, for example https://myvault.vault.azure.net.
// scope - the scope of the role definition to create or update. Managed HSM only supports '/'.
// roleDefinitionName - the name of the role definition to create or update. It can be any valid GUID.
// parameters - parameters for the role definition.
func (client RoleDefinitionsClient) CreateOrUpdate(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, parameters RoleDefinitionCreateParameters) (result RoleDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Properties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("keyvault.RoleDefinitionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, vaultBaseURL, scope, roleDefinitionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RoleDefinitionsClient) CreateOrUpdatePreparer(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string, parameters RoleDefinitionCreateParameters) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	pathParameters := map[string]interface{}{
		"roleDefinitionName": autorest.Encode("path", roleDefinitionName),
		"scope":              scope,
	}

	const APIVersion = "7.5-preview.1"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result RoleDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a custom role definition.
// Parameters:
// vaultBaseURL - the vault name, for example https://myvault.vault.azure.net.
// scope - the scope of the role definition to delete. Managed HSM only supports '/'.
// roleDefinitionName - the name (GUID) of the role definition to delete.
func (client RoleDefinitionsClient) Delete(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string) (result RoleDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, vaultBaseURL, scope, roleDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RoleDefinitionsClient) DeletePreparer(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	pathParameters := map[string]interface{}{
		"roleDefinitionName": autorest.Encode("path", roleDefinitionName),
		"scope":              scope,
	}

	const APIVersion = "7.5-preview.1"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) DeleteResponder(resp *http.Response) (result RoleDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get the specified role definition.
// Parameters:
// vaultBaseURL - the vault name, for example https://myvault.vault.azure.net.
// scope - the scope of the role definition to get. Managed HSM only supports '/'.
// roleDefinitionName - the name of the role definition to get.
func (client RoleDefinitionsClient) Get(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string) (result RoleDefinition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, vaultBaseURL, scope, roleDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RoleDefinitionsClient) GetPreparer(ctx context.Context, vaultBaseURL string, scope string, roleDefinitionName string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	pathParameters := map[string]interface{}{
		"roleDefinitionName": autorest.Encode("path", roleDefinitionName),
		"scope":              scope,
	}

	const APIVersion = "7.5-preview.1"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) GetResponder(resp *http.Response) (result RoleDefinition, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List get all role definitions that are applicable at scope and above.
// Parameters:
// vaultBaseURL - the vault name, for example https://myvault.vault.azure.net.
// scope - the scope of the role definition.
// filter - the filter to apply on the operation. Use atScopeAndBelow filter to search below the given scope as
// well.
func (client RoleDefinitionsClient) List(ctx context.Context, vaultBaseURL string, scope string, filter string) (result RoleDefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.List")
		defer func() {
			sc := -1
			if result.rdlr.Response.Response != nil {
				sc = result.rdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, vaultBaseURL, scope, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "List", resp, "Failure sending request")
		return
	}

	result.rdlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rdlr.hasNextLink() && result.rdlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client RoleDefinitionsClient) ListPreparer(ctx context.Context, vaultBaseURL string, scope string, filter string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"vaultBaseUrl": vaultBaseURL,
	}

	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "7.5-preview.1"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{vaultBaseUrl}", urlParameters),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/roleDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RoleDefinitionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RoleDefinitionsClient) ListResponder(resp *http.Response) (result RoleDefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client RoleDefinitionsClient) listNextResults(ctx context.Context, lastResults RoleDefinitionListResult) (result RoleDefinitionListResult, err error) {
	req, err := lastResults.roleDefinitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keyvault.RoleDefinitionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RoleDefinitionsClient) ListComplete(ctx context.Context, vaultBaseURL string, scope string, filter string) (result RoleDefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RoleDefinitionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, vaultBaseURL, scope, filter)
	return
}
