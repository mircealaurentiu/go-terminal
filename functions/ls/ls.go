package goterminal

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
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
		//TODO handle error
	}

	/*
		for _, elem := range files {
			dir_content = append(dir_content, elem)
		}
	*/

	dir_content = append(dir_content, files...)

	return dir_content

}

func ListFiles(fileLister FileLister) []string {
	return fileLister.LS()
}

func HandleLS(args []string, fileLister FileLister) {

	files := ListFiles(fileLister)

	if len(args) == 0 {
		fmt.Println(strings.Join(files, " "))
		return
	}

	//TODO handle args
	for _, arg := range args {
		switch arg {
		case "-x":
			fmt.Println("case -x")
		}
	}

}
