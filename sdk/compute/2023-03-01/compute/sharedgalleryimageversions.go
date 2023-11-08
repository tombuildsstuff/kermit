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
)

// SharedGalleryImageVersionsClient is the compute Client
type SharedGalleryImageVersionsClient struct {
    BaseClient
}
// NewSharedGalleryImageVersionsClient creates an instance of the SharedGalleryImageVersionsClient client.
func NewSharedGalleryImageVersionsClient(subscriptionID string) SharedGalleryImageVersionsClient {
    return NewSharedGalleryImageVersionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSharedGalleryImageVersionsClientWithBaseURI creates an instance of the SharedGalleryImageVersionsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
    func NewSharedGalleryImageVersionsClientWithBaseURI(baseURI string, subscriptionID string) SharedGalleryImageVersionsClient {
        return SharedGalleryImageVersionsClient{ NewWithBaseURI(baseURI, subscriptionID)}
    }

// Get get a shared gallery image version by subscription id or tenant id.
    // Parameters:
        // location - resource location.
        // galleryUniqueName - the unique name of the Shared Gallery.
        // galleryImageName - the name of the Shared Gallery Image Definition from which the Image Versions are to be
        // listed.
        // galleryImageVersionName - the name of the gallery image version to be created. Needs to follow semantic
        // version name pattern: The allowed characters are digit and period. Digits must be within the range of a
        // 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
func (client SharedGalleryImageVersionsClient) Get(ctx context.Context, location string, galleryUniqueName string, galleryImageName string, galleryImageVersionName string) (result SharedGalleryImageVersion, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SharedGalleryImageVersionsClient.Get")
        defer func() {
            sc := -1
        if result.Response.Response != nil {
        sc = result.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
    req, err := client.GetPreparer(ctx, location, galleryUniqueName, galleryImageName, galleryImageVersionName)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.SharedGalleryImageVersionsClient", "Get", nil , "Failure preparing request")
    return
    }

        resp, err := client.GetSender(req)
        if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "compute.SharedGalleryImageVersionsClient", "Get", resp, "Failure sending request")
        return
        }

        result, err = client.GetResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "compute.SharedGalleryImageVersionsClient", "Get", resp, "Failure responding to request")
        return
        }

    return
}

    // GetPreparer prepares the Get request.
    func (client SharedGalleryImageVersionsClient) GetPreparer(ctx context.Context, location string, galleryUniqueName string, galleryImageName string, galleryImageVersionName string) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "galleryImageName": autorest.Encode("path",galleryImageName),
        "galleryImageVersionName": autorest.Encode("path",galleryImageVersionName),
        "galleryUniqueName": autorest.Encode("path",galleryUniqueName),
        "location": autorest.Encode("path",location),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-03-03"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries/{galleryUniqueName}/images/{galleryImageName}/versions/{galleryImageVersionName}",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // GetSender sends the Get request. The method will close the
    // http.Response Body if it receives an error.
    func (client SharedGalleryImageVersionsClient) GetSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // GetResponder handles the response to the Get request. The method always
    // closes the http.Response Body.
    func (client SharedGalleryImageVersionsClient) GetResponder(resp *http.Response) (result SharedGalleryImageVersion, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

// List list shared gallery image versions by subscription id or tenant id.
    // Parameters:
        // location - resource location.
        // galleryUniqueName - the unique name of the Shared Gallery.
        // galleryImageName - the name of the Shared Gallery Image Definition from which the Image Versions are to be
        // listed.
        // sharedTo - the query parameter to decide what shared galleries to fetch when doing listing operations.
func (client SharedGalleryImageVersionsClient) List(ctx context.Context, location string, galleryUniqueName string, galleryImageName string, sharedTo SharedToValues) (result SharedGalleryImageVersionListPage, err error) {
    if tracing.IsEnabled() {
        ctx = tracing.StartSpan(ctx, fqdn + "/SharedGalleryImageVersionsClient.List")
        defer func() {
            sc := -1
        if result.sgivl.Response.Response != nil {
        sc = result.sgivl.Response.Response.StatusCode
        }
            tracing.EndSpan(ctx, sc, err)
        }()
    }
        result.fn = client.listNextResults
    req, err := client.ListPreparer(ctx, location, galleryUniqueName, galleryImageName, sharedTo)
    if err != nil {
    err = autorest.NewErrorWithError(err, "compute.SharedGalleryImageVersionsClient", "List", nil , "Failure preparing request")
    return
    }

        resp, err := client.ListSender(req)
        if err != nil {
        result.sgivl.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "compute.SharedGalleryImageVersionsClient", "List", resp, "Failure sending request")
        return
        }

        result.sgivl, err = client.ListResponder(resp)
        if err != nil {
        err = autorest.NewErrorWithError(err, "compute.SharedGalleryImageVersionsClient", "List", resp, "Failure responding to request")
        return
        }
            if result.sgivl.hasNextLink() && result.sgivl.IsEmpty() {
            err = result.NextWithContext(ctx)
            return
            }

    return
}

    // ListPreparer prepares the List request.
    func (client SharedGalleryImageVersionsClient) ListPreparer(ctx context.Context, location string, galleryUniqueName string, galleryImageName string, sharedTo SharedToValues) (*http.Request, error) {
        pathParameters := map[string]interface{} {
        "galleryImageName": autorest.Encode("path",galleryImageName),
        "galleryUniqueName": autorest.Encode("path",galleryUniqueName),
        "location": autorest.Encode("path",location),
        "subscriptionId": autorest.Encode("path",client.SubscriptionID),
        }

            const APIVersion = "2022-03-03"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
        if len(string(sharedTo)) > 0 {
        queryParameters["sharedTo"] = autorest.Encode("query",sharedTo)
        }

    preparer := autorest.CreatePreparer(
autorest.AsGet(),
autorest.WithBaseURL(client.BaseURI),
autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries/{galleryUniqueName}/images/{galleryImageName}/versions",pathParameters),
autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare((&http.Request{}).WithContext(ctx))
    }

    // ListSender sends the List request. The method will close the
    // http.Response Body if it receives an error.
    func (client SharedGalleryImageVersionsClient) ListSender(req *http.Request) (*http.Response, error) {
                return client.Send(req, azure.DoRetryWithRegistration(client.Client))
                }

    // ListResponder handles the response to the List request. The method always
    // closes the http.Response Body.
    func (client SharedGalleryImageVersionsClient) ListResponder(resp *http.Response) (result SharedGalleryImageVersionList, err error) {
            err = autorest.Respond(
            resp,
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
            result.Response = autorest.Response{Response: resp}
            return
    }

            // listNextResults retrieves the next set of results, if any.
            func (client SharedGalleryImageVersionsClient) listNextResults(ctx context.Context, lastResults SharedGalleryImageVersionList) (result SharedGalleryImageVersionList, err error) {
            req, err := lastResults.sharedGalleryImageVersionListPreparer(ctx)
            if err != nil {
            return result, autorest.NewErrorWithError(err, "compute.SharedGalleryImageVersionsClient", "listNextResults", nil , "Failure preparing next results request")
            }
            if req == nil {
            return
            }
            resp, err := client.ListSender(req)
            if err != nil {
            result.Response = autorest.Response{Response: resp}
            return result, autorest.NewErrorWithError(err, "compute.SharedGalleryImageVersionsClient", "listNextResults", resp, "Failure sending next results request")
            }
            result, err = client.ListResponder(resp)
            if err != nil {
            err = autorest.NewErrorWithError(err, "compute.SharedGalleryImageVersionsClient", "listNextResults", resp, "Failure responding to next results request")
            }
            return
                    }

            // ListComplete enumerates all values, automatically crossing page boundaries as required.
            func (client SharedGalleryImageVersionsClient) ListComplete(ctx context.Context, location string, galleryUniqueName string, galleryImageName string, sharedTo SharedToValues) (result SharedGalleryImageVersionListIterator, err error) {
            if tracing.IsEnabled() {
            ctx = tracing.StartSpan(ctx, fqdn + "/SharedGalleryImageVersionsClient.List")
            defer func() {
            sc := -1
            if result.Response().Response.Response != nil {
            sc = result.page.Response().Response.Response.StatusCode
            }
            tracing.EndSpan(ctx, sc, err)
            }()
            }
                    result.page, err = client.List(ctx, location, galleryUniqueName, galleryImageName, sharedTo)
                            return
            }

