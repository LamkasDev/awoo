package paths

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ResolvePath(input string, ext string) string {
	if !filepath.IsAbs(input) {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		input = filepath.Join(wd, input)
	}
	if filepath.Ext(input) == "" {
		files, err := ioutil.ReadDir(input)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			if filepath.Ext(file.Name()) == ext {
				input = filepath.Join(input, file.Name())
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
	if filepath.Ext(output) == "" {
		if filepath.IsAbs(output) {
			output = filepath.Join(output, fmt.Sprintf("%s%s", inputName, outputExt))
		} else {
			output = filepath.Join(filepath.Dir(input), output, fmt.Sprintf("%s%s", inputName, outputExt))
		}
	} else if !filepath.IsAbs(output) {
		output = filepath.Join(filepath.Dir(input), output)
	}

	return input, output
}
