package main

import (
	"bufio"
	"fmt"
	goterminal "go-terminal/functions/ls"
	"log"
	"os"
	"strings"
)

const (
	ls   = "ls"
	menu = "menu"
	help = "help"
	exit = "exit"
)

func main() {

	fmt.Println("--- terminal emulator written in go ---")

	realListener := goterminal.RealListFiles{}

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
			goterminal.HandleLS(args, realListener)

		case menu, help:
			fmt.Println("list of command available commands: ls")

		case exit:
			return

		default:
			fmt.Fprintf(os.Stdout, "command not known: %s\n", string(mainCommand))
		}
	}

}
