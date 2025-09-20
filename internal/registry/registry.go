package registry

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/Keshav-Aneja/biz/internal/utils"
	"github.com/Keshav-Aneja/biz/internal/validators"
	"github.com/Masterminds/semver"
)

var REGISTRY_URL = "https://registry.npmjs.org/"

func fetchFromRegistry(path string, moduleDetails interface{}) error {
	resp, err := http.Get(REGISTRY_URL + path)
	if err != nil {
		return fmt.Errorf("error resolving package from registry: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(moduleDetails); err != nil {
		return fmt.Errorf("invalid response received from registry: %w", err)
	}

	return nil
}

func GetModuleDetails(moduleName string) (*ModuleInfo, error) {
	var module ModuleInfo
	if err := fetchFromRegistry(moduleName, &module); err != nil {
		return nil, fmt.Errorf("unable to parse module details: %w", err)
	}
	return &module, nil
}

func GetModuleVersionDetails(moduleName string, version string) (*ModuleVersionInfo, error) {
	path := moduleName + "/" + version
	var module ModuleVersionInfo
	if err := fetchFromRegistry(path, &module); err != nil {
		return nil, fmt.Errorf("unable to parse module version details: %w", err)
	}
	return &module, nil
}

func resolveVersion(moduleName string, versionRange string) (string, error) {
	if versionRange == "latest" {
		// fmt.Print(moduleDetails.DistTags.Latest)
		// return moduleDetails.DistTags.Latest, nil
		// [NOTE] : For now let's return this latest version from here, afterwards we need to check for the node version as well
		// And download according to the node version
		return "latest", nil
	}

	moduleDetails, err := GetModuleDetails(moduleName)
	if err != nil {
		return "", err
	}

	constraint, err := semver.NewConstraint(versionRange)
	if err != nil {
		return "", fmt.Errorf("invalid version constraint: %w", err)
	}

	var bestVersion *semver.Version
	for v := range moduleDetails.Versions {
		sv, err := semver.NewVersion(v)
		if err != nil {
			continue
		}
		if constraint.Check(sv) {
			if bestVersion == nil || sv.GreaterThan(bestVersion) {
				bestVersion = sv
			}
		}
	}

	if bestVersion == nil {
		return "", fmt.Errorf("no matching version found for %s", versionRange)
	}
	return bestVersion.String(), nil
}

func ResolveModule(requestedModule string) error {
	fmt.Println(requestedModule)
	moduleName, versionRange, err := validators.ValidateModuleName(requestedModule)
	if err != nil {
		return err
	}

	version, err := resolveVersion(moduleName, versionRange)
	if err != nil {
		return err
	}

	module, err := GetModuleVersionDetails(moduleName, version)
	if err != nil {
		return err
	}
	tarball := module.Dist.Tarball

	if err := utils.Download(module.Name, tarball); err != nil {
		return fmt.Errorf("error downloading the module: %w", err)
	}

	if err := utils.Extract(module.Name); err != nil {
		return fmt.Errorf("error extracting the module: %w", err)
	}

	dependencies, err := utils.ReadDependencies(module.Name)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, module := range dependencies {
		wg.Add(1)
		go func(module string) {
			defer wg.Done()
			if err := ResolveModule(module); err != nil {
				fmt.Printf("Error resolving module %s: %v\n", module, err)
			}
		}(module)
	}
	wg.Wait()

	return nil
}
