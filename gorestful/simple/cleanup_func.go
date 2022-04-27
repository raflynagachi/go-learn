package simple

import "fmt"

type File struct {
	Name string
}

func NewFile(name string) (*File, func(), error) {
	file := &File{
		Name: name,
	}
	return file,
		func() {
			file.Close()
		},
		nil
}

func (f *File) Close() {
	fmt.Println("Close file", f.Name)
}
