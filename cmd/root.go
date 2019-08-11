package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:  "wifi-manager",
	Long: "a GOlang cli to change, list and forget wifi networks",
}
