package cmd

import (
	"fmt"
	"os/exec"

	"github.com/omaressameldin/wifi-selector/network"
	"github.com/omaressameldin/wifi-selector/utils"
	"github.com/spf13/cobra"
)

func showInfo(selectedNet network.Network) {
	c := exec.Command(
		"bash",
		"-c",
		fmt.Sprintf("sudo cat /etc/NetworkManager/system-connections/\"%v\"", selectedNet.Filename),
	)
	o, err := c.Output()
	utils.Must(err)
	fmt.Println(string(o))
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "displays info about saved wifi networks",
	Run: func(cmd *cobra.Command, args []string) {
		wifiLists := utils.GetSavedWifis()
		selected := utils.SelectFromList(
			"Select wifi network you want to reveal",
			wifiLists,
			"👀",
		)
		showInfo(wifiLists[selected])
	},
}

func init() {
	RootCmd.AddCommand(infoCmd)
}
