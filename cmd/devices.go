package cmd

import (
	"fmt"
	// "regexp"
	"strings"

	"github.com/omaressameldin/wifi-manager/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)


func parseDevicesFlags(flags *pflag.FlagSet) string {
	var cmdColumns []string
	d, _ := flags.GetBool("device")
	if d {
		cmdColumns = append(cmdColumns, "device")
	}
	t, _ := flags.GetBool("type")
	if t {
		cmdColumns = append(cmdColumns, "type")
	}
	s, _ := flags.GetBool("state")
	if s {
		cmdColumns = append(cmdColumns, "state")
	}
	c, _ := flags.GetBool("connection")
	if c {
		cmdColumns = append(cmdColumns, "connection")
	}
	dp, _ := flags.GetBool("dbus-path")
	if dp {
		cmdColumns = append(cmdColumns, "dbus-path")
	}
	cp, _ := flags.GetBool("con-path")
	if cp {
		cmdColumns = append(cmdColumns, "con-path")
	}
	cu, _ := flags.GetBool("con-uuid")
	if cu {
		cmdColumns = append(cmdColumns, "con-uuid")
	}

	if len(cmdColumns) > 0 {
		return fmt.Sprintf("-f %v", strings.Join(cmdColumns, ","))
	} else {
		return ""
	}
}

func renderDevices(flags *pflag.FlagSet, devices []string) {
	for i, n := range devices {
		if i == len(devices)-1 {
			break
		}
		fmt.Println(n)
	}
}

var devsCmd = &cobra.Command{
	Use:       "devs",
	Short:     "List available devices",
	ValidArgs: []string{},

	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()

		devices := utils.GetAvailableDevices(parseDevicesFlags(flags))
		renderDevices(flags, devices)
	},
}


func init() {
	RootCmd.AddCommand(devsCmd)
	devsCmd.Flags().BoolP("device", "d", false, "show only device names")
	devsCmd.Flags().BoolP("type", "t", false, "show only device type")
	devsCmd.Flags().BoolP("state", "s", false, "show only device state")
	devsCmd.Flags().BoolP("dbus-path", "p", false, "show only dbus path")
	devsCmd.Flags().BoolP("connection", "c", false, "show only connection")
	devsCmd.Flags().BoolP("con-uuid", "u", false, "show only connectios uuid")
	devsCmd.Flags().BoolP("con-path", "q", false, "show only connections path")
}
