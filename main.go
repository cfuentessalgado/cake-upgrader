package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

var ROOT_ONE_TO_TWO = map[string]string{
	"config":      "Config",
	"console":     "Console",
	"controllers": "Controller",
	"lib":         "Lib",
	"locale":      "Locale",
	"models":      "Model",
	"plugins":     "Plugin",
	"tests":       "Test",
	"vendors":     "Vendor",
	"views":       "View",
}

func main() {
	args := os.Args

	path := args[1]
	start, _ := strconv.Atoi(args[2])
	target, _ := strconv.Atoi(args[3])

	fmt.Println(path, start, target)

	if target-start > 1 {
		error := fmt.Errorf("Target version must not excede 1 version (start %d target %d)", start, target)
		fmt.Println(error)
	}

	if start == 1 && target == 2 {
		oneToTwo(path)
	}
}

func oneToTwo(path string) {
	camelCaseDirectories(path + "/app")
}

func camelCaseDirectories(path string) {
	files, err := ioutil.ReadDir(path)
	handleError(err)
	fmt.Println("Listing directories to camel case:")
	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		if directoryIsNotAllowed(f.Name()) {
			continue
		}
		if !shouldChange(f.Name()) {
			continue
		}
		fmt.Println(f.Name())
		target := strcase.ToCamel(f.Name())
		renameFileInPath(path, f.Name(), target)
	}
}

func directoryIsNotAllowed(f string) bool {
	return f == "tmp" || f == "webroot" || f == "composer-modules"
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func renameFileInPath(path string, old string, new string) {
	_ = os.Rename(path+"/"+old, path+"/"+new)

}

func directoryContainsClasses(dir string) bool {
	files, err := ioutil.ReadDir(dir)
	handleError(err)

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		content, err := ioutil.ReadFile(dir + "/" + f.Name())
		handleError(err)
		if strings.Contains(string(content), "class ") {
			return true
		}
	}
	return false
}

func shouldChange(folder string) bool {
    fmt.Println(folder)
	if _, k := ROOT_ONE_TO_TWO[folder]; k {
		return true
	}
	return false
}
