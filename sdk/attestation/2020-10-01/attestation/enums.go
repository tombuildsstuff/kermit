package attestation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CertificateModification enumerates the values for certificate modification.
type CertificateModification string

const (
	// CertificateModificationIsAbsent After the operation was performed, the certificate is no longer present
	// in the set of certificates.
	CertificateModificationIsAbsent CertificateModification = "IsAbsent"
	// CertificateModificationIsPresent After the operation was performed, the certificate is in the set of
	// certificates.
	CertificateModificationIsPresent CertificateModification = "IsPresent"
)

// PossibleCertificateModificationValues returns an array of possible values for the CertificateModification const type.
func PossibleCertificateModificationValues() []CertificateModification {
	return []CertificateModification{CertificateModificationIsAbsent, CertificateModificationIsPresent}
}

// DataType enumerates the values for data type.
type DataType string

const (
	// DataTypeBinary The field's content should be treated as binary and not interpreted by MAA.
	DataTypeBinary DataType = "Binary"
	// DataTypeJSON The field's content should be treated as UTF-8 JSON text that may be further interpreted by
	// MAA. Refer to RFC 8259 for a description of JSON serialization standards for interoperability.
	DataTypeJSON DataType = "JSON"
)

// PossibleDataTypeValues returns an array of possible values for the DataType const type.
func PossibleDataTypeValues() []DataType {
	return []DataType{DataTypeBinary, DataTypeJSON}
}

// PolicyModification enumerates the values for policy modification.
type PolicyModification string

const (
	// PolicyModificationRemoved The specified policy object was removed.
	PolicyModificationRemoved PolicyModification = "Removed"
	// PolicyModificationUpdated The specified policy object was updated.
	PolicyModificationUpdated PolicyModification = "Updated"
)

// PossiblePolicyModificationValues returns an array of possible values for the PolicyModification const type.
func PossiblePolicyModificationValues() []PolicyModification {
	return []PolicyModification{PolicyModificationRemoved, PolicyModificationUpdated}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeOpenEnclave OpenEnclave extensions to SGX
	TypeOpenEnclave Type = "OpenEnclave"
	// TypeSgxEnclave Intel Software Guard eXtensions
	TypeSgxEnclave Type = "SgxEnclave"
	// TypeTpm Edge TPM Virtualization Based Security
	TypeTpm Type = "Tpm"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeOpenEnclave, TypeSgxEnclave, TypeTpm}
}
