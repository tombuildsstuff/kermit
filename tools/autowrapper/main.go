package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

const autoRestGoVersion = "2.1.188"

func main() {
	log.Printf("Launched AutoWrapper..")
	if err := run(); err != nil {
		log.Fatalf("running: %+v", err)
		return
	}
}

func run() error {
	dataPlaneServices, err := parseDataPlaneServicesFromConfig("../../config/data-plane.hcl")
	if err != nil {
		return fmt.Errorf("loading the Data Plane config: %+v", err)
	}

	log.Printf("Loading Resource Manager Services from the Config..")
	resourceManagerServices, err := parseResourceManagerServicesFromConfig("../../config/resource-manager.hcl")
	if err != nil {
		return fmt.Errorf("loading the Resource Manager config: %+v", err)
	}

	log.Printf("Generating Data Plane SDKs..")
	for _, dataPlane := range *dataPlaneServices {
		log.Printf("[DEBUG] Generating Data Plane Service %q / API Version %q / Swagger Tag %q / Readme %q", dataPlane.ServiceName, dataPlane.ApiVersion, dataPlane.SwaggerTag, dataPlane.ReadmeFilePath)
		err = runAutoRestForService(dataPlane.ServiceName, dataPlane.ApiVersion, dataPlane.SwaggerTag, dataPlane.ReadmeFilePath, false)
		if err != nil {
			log.Printf("Error running generator: %+v", err)
		}
	}
	log.Printf("Finished Generating Data Plane SDKs.")

	log.Printf("Generating Resource Manager SDKs..")
	for _, resourceManager := range *resourceManagerServices {
		log.Printf("[DEBUG] Generating Resource Manager Service %q / API Version %q / Swagger Tag %q / Readme %q", resourceManager.ServiceName, resourceManager.ApiVersion, resourceManager.SwaggerTag, resourceManager.ReadmeFilePath)
		err = runAutoRestForService(resourceManager.ServiceName, resourceManager.ApiVersion, resourceManager.SwaggerTag, resourceManager.ReadmeFilePath, true)
		if err != nil {
			log.Printf("Error running generator: %+v", err)
		}
	}
	log.Printf("Finished Generating Resource Manager SDKs.")
	return nil
}

func runAutoRestForService(serviceName, apiVersion, tag, readmeFilePath string, isResourceManager bool) error {
	// autorest --use=@microsoft.azure/autorest.go@2.1.187 --tag=package-2022-02-preview --go-sdk-folder=./azure-sdk-for-go --go --verbose --use-onever --version=2.0.4421 --go.license-header=MICROSOFT_MIT_NO_VERSION --enum-prefix ./azure-rest-api-specs/specification/containerregistry/resource-manager/readme.md --output-folder=./example-dir
	outputDirectory := fmt.Sprintf("../../sdk/%[1]s/%[2]s/%[1]s", serviceName, apiVersion)
	log.Printf("[DEBUG] Generating Service %q / Tag %q into %q..", serviceName, tag, outputDirectory)
	typeArg := "--azure-arm"
	if !isResourceManager {
		typeArg = "--openapi-type=data-plane"
	}
	args := []string{
		fmt.Sprintf("--use=@microsoft.azure/autorest.go@%s", autoRestGoVersion), // autorest.go plugin version
		fmt.Sprintf("--tag=%s", tag),
		"--go",
		"--verbose",
		"--use-onever",
		"--version=2.0.4421", // the version of autorest itself
		"--go.license-header=MICROSOFT_MIT_NO_VERSION",
		"--enum-prefix",
		fmt.Sprintf("--namespace=%s", strings.ToLower(serviceName)),
		typeArg,
		fmt.Sprintf("--output-folder=%s", outputDirectory),
		readmeFilePath,
	}
	log.Printf("[DEBUG] Launching `autorest %s`", strings.Join(args, " "))
	cmd := exec.Command("autorest", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	if err := cmd.Run(); err != nil {
		exitCode := "unknown"
		if exit, ok := err.(*exec.ExitError); ok {
			exitCode = fmt.Sprintf("%d", exit.ExitCode())
		}
		return fmt.Errorf("running generator (Command: autorest %q) had exit code %q: %+v", exitCode, strings.Join(args, " "), err)
	}
	log.Printf("[DEBUG] Generated SDK for %q / Tag %q", serviceName, tag)

	absPath, err := filepath.Abs(outputDirectory)
	if err != nil {
		return fmt.Errorf("determing absolute path for %q: %+v", outputDirectory)
	}
	log.Printf("[DEBUG] Removing Interfaces from the SDK for %q / Tag %q from %q", serviceName, tag, absPath)
	os.RemoveAll(fmt.Sprintf("%s/%sapi", outputDirectory, serviceName))
	log.Printf("[DEBUG] Removed Interfaces from the SDK for Service %q / Tag %q", serviceName, tag)
	return nil
}

type dataPlaneService struct {
	ServiceName    string `hcl:"serviceName,label"`
	ApiVersion     string `hcl:"apiVersion,label"`
	SwaggerTag     string `hcl:"swagger_tag"`
	ReadmeFilePath string `hcl:"readme_file_path"`
}

func parseDataPlaneServicesFromConfig(dataPlaneFilePath string) (*[]dataPlaneService, error) {
	type Config struct {
		DataPlane []dataPlaneService `hcl:"data_plane,block"`
	}

	var config Config
	err := hclsimple.DecodeFile(dataPlaneFilePath, nil, &config)
	if err != nil {
		return nil, fmt.Errorf("parsing: %+v", err)
	}

	return &config.DataPlane, nil
}

type resourceManagerService struct {
	ServiceName    string `hcl:"serviceName,label"`
	ApiVersion     string `hcl:"apiVersion,label"`
	SwaggerTag     string `hcl:"swagger_tag"`
	ReadmeFilePath string `hcl:"readme_file_path"`
}

func parseResourceManagerServicesFromConfig(resourceManagerFile string) (*[]resourceManagerService, error) {
	type Config struct {
		Resources []resourceManagerService `hcl:"resource_manager,block"`
	}

	var config Config
	err := hclsimple.DecodeFile(resourceManagerFile, nil, &config)
	if err != nil {
		return nil, fmt.Errorf("parsing: %+v", err)
	}

	return &config.Resources, nil
}
