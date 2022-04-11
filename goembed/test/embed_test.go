package goembed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

//go:embed img/BasicLinuxCommands.jfif
var img []byte

func TestEmbed(t *testing.T) {
	fmt.Println("version", version)
}

func TestEmbedImg(t *testing.T) {
	err := ioutil.WriteFile("img/LinuxCommand.png", img, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed version.txt
//go:embed img/BasicLinuxCommands.jfif
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	file1, _ := files.ReadFile("version.txt")
	file2, _ := files.ReadFile("img/BasicLinuxCommands.jfif")

	fmt.Println(string(file1))
	fmt.Println(file2)
}

//go:embed files/*.txt
var paths embed.FS

func TestMultiPathMatcher(t *testing.T) {
	dir, _ := paths.ReadDir("files")
	for _, entry := range dir {
		fmt.Println(entry.Name())
		file, _ := paths.ReadFile("files/" + entry.Name()) //byte representation
		fmt.Println(string(file))
	}
}
