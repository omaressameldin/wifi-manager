package utils

import (
	"fmt"
	"math"
	"os"

	"github.com/manifoldco/promptui"
)

func Must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func SelectFromList(label string, list []string, icon string) int {
	prompt := promptui.Select{
		Label: label,
		Items: list,
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
