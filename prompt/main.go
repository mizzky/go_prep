package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	fmt.Println(getdir("./"))
}

func getdir(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, getdir(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}
