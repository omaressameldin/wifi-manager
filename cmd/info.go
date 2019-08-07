package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/omaressameldin/wifi-selector/network"
	"github.com/omaressameldin/wifi-selector/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func parseInfoFlags(flags *pflag.FlagSet) string {
	var infoToSho []string

	p, _ := flags.GetBool("password")
	if p {
		infoToSho = append(infoToSho, "'\\[wifi-security\\]'")
	}

	m, _ := flags.GetBool("main")
	if m {
		infoToSho = append(infoToSho, "'\\[wifi\\]'")
	}

	c, _ := flags.GetBool("connection")
	if c {
		infoToSho = append(infoToSho, "'\\[connection\\]'")
	}

	if len(infoToSho) > 0 {
		return fmt.Sprintf("| grep -A 4 -E %v", strings.Join(infoToSho, "\\|"))
	} else {
		return ""
	}
}

func showInfo(selectedNet network.SavedNetwork, extraOptions string) {
	c := exec.Command(
		"bash",
		"-c",
		fmt.Sprintf(
			"sudo cat /etc/NetworkManager/system-connections/\"%v\" %v",
			selectedNet.Filename,
			extraOptions,
		),
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
			network.ListNames(wifiLists),
			"👀",
			true,
		)
		flags := cmd.Flags()

		showInfo(wifiLists[selected], parseInfoFlags(flags))
	},
}

func init() {
	RootCmd.AddCommand(infoCmd)
	infoCmd.Flags().BoolP("password", "p", false, "show only networks' password")
	infoCmd.Flags().BoolP("main", "m", false, "show only networks' main info")
	infoCmd.Flags().BoolP("connection", "c", false, "show only networks' connection")
}
