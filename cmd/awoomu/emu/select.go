package emu

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/LamkasDev/awoo-emu/cmd/common/logger"
	"github.com/manifoldco/promptui"
)

func SelectProgram() (string, error) {
	paths := make(map[string]string)
	list := []string{}
	filepath.Walk("data", func(path string, file os.FileInfo, err error) error {
		if err == nil && strings.Contains(file.Name(), ".awoobj") {
			name := strings.Replace(file.Name(), ".awoobj", "", -1)
			paths[name] = path
			list = append(list, name)
		}

		return nil
	})
	prompt := promptui.Select{
		Label:             "Select program to run",
		Items:             list,
		Size:              10,
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			return strings.Contains(strings.ToLower(list[index]), strings.ToLower(input))
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		logger.Log("Prompt failed %v\n", err)
		return "", err
	}

	return paths[result], nil
}
