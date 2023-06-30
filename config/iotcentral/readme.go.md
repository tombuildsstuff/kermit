### Go

These settings apply only when `--go` is specified on the command line.

``` yaml $(go)
license-header: MICROSOFT_MIT_NO_VERSION
directive:
  - remove-model: CommandConfiguration
  - where-model: DeviceCommand
    rename-property:
      from: response
      to: apiResponse
```