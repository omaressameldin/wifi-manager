package main

import (
	"github.com/omaressameldin/wifi-selector/cmd"
	"github.com/omaressameldin/wifi-selector/utils"
)

func main() {
	utils.Must(cmd.RootCmd.Execute())
}
