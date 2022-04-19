package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

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
		if f.IsDir() && !directoryContainsClasses(path + "/" + f.Name()) {
			continue
		}
		renameFileInPath(path, f.Name(), strcase.ToCamel(f.Name()))
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
	err := os.Rename(path+"/"+old, path+"/"+new)
	handleError(err)
}

func directoryContainsClasses(dir string) bool {
    fmt.Println("Checking for classes in ", dir)
	files, err := ioutil.ReadDir(dir)
	handleError(err)

	for _, f := range files {
		content, err := ioutil.ReadFile(dir+"/"+f.Name())
		handleError(err)
		if strings.Contains(string(content), "class ") {
		    fmt.Println(string(content))
			return true
		}
	}
	return false
}
