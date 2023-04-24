// Package attestation implements the Azure ARM Attestation service API version 2022-08-01.
//
// Describes the interface for the per-tenant enclave service.
package attestation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
)


// BaseClient is the base client for Attestation.
type BaseClient struct {
    autorest.Client
}

// New creates an instance of the BaseClient client.
func New()BaseClient {
    return NewWithoutDefaults()
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults() BaseClient {
    return BaseClient{
        Client: autorest.NewClientWithUserAgent(UserAgent()),
    }
}
