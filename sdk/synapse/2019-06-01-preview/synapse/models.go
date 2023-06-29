package synapse

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
)

// The package's fully qualified name.
const fqdn = "home/runner/work/kermit/kermit/sdk/synapse/2019-06-01-preview/synapse"

// ManagedPrivateEndpoint managed private endpoint
type ManagedPrivateEndpoint struct {
	autorest.Response `json:"-"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
	// Properties - Managed private endpoint properties
	Properties *ManagedPrivateEndpointProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for ManagedPrivateEndpoint.
func (mpe ManagedPrivateEndpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mpe.Properties != nil {
		objectMap["properties"] = mpe.Properties
	}
	return json.Marshal(objectMap)
}

// ManagedPrivateEndpointConnectionState the connection state of a managed private endpoint
type ManagedPrivateEndpointConnectionState struct {
	// Status - READ-ONLY; The approval status
	Status *string `json:"status,omitempty"`
	// Description - The managed private endpoint description
	Description *string `json:"description,omitempty"`
	// ActionsRequired - The actions required on the managed private endpoint
	ActionsRequired *string `json:"actionsRequired,omitempty"`
}

// MarshalJSON is the custom marshaler for ManagedPrivateEndpointConnectionState.
func (mpecs ManagedPrivateEndpointConnectionState) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mpecs.Description != nil {
		objectMap["description"] = mpecs.Description
	}
	if mpecs.ActionsRequired != nil {
		objectMap["actionsRequired"] = mpecs.ActionsRequired
	}
	return json.Marshal(objectMap)
}

// ManagedPrivateEndpointListResponse a list of managed private endpoints
type ManagedPrivateEndpointListResponse struct {
	autorest.Response `json:"-"`
	// Value - List of managed private endpoints
	Value *[]ManagedPrivateEndpoint `json:"value,omitempty"`
	// NextLink - READ-ONLY; The link to the next page of results, if any remaining results exist.
	NextLink *string `json:"nextLink,omitempty"`
}

// MarshalJSON is the custom marshaler for ManagedPrivateEndpointListResponse.
func (mpelr ManagedPrivateEndpointListResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mpelr.Value != nil {
		objectMap["value"] = mpelr.Value
	}
	return json.Marshal(objectMap)
}

// ManagedPrivateEndpointListResponseIterator provides access to a complete listing of
// ManagedPrivateEndpoint values.
type ManagedPrivateEndpointListResponseIterator struct {
	i    int
	page ManagedPrivateEndpointListResponsePage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ManagedPrivateEndpointListResponseIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedPrivateEndpointListResponseIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ManagedPrivateEndpointListResponseIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ManagedPrivateEndpointListResponseIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ManagedPrivateEndpointListResponseIterator) Response() ManagedPrivateEndpointListResponse {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ManagedPrivateEndpointListResponseIterator) Value() ManagedPrivateEndpoint {
	if !iter.page.NotDone() {
		return ManagedPrivateEndpoint{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the ManagedPrivateEndpointListResponseIterator type.
func NewManagedPrivateEndpointListResponseIterator(page ManagedPrivateEndpointListResponsePage) ManagedPrivateEndpointListResponseIterator {
	return ManagedPrivateEndpointListResponseIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (mpelr ManagedPrivateEndpointListResponse) IsEmpty() bool {
	return mpelr.Value == nil || len(*mpelr.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (mpelr ManagedPrivateEndpointListResponse) hasNextLink() bool {
	return mpelr.NextLink != nil && len(*mpelr.NextLink) != 0
}

// managedPrivateEndpointListResponsePreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (mpelr ManagedPrivateEndpointListResponse) managedPrivateEndpointListResponsePreparer(ctx context.Context) (*http.Request, error) {
	if !mpelr.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(mpelr.NextLink)))
}

// ManagedPrivateEndpointListResponsePage contains a page of ManagedPrivateEndpoint values.
type ManagedPrivateEndpointListResponsePage struct {
	fn    func(context.Context, ManagedPrivateEndpointListResponse) (ManagedPrivateEndpointListResponse, error)
	mpelr ManagedPrivateEndpointListResponse
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ManagedPrivateEndpointListResponsePage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ManagedPrivateEndpointListResponsePage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.mpelr)
		if err != nil {
			return err
		}
		page.mpelr = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ManagedPrivateEndpointListResponsePage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ManagedPrivateEndpointListResponsePage) NotDone() bool {
	return !page.mpelr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ManagedPrivateEndpointListResponsePage) Response() ManagedPrivateEndpointListResponse {
	return page.mpelr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ManagedPrivateEndpointListResponsePage) Values() []ManagedPrivateEndpoint {
	if page.mpelr.IsEmpty() {
		return nil
	}
	return *page.mpelr.Value
}

// Creates a new instance of the ManagedPrivateEndpointListResponsePage type.
func NewManagedPrivateEndpointListResponsePage(cur ManagedPrivateEndpointListResponse, getNextPage func(context.Context, ManagedPrivateEndpointListResponse) (ManagedPrivateEndpointListResponse, error)) ManagedPrivateEndpointListResponsePage {
	return ManagedPrivateEndpointListResponsePage{
		fn:    getNextPage,
		mpelr: cur,
	}
}

// ManagedPrivateEndpointProperties properties of a managed private endpoint
type ManagedPrivateEndpointProperties struct {
	// PrivateLinkResourceID - The ARM resource ID of the resource to which the managed private endpoint is created
	PrivateLinkResourceID *string `json:"privateLinkResourceId,omitempty"`
	// GroupID - The groupId to which the managed private endpoint is created
	GroupID *string `json:"groupId,omitempty"`
	// ProvisioningState - READ-ONLY; The managed private endpoint provisioning state
	ProvisioningState *string `json:"provisioningState,omitempty"`
	// ConnectionState - The managed private endpoint connection state
	ConnectionState *ManagedPrivateEndpointConnectionState `json:"connectionState,omitempty"`
	// IsReserved - READ-ONLY; Denotes whether the managed private endpoint is reserved
	IsReserved *bool `json:"isReserved,omitempty"`
}

// MarshalJSON is the custom marshaler for ManagedPrivateEndpointProperties.
func (mpep ManagedPrivateEndpointProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if mpep.PrivateLinkResourceID != nil {
		objectMap["privateLinkResourceId"] = mpep.PrivateLinkResourceID
	}
	if mpep.GroupID != nil {
		objectMap["groupId"] = mpep.GroupID
	}
	if mpep.ConnectionState != nil {
		objectMap["connectionState"] = mpep.ConnectionState
	}
	return json.Marshal(objectMap)
}
