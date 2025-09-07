package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



var rootCmd = &cobra.Command{
	Use:   "biz",
	Short: "A fast, minimal JavaScript package manager",
	Long: `Biz 🐼 is a lightweight JavaScript package manager developed in Go and designed to simplify 
dependency management for your projects. With Biz, you can quickly add, remove, 
and manage packages, while keeping your project manifest and lockfile up to date.

Features include:
  - Fast installation of packages from the npm registry
  - Automatic dependency resolution and lockfile updates
  - Support for development and production dependencies
  - Easy-to-use CLI commands like 'add', 'remove', and 'install'

Usage Examples:

  # Initialize a new project
  biz init

  # Add a package to your dependencies
  biz get lodash

  # Remove a package from your project
  biz remove lodash

  # Install all dependencies listed in the manifest
  biz install

Biz empowers developers to manage dependencies efficiently without the overhead 
of larger package managers, making project setup and maintenance faster and simpler.`,
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.biz.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


