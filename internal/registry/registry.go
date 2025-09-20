package registry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Keshav-Aneja/biz/internal/utils"
	"github.com/Keshav-Aneja/biz/internal/validators"
)

var REGISTRY_URL = "https://registry.npmjs.org/"

func fetchFromRegistry(path string) ([]byte) {
	resp, err := http.Get(REGISTRY_URL + path)
	if err != nil {
		fmt.Println("Error resolving package from registry", err)
		return nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Invalid response recieved from registry", err)
		return nil
	}

	return body
}


func GetModuleDetails(moduleName string) (*ModuleInfo) {
	moduleDetails := fetchFromRegistry(moduleName)

	var module ModuleInfo
	err := json.Unmarshal(moduleDetails, &module)
	if err != nil {
		fmt.Println("Unable to parse the module details", err)
		return nil
	}

	return &module
}

func GetModuleVersionDetails(moduleName string, version string) (*ModuleVersionInfo) {
	path := moduleName + "/" + version
	
	moduleDetails := fetchFromRegistry(path)

	var module ModuleVersionInfo
	err := json.Unmarshal(moduleDetails, &module)
	if err != nil {
		fmt.Println("Unable to parse the module details", err)
		return nil
	}

	return &module
}


func ResolveModule(requestedModule string) error {
	fmt.Println(requestedModule)
	moduleName, version, err := validators.ValidateModuleName(requestedModule)
	if err != nil {
		return err;
	}

	module := GetModuleVersionDetails(moduleName, version)
	tarball := module.Dist.Tarball

	err = utils.Download(module.Name, tarball)
	if err != nil {
		return fmt.Errorf("%s", "Error downloading the module"  + err.Error())
	}

	err = utils.Extract(module.Name)
	if err != nil {
		return fmt.Errorf("%s", "Error extracting the module " + err.Error())
	}

	dependencies, err := utils.ReadDependencies(module.Name)
	if err != nil {
		return err;
	}


	for _, module := range dependencies {
		// fmt.Println(module)
		ResolveModule(module)
	}

	return nil
}
