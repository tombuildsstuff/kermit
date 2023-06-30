# IoT Central - API client generation
> see https://aka.ms/autorest
This is the AutoRest configuration file for Azure IoT Central.

## Getting Started
To build the SDK for Azure IoT Central, simply [Install AutoRest](https://aka.ms/autorest/install) and in this folder, run:
> `autorest readme.md`
To see additional help and options, run:
> `autorest --help`
---

## Configuration

### Basic Information

These are the global settings for the IoT Central API.

``` yaml
openapi-type: data-plane
tag: package-2022-10-31-preview
```

### Tag: package-2022-10-31-preview
These settings apply only when `--tag=package-2022-10-31-preview` is specified on the command line.

```yaml $(tag) == 'package-2022-10-31-preview'
input-file:
  - ../../swagger/specification/iotcentral/data-plane/Microsoft.IoTCentral/preview/2022-10-31-preview/iotcentral.json
```

---

# Code Generation

## Go

See configuration in [readme.go.md](./readme.go.md)