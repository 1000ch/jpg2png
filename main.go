package main

import (
	"flag"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

func jpg2png(arg string) error {
	src, error := os.Open(arg)
	if error != nil {
		return error
	}
	defer src.Close()

	dest, error := os.Create(strings.Replace(arg, ".jpg", ".png", 1))
	if error != nil {
		return error
	}
	defer dest.Close()

	image, error := jpeg.Decode(src)
	if error != nil {
		return error
	}

	return png.Encode(dest, image)
}

func main() {
	flag.Parse()

	var fileList []string
	for _, arg := range flag.Args() {
		fileList = append(fileList, walkDir(arg)...)
	}

	for _, file := range distinct(fileList) {
		if strings.HasSuffix(file, ".jpg") {
			jpg2png(file)
		}
	}
}
