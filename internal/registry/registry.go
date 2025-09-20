package registry

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func GetModuleVersionDetails(moduleName string, latest bool) (*ModuleVersionInfo) {
	path := moduleName
	if latest {
		path += "/latest"
	}
	moduleDetails := fetchFromRegistry(path)

	var module ModuleVersionInfo
	err := json.Unmarshal(moduleDetails, &module)
	if err != nil {
		fmt.Println("Unable to parse the module details", err)
		return nil
	}

	return &module
}
