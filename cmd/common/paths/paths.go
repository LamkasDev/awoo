package paths

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type AwooPath struct {
	Absolute string
	Anchor   *string
}

func CreatePath(path string) AwooPath {
	path = filepath.Clean(path)
	if !filepath.IsAbs(path) {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		return AwooPath{
			Absolute: filepath.Join(wd, path),
			Anchor:   &wd,
		}
	}
	if filepath.Ext(path) == "" {
		return AwooPath{
			Absolute: path,
			Anchor:   &path,
		}
	}

	return AwooPath{
		Absolute: path,
	}
}

func CreatePathList(path string, ext string) []AwooPath {
	resultPath := CreatePath(path)

	// If a path is folder, recursively create a list of files anchored from path
	if filepath.Ext(resultPath.Absolute) == "" {
		resultPaths := []AwooPath{}
		err := filepath.WalkDir(resultPath.Absolute, func(cpath string, _ fs.DirEntry, cerr error) error {
			if cerr != nil {
				return cerr
			}
			if filepath.Ext(cpath) == ext {
				resultPaths = append(resultPaths, AwooPath{
					Absolute: cpath,
					Anchor:   resultPath.Anchor,
				})
			}
			return nil
		})
		if err != nil {
			panic(err)
		}

		return resultPaths
	}

	// If path is a file, return single path
	if _, err := os.Stat(resultPath.Absolute); err != nil {
		panic(err)
	}

	return []AwooPath{resultPath}
}

func CreatePathListCombined(paths []string, ext string) []AwooPath {
	resultPaths := []AwooPath{}
	for _, path := range paths {
		resultPaths = append(resultPaths, CreatePathList(path, ext)...)
	}

	return resultPaths
}

func ResolveOutputPath(input AwooPath, output string, ext string) AwooPath {
	resultPath := CreatePath(output)
	unanchoredOutput := strings.TrimPrefix(input.Absolute, *input.Anchor)
	unanchoredOutput = strings.ReplaceAll(unanchoredOutput, filepath.Ext(unanchoredOutput), ext)
	resultPath.Absolute = filepath.Join(*resultPath.Anchor, unanchoredOutput)

	return resultPath
}

func ResolveAllPaths(inputs []string, inputExt string, output string, outputExt string) ([]AwooPath, []AwooPath) {
	resultInputs := CreatePathListCombined(inputs, inputExt)
	resultOutputs := []AwooPath{}
	if filepath.Ext(output) == "" {
		for _, input := range resultInputs {
			resultOutputs = append(resultOutputs, ResolveOutputPath(input, output, outputExt))
		}
	} else {
		resultOutputs = append(resultOutputs, CreatePath(output))
	}

	return resultInputs, resultOutputs
}
