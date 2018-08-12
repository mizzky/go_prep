package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		println("invalid args.")
		os.Exit(1)
	}
	//fmt.Println(getdir(os.Args[1]))
	fmt.Println(getdir(os.Args[1]))
	fmt.Println(getFilesize(os.Args[1]))
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

func getFilesize(dir string) []int64 {
	var sizes []int64

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fileInfo, statErr := os.Stat(filepath.Join(dir, file.Name()))
		if statErr != nil {
			println("err file stat")
		}

		fileSize := fileInfo.Size()
		sizes = append(sizes, fileSize)
		//fmt.Printf("fileName --> %v", file.Name())
		//fmt.Printf("fileSize --> %v \n", fileSize)
	}
	return sizes
}
