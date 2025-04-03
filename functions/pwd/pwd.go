package pwd

import "fmt"

type PrintWorkDir interface {
	PWD() string
}

type RealPrintWorkDir struct{}

func (r RealPrintWorkDir) PWD() string {

	//TODO have actual path
	dir_path := "/path/to/current/dir"

	return dir_path
}

func PrintWD(printWorkDir PrintWorkDir) string {
	return printWorkDir.PWD()
}

func HandlePWD(args []string, printWorkDir PrintWorkDir) {

	currentWorkDir := PrintWD(printWorkDir)

	if len(args) == 0 {
		fmt.Println(currentWorkDir)
	}

	//TODO handle args
	for _, arg := range args {
		switch arg {
		case "-x":
			fmt.Println("case -x")
		}
	}

}
