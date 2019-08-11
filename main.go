package main

import (
	"github.com/omaressameldin/wifi-manager/cmd"
	"github.com/omaressameldin/wifi-manager/utils"
)

func main() {
	utils.Must(cmd.RootCmd.Execute())
}
