package utils

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	spinner "github.com/omaressameldin/wifi-manager/Spinner"
	"github.com/omaressameldin/wifi-manager/network"
)

func Must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func SelectFromList(label string, list []string, icon string, isPositive bool) int {
	color := "red"
	if isPositive {
		color = "magenta"
	}

	prompt := promptui.Select{
		Label: label,
		Items: list,
		Size:  int(math.Min(float64(len(list)), 10)),
		Templates: &promptui.SelectTemplates{
			Active:   fmt.Sprintf("%v  {{ . | %v | underline}}", icon, color),
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

func GetAvailableNetworks(extraOptions string) []string {
	s := spinner.Spinner{
		Shape: 14,
	}
	s.StartSpinner("Searching")

	command := fmt.Sprintf("nmcli %v dev wifi list", extraOptions)
	c := exec.Command("bash", "-c", command)

	o, err := c.Output()
	s.StopSpinner()
	Must(err)

	return strings.Split(string(o), "\n")
}


func GetAvailableDevices(extraOptions string) []string {
	s := spinner.Spinner{
		Shape: 14,
	}
	s.StartSpinner("Searching")

	command := fmt.Sprintf("nmcli %v d s", extraOptions)
	c := exec.Command("bash", "-c", command)

	o, err := c.Output()
	s.StopSpinner()
	Must(err)

	return strings.Split(string(o), "\n")
}


