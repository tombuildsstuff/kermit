## Go

These settings apply only when `--go` is specified on the command line.

``` yaml $(go)
go:
  license-header: MICROSOFT_MIT_NO_VERSION
  namespace: keyvault
  clear-output-folder: true
```

### Go multi-api

``` yaml $(go) && $(multiapi)
batch:
  - tag: package-preview-7.5-preview.1
  - tag: package-7.4
```

### Tag: package-preview-7.5-preview.1 and go

These settings apply only when `--tag=package-preview-7.5-preview.1 --go` is specified on the command line.
Please also specify `--go-sdk-folder=<path to the root directory of your azure-sdk-for-go clone>`.

``` yaml $(tag) == 'package-preview-7.5-preview.1' && $(go)
output-folder: $(go-sdk-folder)/services/$(namespace)/v7.5-preview.1/$(namespace)
```

### Tag: package-7.4 and go

These settings apply only when `--tag=package-7.4 --go` is specified on the command line.
Please also specify `--go-sdk-folder=<path to the root directory of your azure-sdk-for-go clone>`.

``` yaml $(tag) == 'package-7.4' && $(go)
output-folder: $(go-sdk-folder)/services/$(namespace)/v7.4/$(namespace)
```
