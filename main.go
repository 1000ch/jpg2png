package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func walkDir(arg string) []string {
	var fileList []string

	pathList, _ := ioutil.ReadDir(arg)
	for _, path := range pathList {
		name := path.Name()
		if path.IsDir() {
			fileList = append(fileList, walkDir(name)...)
		} else {
			fileList = append(fileList, name)
		}
	}

	return fileList
}

func main() {
	flag.Parse()

	var fileList []string
	for _, arg := range flag.Args() {
		fileList = append(fileList, walkDir(arg)...)
	}

	for _, file := range fileList {
		fmt.Println(file)
	}
}
