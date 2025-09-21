package registry

import (
	"fmt"

	"github.com/Keshav-Aneja/biz/internal/models"
	"github.com/Masterminds/semver"
)

func resolveVersionAndTarball(pkgName string, versionRange string) (string, *models.PkgVersionInfo, error)  {
	if versionRange == "latest" {
		// [NOTE] : For now let's return this latest version from here, afterwards we need to check for the node version as well
		// And download according to the node version
		return "latest", nil, nil
	}

	pkgDetails, err := GetPkgDetails(pkgName)
	if err != nil {
		return "", nil, err
	}

	constraint, err := semver.NewConstraint(versionRange)
	if err != nil {
		return "", nil, fmt.Errorf("invalid version constraint: %w", err)
	}

	var resolvedPackage models.PkgVersionInfo
	var bestVersion *semver.Version
	for v := range pkgDetails.Versions {
		sv, err := semver.NewVersion(v)
		if err != nil {
			continue
		}
		if constraint.Check(sv) {
			if bestVersion == nil || sv.GreaterThan(bestVersion) {
				bestVersion = sv
				resolvedPackage = pkgDetails.Versions[v]
			}
		}
	}

	if bestVersion == nil {
		return "", nil, fmt.Errorf("no matching version found for %s", versionRange)
	}
	return bestVersion.String(), &resolvedPackage, nil
}
