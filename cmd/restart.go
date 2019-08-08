package cmd

import (
	"fmt"
	"os/exec"

	spinner "github.com/omaressameldin/wifi-selector/Spinner"
	"github.com/omaressameldin/wifi-selector/utils"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restart nework manager",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.Spinner{
			Shape: 14,
		}
		s.StartSpinner("Restarting")

		command := "sudo service NetworkManager restart"
		c := exec.Command("bash", "-c", command)
		_, err := c.Output()
		s.StopSpinner()
		utils.Must(err)
		fmt.Println("✔️  Restarted Successfully")
	},
}

func init() {
	RootCmd.AddCommand(restartCmd)
}
