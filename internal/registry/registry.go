package registry

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/Keshav-Aneja/biz/internal/models"
	"github.com/Keshav-Aneja/biz/internal/utils"
	"github.com/Keshav-Aneja/biz/internal/validators"
)

var REGISTRY_URL = "https://registry.npmjs.org/"

func fetchFromRegistry(path string, pkgDetails interface{}) error {
	resp, err := http.Get(REGISTRY_URL + path)
	if err != nil {
		return fmt.Errorf("error resolving package from registry: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(pkgDetails); err != nil {
		return fmt.Errorf("invalid response received from registry: %w", err)
	}
	return nil
}

func GetPkgDetails(pkgName string) (*models.PkgInfo, error) {
	var pkg models.PkgInfo
	if err := fetchFromRegistry(pkgName, &pkg); err != nil {
		return nil, fmt.Errorf("unable to parse package details: %w", err)
	}
	return &pkg, nil
}

func GetPkgVersionDetails(pkgName string, version string) (*models.PkgVersionInfo, error) {
	path := pkgName + "/" + version
	var pkg models.PkgVersionInfo
	if err := fetchFromRegistry(path, &pkg); err != nil {
		return nil, fmt.Errorf("unable to parse package version details: %w", err)
	}
	return &pkg, nil
}


func ResolvePackage(requestedPkg string) error {
	pkgName, version, needVersionResolution, err := validators.ValidatePkgName(requestedPkg)
	if err != nil {
		return err
	}
	
	var pkg *models.PkgVersionInfo
	if needVersionResolution {
		version, pkg, err = resolveVersionAndTarball(pkgName, version)
		if err != nil {
			return err
		}
	}
		
	if pkg == nil {
		pkg, err = GetPkgVersionDetails(pkgName, version)
		if err != nil {
			return err
		}
	}
	
	tarball := pkg.Dist.Tarball

	if err := utils.DownloadAndExtractPkg(pkg.Name, tarball); err != nil {
		return err
	}

	dependencies, err := utils.ReadDependencies(pkg.Name)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, pkg := range dependencies {
		wg.Add(1)
		go func(pkg string) {
			defer wg.Done()
			if err := ResolvePackage(pkg); err != nil {
				fmt.Printf("Error resolving package %s: %v\n", pkg, err)
			}
		}(pkg)
	}
	wg.Wait()

	return nil
}
