package cmd

import (
	"github.com/Keshav-Aneja/biz/internal/printer"
	"github.com/Keshav-Aneja/biz/internal/registry"
	"github.com/Keshav-Aneja/biz/internal/validators"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [package_name]",
	Short: "Installs a package and gets it to your project dependencies",
	Long: `The 'get' command installs the specified package into your project's 
node_modules folder and updates your manifest file (biz.json) with the dependency.

Usage Examples:

  # Install a package and get it to dependencies
  biz get lodash

  # Install a specific version of a package
  biz get react@18.2.0

  # Install a package as a development dependency
  biz get --dev typescript

This command handles downloading the package from the registry, resolving 
dependencies, and updating your lockfile to ensure consistent installations.`,
	Run: func(cmd *cobra.Command, args []string) {
		printer.Gradient("Biz - Package manager 🐼")
		
		moduleName, latest, err := validators.ValidateModuleName(args[0])
		if err != nil {
			printer.Error(err.Error())
		}

		module := registry.GetModuleVersionDetails(moduleName, latest)


		printer.Gradient("Package acquired successfully " + module.Name + " - " + module.Description)
	},
}


func init() {
	rootCmd.AddCommand(getCmd)
}
