package cmd

import (
	"github.com/Keshav-Aneja/biz/internal/printer"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project with a biz.json manifest",
	Long: `The 'init' command sets up a new project by creating a biz.json file 
in the current directory. This manifest file is used by Biz to manage your 
project's dependencies, scripts, and metadata.

You will be prompted to provide details such as:
  - Project name
  - Version
  - Description
  - Entry point (e.g., index.js)
  - Author and license

If you prefer, you can skip prompts to quickly generate a default biz.json.

Usage Examples:

  # Start an interactive setup to create a biz.json
  biz init

  # Quickly generate a default biz.json without prompts
  biz init --yes`,
	Run: func(cmd *cobra.Command, args []string) {
		printer.Gradient("Biz - Package manager 🐼")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
