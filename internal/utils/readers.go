package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Keshav-Aneja/biz/internal/constants"
)

type PackageDependencies struct {
	Dependencies map[string]string `json:"dependencies"`
}

func ReadDependencies(package_name string) ([]string, error) {
	content, err := os.ReadDir(filepath.Join(constants.Directories.BIZ_MODULES, package_name))
	if err != nil {
		return nil, fmt.Errorf("unable to read package.json file %w",err);
	}

	for _, val := range content {
		if val.Name() == constants.Directories.PACKAGE_FILE {

		}
	}

	package_file, err := os.ReadFile(filepath.Join(constants.Directories.BIZ_MODULES, package_name, constants.Directories.PACKAGE_FILE))
	if err != nil {
		return nil, err
	}

	var package_content PackageDependencies
	err = json.Unmarshal(package_file, &package_content)
	if err != nil {
		return nil, fmt.Errorf("%s", "unable to parse the module details" + err.Error())
	}

	var dependencies []string
	
	for name, version := range package_content.Dependencies {
		dependencies = append(dependencies, name + "@" + version)
	}

	return dependencies, nil
}
