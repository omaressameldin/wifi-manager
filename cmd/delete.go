package cmd

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/omaressameldin/wifi-selector/utils"
	"github.com/spf13/cobra"
)

func listSavedWifis() []string {
	c := exec.Command("bash", "-c", "ls /etc/NetworkManager/system-connections/")
	o, err := c.Output()
	utils.Must(err)

	savedNetworks := strings.Split(string(o), "\n")
	savedNetworksSet := make(map[string]bool)
	var savedNetworksUniqArr []string

	for i, n := range savedNetworks {
		if i == len(savedNetworks)-1 {
			break
		}
		networkName := strings.Replace(n, ".nmconnection", "", 1)
		if savedNetworksSet[networkName] {
			continue
		}
		savedNetworksUniqArr = append(savedNetworksUniqArr, networkName)
		savedNetworksSet[networkName] = true
	}

	for i, n := range savedNetworksUniqArr {
		num := fmt.Sprintf("%d-", i+1)
		fmt.Println(num, n)
	}

	return savedNetworksUniqArr
}

func selectWifi(wifiNetworks []string) string {
	wifiChoiceChan := make(chan string)
	fmt.Printf("write the number of the wifi network you wish to delete: ")
	go func() {
		var wifiNum string
		fmt.Scanf("%s\n", &wifiNum)
		wifiChoiceChan <- wifiNum
	}()
	select {
	case choice := <-wifiChoiceChan:
		choiceInNum, err := strconv.ParseInt(choice, 10, 32)
		utils.Must(err)
		return wifiNetworks[choiceInNum-1]
	}
}

func deleteWifi(selectedNet string) {
	c := exec.Command("bash", "-c", fmt.Sprintf("nmcli connection delete \"%v\"", selectedNet))
	o, err := c.Output()
	utils.Must(err)
	fmt.Println(string(o))
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete selected wifi network",
	Run: func(cmd *cobra.Command, args []string) {
		wifiLists := listSavedWifis()
		selected := selectWifi(wifiLists)
		deleteWifi(selected)
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
