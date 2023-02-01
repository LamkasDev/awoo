package paths

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ResolvePath(input string, ext string) string {
	if !path.IsAbs(input) {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		input = filepath.Join(wd, input)
	}
	if path.Ext(input) == "" {
		files, err := ioutil.ReadDir(input)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			if filepath.Ext(file.Name()) == ext {
				input = path.Join(input, file.Name())
				break
			}
		}
	}
	if _, err := os.Stat(input); err != nil {
		panic(err)
	}

	return input
}

func ResolvePaths(input string, inputExt string, output string, outputExt string) (string, string) {
	input = ResolvePath(input, inputExt)
	inputName := strings.TrimSuffix(filepath.Base(input), filepath.Ext(input))
	if path.Ext(output) == "" {
		if path.IsAbs(output) {
			output = filepath.Join(output, fmt.Sprintf("%s%s", inputName, outputExt))
		} else {
			output = filepath.Join(filepath.Dir(input), output, fmt.Sprintf("%s%s", inputName, outputExt))
		}
	} else if !path.IsAbs(output) {
		output = filepath.Join(filepath.Dir(input), output)
	}

	return input, output
}
