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
	package_file, err := os.ReadFile(filepath.Join(constants.Directories.BIZ_MODULES, package_name, constants.Directories.PACKAGE_FILE))
	if err != nil {
		return nil, err
	}

	var package_content PackageDependencies
	err = json.Unmarshal(package_file, &package_content)
	if err != nil {
		return nil, fmt.Errorf("%s", "unable to parse the module details" + err.Error())
	}

	var dependencies = make([]string,0, len(package_content.Dependencies))
	
	for name, version := range package_content.Dependencies {
		dependencies = append(dependencies, name + "@" + version)
	}

	return dependencies, nil
}
