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
)

// BestPracticesVersionsClient is the automanage Client
type BestPracticesVersionsClient struct {
    BaseClient
}
// NewBestPracticesVersionsClient creates an instance of the BestPracticesVersionsClient client.
func NewBestPracticesVersionsClient(subscriptionID string) BestPracticesVersionsClient {
    return NewBestPracticesVersionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBestPracticesVersionsClientWithBaseURI creates an instance of the BestPracticesVersionsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
    func NewBestPracticesVersionsClientWithBaseURI(baseURI string, subscriptionID string) BestPracticesVersionsClient {
        return BestPracticesVersionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get get information about a Automanage best practice version
    // Parameters:
        // bestPracticeName - the Automanage best practice name.
        // versionName - the Automanage best practice version name.
func (client BestPracticesVersionsClient) Get(ctx context.Context, bestPracticeName string, versionName string) (result BestPractice, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BestPracticesVersionsClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, bestPracticeName, versionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automanage.BestPracticesVersionsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "automanage.BestPracticesVersionsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "automanage.BestPracticesVersionsClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client BestPracticesVersionsClient) GetPreparer(ctx context.Context, bestPracticeName string, versionName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "bestPracticeName": autorest.Encode("path",bestPracticeName),
        "versionName": autorest.Encode("path",versionName),
        }

            const APIVersion = "2022-05-04"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/providers/Microsoft.Automanage/bestPractices/{bestPracticeName}/versions/{versionName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client BestPracticesVersionsClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client BestPracticesVersionsClient) GetResponder(resp *http.Response) (result BestPractice, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// ListByTenant retrieve a list of Automanage best practices versions
    // Parameters:
        // bestPracticeName - the Automanage best practice name.
func (client BestPracticesVersionsClient) ListByTenant(ctx context.Context, bestPracticeName string) (result BestPracticeList, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/BestPracticesVersionsClient.ListByTenant")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.ListByTenantPreparer(ctx, bestPracticeName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "automanage.BestPracticesVersionsClient", "ListByTenant", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListByTenantSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "automanage.BestPracticesVersionsClient", "ListByTenant", resp, "Failure sending request")
        return
        }

        result, err = client.ListByTenantResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "automanage.BestPracticesVersionsClient", "ListByTenant", resp, "Failure responding to request")
        return
        }

    return
}

    // ListByTenantPreparer prepares the ListByTenant request.
    func (client BestPracticesVersionsClient) ListByTenantPreparer(ctx context.Context, bestPracticeName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "bestPracticeName": autorest.Encode("path",bestPracticeName),
        }

            const APIVersion = "2022-05-04"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/providers/Microsoft.Automanage/bestPractices/{bestPracticeName}/versions",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListByTenantSender sends the ListByTenant request. The method will close the
    // http.Response Body if it receives an error.
    func (client BestPracticesVersionsClient) ListByTenantSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
                }

    // ListByTenantResponder handles the response to the ListByTenant request. The method always
    // closes the http.Response Body.
    func (client BestPracticesVersionsClient) ListByTenantResponder(resp *http.Response) (result BestPracticeList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

