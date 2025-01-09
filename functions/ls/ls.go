package goterminal

import (
	"fmt"
	"io/fs"
	"os"
)

type FileLister interface {
	LS() []string
}

type RealListFiles struct{}

func (r RealListFiles) LS() []string {

	dir_content := []string{}

	dir_path := os.DirFS(".")

	files, err := fs.Glob(dir_path, "*.*")

	if err != nil {
		//aici pt mine, handle error
	}

	for _, elem := range files {
		dir_content = append(dir_content, elem)
	}

	return dir_content

}

func ListFiles(fileLister FileLister) []string {
	return fileLister.LS()
}

func HandleLS(args []string, fileLister FileLister) {

	files := ListFiles(fileLister)

	if len(args) == 0 {
		fmt.Println(files)
		return
	}

	for _, arg := range args {
		switch arg {
		case "-x":
			fmt.Println("case -x")
		}
	}

}
