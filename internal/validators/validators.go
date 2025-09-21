package validators

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
)

func ValidatePkgName(requestedPkg string) (ValidPkgName string, ValidPkgVersion string, NeedsVersionResolution bool, Error error) {
	pkg := strings.TrimSpace(requestedPkg)

	if len(pkg) == 0 {
		return "", "", false, fmt.Errorf("invalid package name! Please provide the correct name")
	}

	lastAt := strings.LastIndex(pkg, "@")

	// No "@" symbol, e.g., "react"
	if lastAt == -1 {
		return pkg, "latest", false, nil
	}

	// Starts with "@", e.g. "@angular/core"
	if lastAt == 0 {
		// This could be just "@" or a scoped package name.
		// If it's just "@", pkg[1:] will be empty.
		if len(pkg) > 1 {
			return pkg, "latest", false, nil
		}
		return "", "", false, fmt.Errorf("invalid package name! Please provide the correct name")
	}

	// "@" is somewhere in the middle or end
	pkgName := pkg[:lastAt]
	pkgVersion := pkg[lastAt+1:]

	if len(pkgVersion) == 0 {
		// Case: "pkg@"
		return pkgName, "latest", false, nil
	}
	// Case: "pkg@version"
	// Now we need to validate the package version
	_, err := semver.NewVersion(pkgVersion)
	if err != nil {
		//Need to resolve this module using semantic versioning constraints
		return pkgName, pkgVersion, true, nil
	}

	return pkgName, pkgVersion, false, nil
}
