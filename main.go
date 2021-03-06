package main

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"

	"github.com/1000ch/jpg2png/jpg2png"
)

func walkDir(arg string) []string {
	var fileList []string

	pathList, _ := ioutil.ReadDir(arg)
	for _, path := range pathList {
		name := filepath.Join(arg, path.Name())

		if path.IsDir() {
			fileList = append(fileList, walkDir(name)...)
		} else {
			fileList = append(fileList, name)
		}
	}

	return fileList
}

func distinct(args []string) []string {
	var fileList []string
	appeared := map[string]bool{}

	for _, arg := range args {
		if !appeared[arg] {
			appeared[arg] = true
			fileList = append(fileList, arg)
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

	wg := &sync.WaitGroup{}
	for _, file := range distinct(fileList) {
		if strings.HasSuffix(file, ".jpg") {
			wg.Add(1)
			go func(file string) {
				jpg2png.Convert(file)
				wg.Done()
			}(file)
		}
	}
	wg.Wait()
}
