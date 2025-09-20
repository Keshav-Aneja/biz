package validators

import (
	"fmt"
	"strings"
)

func ValidateModuleName(requestedModule string) (string, bool, error) {
	module := strings.TrimSpace(requestedModule)

	if len(module) == 0 {
		return "", false, fmt.Errorf("invalid module name! Please provide the correct name")
	}

	lastAt := strings.LastIndex(module, "@")

	// No "@" symbol, e.g., "react"
	if lastAt == -1 {
		return module, true, nil
	}

	// Starts with "@", e.g. "@angular/core"
	if lastAt == 0 {
		// This could be just "@" or a scoped package name.
		// If it's just "@", module[1:] will be empty.
		if len(module) > 1 {
			return module, true, nil
		}
		return "", false, fmt.Errorf("invalid module name! Please provide the correct name")
	}

	// "@" is somewhere in the middle or end
	moduleName := module[:lastAt]
	moduleVersion := module[lastAt+1:]

	if len(moduleVersion) == 0 {
		// Case: "module@"
		return moduleName, true, nil
	}

	// Case: "module@version"
	return module, false, nil
}
