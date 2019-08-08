package cmd

import (
	"os/exec"

	"github.com/omaressameldin/wifi-selector/utils"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restart nework manager",
	Run: func(cmd *cobra.Command, args []string) {
		command := "sudo service NetworkManager restart	"
		c := exec.Command("bash", "-c", command)
		_, err := c.Output()
		utils.Must(err)
	},
}

func init() {
	RootCmd.AddCommand(restartCmd)
}
