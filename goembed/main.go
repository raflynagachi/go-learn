package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed img/LinuxCommand.png
var img []byte

//go:embed files/*.txt
var files embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("LinuxCommand.png", img, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := files.ReadDir("files")
	for _, entry := range dirEntries {
		fmt.Println(entry.Name())
		file, _ := files.ReadFile("files/" + entry.Name())
		fmt.Println(string(file))
	}
}
