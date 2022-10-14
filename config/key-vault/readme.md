# KeyVault

> see https://aka.ms/autorest

This is the AutoRest configuration file for KeyVault.

---

## Getting Started

To build the SDK for KeyVault, simply [Install AutoRest](https://aka.ms/autorest/install) and in this folder, run:

> `autorest`

To see additional help and options, run:

> `autorest --help`

---

## Configuration

### Basic Information

These are the global settings for the KeyVault API.

``` yaml
openapi-type: data-plane
tag: package-7.3
```


### Tag: package-7.3

These settings apply only when `--tag=package-7.3` is specified on the command line.

```yaml $(tag) == 'package-7.3'
input-file:
  - ../../swagger/specification/keyvault/data-plane/Microsoft.KeyVault/stable/7.3/backuprestore.json
  - ../../swagger/specification/keyvault/data-plane/Microsoft.KeyVault/stable/7.3/certificates.json
  - ../../swagger/specification/keyvault/data-plane/Microsoft.KeyVault/stable/7.3/common.json
  - ../../swagger/specification/keyvault/data-plane/Microsoft.KeyVault/stable/7.3/keys.json
  - ../../swagger/specification/keyvault/data-plane/Microsoft.KeyVault/stable/7.3/rbac.json
  - ../../swagger/specification/keyvault/data-plane/Microsoft.KeyVault/stable/7.3/secrets.json
  - ../../swagger/specification/keyvault/data-plane/Microsoft.KeyVault/stable/7.3/securitydomain.json
  - ../../swagger/specification/keyvault/data-plane/Microsoft.KeyVault/stable/7.3/storage.json
```
---

# Code Generation

## Go

See configuration in [readme.go.md](./readme.go.md)

## Suppression

``` yaml
directive:
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.CertificateOperation.properties.cancellation_requested
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.CertificateOperation.properties.status_details
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.CertificateOperation.properties.request_id
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.CertificatePolicy.properties.key_props
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.CertificatePolicy.properties.secret_props
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.CertificatePolicy.properties.x509_props
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.CertificatePolicy.properties.lifetime_actions
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.X509CertificateProperties.properties.key_usage
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.X509CertificateProperties.properties.validity_months
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.IssuerParameters.properties.cert_transparency
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.Action.properties.action_type
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.Trigger.properties.lifetime_percentage
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.Trigger.properties.days_before_expiry
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.SubjectAlternativeNames.properties.dns_names
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.IssuerBundle.properties.org_details
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.IssuerCredentials.properties.account_id
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.OrganizationDetails.properties.admin_details
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.AdministratorDetails.properties.first_name
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.AdministratorDetails.properties.last_name
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.CertificateIssuerSetParameters.properties.org_details
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.CertificateIssuerUpdateParameters.properties.org_details
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: certificates.json
    where: $.definitions.CertificateOperationUpdateParameter.properties.cancellation_requested
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: keys.json
    where: $.definitions.KeyProperties.properties.key_size
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: keys.json
    where: $.definitions.KeyProperties.properties.reuse_key
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: keys.json
    where: $.definitions.JsonWebKey.properties.key_ops
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: keys.json
    where: $.definitions.JsonWebKey.properties.key_hsm
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: keys.json
    where: $.definitions.KeyBundle.properties.release_policy
    reason: Consistency with other properties.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: keys.json
    where: $.definitions.KeyCreateParameters.properties.key_size
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: keys.json
    where: $.definitions.KeyCreateParameters.properties.public_exponent
    reason: Consistency with other properties.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: keys.json
    where: $.definitions.KeyCreateParameters.properties.release_policy
    reason: Consistency with other properties.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: keys.json
    where: $.definitions.KeyCreateParameters.properties.key_ops
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: keys.json
    where: $.definitions.KeyImportParameters.properties.Hsm
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: keys.json
    where: $.definitions.KeyImportParameters.properties.release_policy
    reason: Consistency with other properties.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: keys.json
    where: $.definitions.KeyUpdateParameters.properties.key_ops
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: keys.json
    where: $.definitions.KeyUpdateParameters.properties.release_policy
    reason: Consistency with other properties.
  - suppress: MISSING_REQUIRED_PARAMETER
    from: certificates.json
    where: $..parameters[?(@.name=='vaultBaseUrl')]
    reason: Suppress an invalid error caused by a bug in the linter.
  - suppress: MISSING_REQUIRED_PARAMETER
    from: keys.json
    where: $..parameters[?(@.name=='vaultBaseUrl')]
    reason: Suppress an invalid error caused by a bug in the linter.
  - suppress: MISSING_REQUIRED_PARAMETER
    from: secrets.json
    where: $..parameters[?(@.name=='vaultBaseUrl')]
    reason: Suppress an invalid error caused by a bug in the linter.
  - suppress: MISSING_REQUIRED_PARAMETER
    from: storage.json
    reason: Suppress an invalid error caused by a bug in the linter.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: securitydomain.json
    where: $.definitions.TransferKey.properties.transfer_key
    reason: Merely refactored existing definitions into new files.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: securitydomain.json
    where: $.definitions.UploadPendingResponse.properties.status_details
    reason: Consistency with other properties.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: securitydomain.json
    where: $.definitions.SecurityDomainOperationStatus.properties.status_details
    reason: Consistency with other properties.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: securitydomain.json
    where: $.definitions.SecurityDomainJsonWebKey.properties.key_ops
    reason: Consistency with other properties.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: securitydomain.json
    where: $.definitions.SecurityDomainJsonWebKey.properties["x5t#S256"]
    reason: Consistency with other properties.
  - suppress: DefinitionsPropertiesNamesCamelCase
    from: securitydomain.json
    where: $.definitions.TransferKey.properties.key_format
    reason: Consistency with other properties
  - suppress: DOUBLE_FORWARD_SLASHES_IN_URL
    from: rbac.json
    reason: / is a valid scope in this scenario.
  - suppress: OBJECT_MISSING_REQUIRED_PROPERTY
    from: rbac.json
    where: $..parameters[?(@.name=='scope')]
    reason: Suppress an invalid error caused by a bug in the linter.
```
