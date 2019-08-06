package utils

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/omaressameldin/wifi-selector/network"
)

func Must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func SelectFromList(label string, list []network.SavedNetwork, icon string) int {
	names := network.ListNames(list)
	prompt := promptui.Select{
		Label: label,
		Items: names,
		Size:  int(math.Max(float64(len(list)), 10)),
		Templates: &promptui.SelectTemplates{
			Active:   fmt.Sprintf("%v  {{ . | red | underline}}", icon),
			Inactive: "   {{ . | white | faint }}",
		},
	}

	index, _, err := prompt.Run()
	Must(err)
	return index
}

func GetSavedWifis() []network.SavedNetwork {
	c := exec.Command("bash", "-c", "ls /etc/NetworkManager/system-connections/")
	o, err := c.Output()
	Must(err)

	savedNetworks := strings.Split(string(o), "\n")
	savedNetworksSet := make(map[string]bool)
	var savedNetworksUniqArr []network.SavedNetwork

	for i, n := range savedNetworks {
		if i == len(savedNetworks)-1 {
			break
		}
		networkName := strings.Replace(n, ".nmconnection", "", 1)
		if savedNetworksSet[networkName] {
			continue
		}
		savedNetworksUniqArr = append(savedNetworksUniqArr, network.SavedNetwork{
			Name:     networkName,
			Filename: n,
		})
		savedNetworksSet[networkName] = true
	}

	return savedNetworksUniqArr
}
