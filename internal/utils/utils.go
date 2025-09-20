package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Keshav-Aneja/biz/internal/constants"
)

func Download(moduleName string, path string) error {
	tempDirectory := constants.Directories.TEMPORARY
	err := os.MkdirAll(tempDirectory, os.FileMode(constants.Permissions.DIRECTORY))
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	file, err := http.Get(path)
	if err != nil {
		return err
	}
	defer file.Body.Close()

	out, err := os.Create(filepath.Join(tempDirectory, moduleName))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, file.Body)
	return err
}

func setupDirectory(dest string) error {
	tempDirectory := constants.Directories.TEMPORARY
	err := os.MkdirAll(tempDirectory, os.FileMode(constants.Permissions.DIRECTORY))
	if err != nil && !os.IsExist(err) {
		return err
	}

	err = os.MkdirAll(dest, os.FileMode(constants.Permissions.DIRECTORY))
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func removeDirectory(dest string) error {
	err := os.Remove(dest)
	if err != nil {
		return err
	}
	// If no more tarballs are there then we remove the directory
	// dir, err := os.ReadDir(constants.Directories.TEMPORARY)
	// if err != nil {
	// 	return err
	// }

	// if len(dir) == 0 {
	// 	return os.Remove(constants.Directories.TEMPORARY)
	// }
	return nil
}

/**
*
* reads the uncompressed tarballs using gzip and tar
* starts extracting all the files and folders from the specific package
*
*/
func Extract(moduleName string) error {
	tgzPath := constants.Directories.TEMPORARY + "/" + moduleName
	dest := constants.Directories.BIZ_MODULES + "/" + moduleName
	//Setup the directories for extracting
	if err := setupDirectory(dest); err != nil {
		return fmt.Errorf("error setting up directories %w", err)
	}

	file, err := os.Open(tgzPath)
	if err != nil {
		return err
	}
	defer file.Close()

	uncompressedStream, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("extracting tarGz failed: %w", err)
	}
	defer uncompressedStream.Close()

	tarReader := tar.NewReader(uncompressedStream)
	createdDirs := make(map[string]bool)

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("extracting tarGz: Next() failed %w", err)
		}

		parts := strings.Split(header.Name, "/")
		target := filepath.Join(dest, filepath.Join(parts[1:]...))

		switch header.Typeflag {
		case tar.TypeDir:
			if !createdDirs[target] {
				if err := os.MkdirAll(target, 0755); err != nil {
					return fmt.Errorf("mkdir %s failed: %w", target, err)
				}
				createdDirs[target] = true
			}
		case tar.TypeReg:
			dir := filepath.Dir(target)
			if !createdDirs[dir] {
				if err := os.MkdirAll(dir, 0755); err != nil {
					return fmt.Errorf("mkdir %s failed: %w", dir, err)
				}
				createdDirs[dir] = true
			}

			outFile, err := os.Create(target)
			if err != nil {
				return fmt.Errorf("create file %s failed: %w", target, err)
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return fmt.Errorf("copy file %s failed: %w", target, err)
			}
			outFile.Close()
		default:
			// skip special files (symlinks, etc.)
		}
	}

	defer removeDirectory(tgzPath)
	return nil
}
