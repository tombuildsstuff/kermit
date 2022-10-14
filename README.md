# Kermit

Kermit provides a stepping-stone from [the Track1 Azure SDK for Go](https://github.com/Azure/azure-sdk-for-go/tree/v67.0.0) to [`hashicorp/go-azure-sdk`](https://github.com/hashicorp/go-azure-sdk).

## Supporting new Versions of Data Plane and Resource Manager Services

See [the documentation in `./config`](./config/README.md) - which is also where the config files live.

## SDKs

SDK's are output into [the `./sdk` directory](./sdk) and can be used like any other Track1 Azure SDK for Go, for example:

```go
package main

import "github.com/tombuildsstuff/kermit/sdk/containerregistry/2022-02-01/containerregistry"

func main() {
    client := containerregistry.NewRegistriesClient("https://management.azure.com", "subscription-id")
	client.Authorizer = ... // see github.com/hashicorp/go-azure-helpers
}
```

## Why is this called Kermit?

Why not.
