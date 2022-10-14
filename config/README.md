## Documentation

Kermit allows the generation of a "Track1" Go SDK using [`Azure/autorest.go`](https://github.com/Azure/autorest.go) where an SDK isn't currently be supported by [`hashicorp/go-azure-sdk`](https://github.com/hashicorp/go-azure-sdk).

This tool supports generating both Data Plane and Resource Manager Services, which are both defined in two distinct files - both of which are output [into the `./sdk` directory within this repository](../sdk).

Data Plane Services can be added to the file `./config/data-plane.hcl` and use the following structure:

```hcl
data_plane "keyvault" "7.3" {
  swagger_tag      = "package-7.3"
  readme_file_path = "../../config/key-vault/readme.md"
}
```

These map to:

* Label `keyvault` - the name of the Service (the directory name used within [the `Azure/azure-rest-api-specs` repository](https://github.com/Azure/azure-rest-api-specs)).
* Label `7.3` - the API Version for this Service.
* Field `swagger_tag` - specifies the Swagger Tag (which must be defined in the associated `readme.md` [example](https://github.com/Azure/azure-rest-api-specs/blob/40f5247a48ff1eec044c8441c422af0628a8a288/specification/elastic/resource-manager/readme.md)) that should be generated.
* Field `readme_file_path` - specifies the relative path to the AutoRest Readme File for this Service, from the `./tools/autowrapper` project.

For Resource Manager:

```hcl
resource_manager "network" "2022-05-01" {
  swagger_tag = "package-2022-05"
  readme_file_path = "../../swagger/specification/network/resource-manager/readme.md"
}
```

These map to:

* Label `network` - the name of the Service (the directory name used within [the `Azure/azure-rest-api-specs` repository](https://github.com/Azure/azure-rest-api-specs)).
* Label `2022-05-01` - the API Version for this Service.
* Field `swagger_tag` - specifies the Swagger Tag (which must be defined in the associated `readme.md` [example](https://github.com/Azure/azure-rest-api-specs/blob/40f5247a48ff1eec044c8441c422af0628a8a288/specification/elastic/resource-manager/readme.md)) that should be generated.
* Field `readme_file_path` - specifies the relative path to the AutoRest Readme File for this Service, from the `./tools/autowrapper` project.

The resulting output is run by the `./tools/autowrapper` tool and sent as an auto-PR to this repository - where the generated SDKs get added to [`./sdk` in the repository root](../sdk).

Any issues? Feel free to add an issue.
