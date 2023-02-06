package emu

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

func SelectProgram() (string, error) {
	paths := make(map[string]string)
	list := []string{}
	err := filepath.Walk("data", func(path string, file os.FileInfo, err error) error {
		if err == nil && strings.Contains(file.Name(), ".awoobj") {
			name := strings.ReplaceAll(file.Name(), ".awoobj", "")
			paths[name] = path
			list = append(list, name)
		}

		return nil
	})
	if err != nil {
		return "", err
	}
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
		return "", err
	}

	return paths[result], nil
}
