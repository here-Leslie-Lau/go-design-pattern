package composite

import "fmt"

type component interface {
	Search(keyword string)
}

// file 文件类 - leaf
type file struct {
	size int
	name string
}

func (f *file) Search(keyword string) {
	fmt.Println("file", keyword, "size:", f.getSize())
}

func (f *file) getSize() int {
	return f.size
}

// folder 文件夹 - container
type folder struct {
	name  string
	files []component
}

func (f *folder) Search(keyword string) {
	fmt.Println("folder:", f.name)
	for _, file := range f.files {
		file.Search(keyword)
	}
}
