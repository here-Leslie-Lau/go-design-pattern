package composite

import "testing"

func TestComposite(t *testing.T) {
	file1 := &file{
		size: 1,
		name: "file",
	}
	file2 := &file{
		size: 2,
		name: "file",
	}
	file3 := &file{
		size: 3,
		name: "file",
	}
	folder := &folder{
		name:  "folder",
		files: []component{file1, file2, file3},
	}

	folder.Search("keyword")
	// output:
	// folder: folder
	// file keyword size: 1
	// file keyword size: 2
	// file keyword size: 3
}
