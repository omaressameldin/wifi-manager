package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/omaressameldin/wifi-selector/utils"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

func connectToNetwork(networkName string) {
	nTrimed := strings.TrimSpace(networkName)
	c := exec.Command(
		"bash",
		"-c",
		fmt.Sprintf("nmcli dev wifi con --ask \"%v\"", nTrimed),
	)

	stdin, err := c.StdinPipe()
	utils.Must(err)
	defer stdin.Close() // the doc says subProcess.Wait will close it, but I'm not sure, so I kept this line

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err = c.Start()
	utils.Must(err)

	password, err := terminal.ReadPassword(1)
	utils.Must(err)
	io.WriteString(stdin, fmt.Sprintf("%v\n", string(password)))

	err = c.Wait()
	utils.Must(err)
}

var conCmd = &cobra.Command{
	Use:       "con",
	Short:     "Connect to a wifi network",
	ValidArgs: []string{},

	Run: func(cmd *cobra.Command, args []string) {
		networks := utils.GetAvailableNetworks("-f ssid")[1:]            // skip headers
		networksToShow := utils.GetAvailableNetworks("-f ssid,bars")[1:] // skip headers
		selected := utils.SelectFromList(
			"Select netowrk to connect",
			networksToShow,
			"👻",
			true,
		)
		connectToNetwork(networks[selected])
	},
}

func init() {
	RootCmd.AddCommand(conCmd)
}
