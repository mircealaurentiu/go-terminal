package main

import (
	"bufio"
	"fmt"
	cmdLS "go-terminal/functions/ls"
	cmdPWD "go-terminal/functions/pwd"
	"log"
	"os"
	"strings"
)

const (
	ls   = "ls"
	pwd  = "pwd"
	menu = "menu"
	help = "help"
	exit = "exit"
)

func main() {

	fmt.Println("--- terminal emulator written in go ---")

	realListener := cmdLS.RealListFiles{}
	realPrintWorkDir := cmdPWD.RealPrintWorkDir{}

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("> ")

		scanner.Scan()
		err := scanner.Err()

		input := scanner.Text()

		if err != nil {
			log.Fatal(err)
		}

		mainCommand := strings.Fields(input)[0]
		args := strings.Fields(input)[1:]

		switch mainCommand {

		case ls:
			cmdLS.HandleLS(args, realListener)
		case pwd:
			cmdPWD.HandlePWD(args, realPrintWorkDir)

		case menu, help:
			fmt.Println("list of available commands: ls, pwd")

		case exit:
			return

		default:
			fmt.Fprintf(os.Stdout, "command not known: %s\n", string(mainCommand))
		}
	}

}
