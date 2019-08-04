package cmd

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/omaressameldin/wifi-selector/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func parseFlags(flags *pflag.FlagSet) string {
	var cmdColumns []string
	n, _ := flags.GetBool("name")
	if n {
		cmdColumns = append(cmdColumns, "ssid")
	}
	m, _ := flags.GetBool("mode")
	if m {
		cmdColumns = append(cmdColumns, "mode")
	}
	c, _ := flags.GetBool("chan")
	if c {
		cmdColumns = append(cmdColumns, "chan")
	}
	r, _ := flags.GetBool("rate")
	if r {
		cmdColumns = append(cmdColumns, "rate")
	}
	g, _ := flags.GetBool("signal")
	if g {
		cmdColumns = append(cmdColumns, "signal")
	}
	b, _ := flags.GetBool("bars")
	if b {
		cmdColumns = append(cmdColumns, "bars")
	}
	s, _ := flags.GetBool("security")
	if s {
		cmdColumns = append(cmdColumns, "security")
	}
	u, _ := flags.GetBool("in-use")
	if u {
		cmdColumns = append(cmdColumns, "in-use")
	}
	if len(cmdColumns) > 0 {
		return fmt.Sprintf("-f %v", strings.Join(cmdColumns, ","))
	} else {
		return ""
	}
}

func renderNetworks(flags *pflag.FlagSet, networks []string) {
	u, _ := flags.GetBool("in-use")

	for i, n := range networks {
		if i == len(networks)-1 {
			break
		}

		if i == 0 {
			fmt.Println(n)
			continue
		} else if u {
			isMatched, _ := regexp.MatchString("(.*)(\\s)+\\*", n)
			if isMatched {
				fmt.Println(n)
				break
			}
		} else {
			num := fmt.Sprintf("%d-", i)
			fmt.Println(num, n)
		}
	}
}

var listCmd = &cobra.Command{
	Use:       "list",
	Short:     "List wifi networks",
	ValidArgs: []string{},

	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		command := fmt.Sprintf("nmcli %v dev wifi list", parseFlags(flags))
		c := exec.Command("bash", "-c", command)
		o, err := c.Output()
		utils.Must(err)

		networks := strings.Split(string(o), "\n")
		renderNetworks(flags, networks)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("name", "n", false, "show only networks' name(ssid)")
	listCmd.Flags().BoolP("mode", "m", false, "show only networks' mode")
	listCmd.Flags().BoolP("chan", "c", false, "show only networks' chan")
	listCmd.Flags().BoolP("rate", "r", false, "show only networks' rate")
	listCmd.Flags().BoolP("signal", "g", false, "show only networks' signal")
	listCmd.Flags().BoolP("bars", "b", false, "show only networks' bars")
	listCmd.Flags().BoolP("security", "s", false, "show only networks' security")
	listCmd.Flags().BoolP("in-use", "u", false, "show only used network")
}
