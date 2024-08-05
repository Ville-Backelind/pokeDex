package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func pokeRepl() {
	scanner := bufio.NewScanner(os.Stdin)
off:
	for {
		fmt.Printf("pokedex > ")
		scanner.Scan()
		switch s := inputSanity(scanner.Text()); s {
		case "":
		case "exit":
			break off
		default:
			fmt.Printf("'%v', is an unknown command. \n- Type help for help page. \n- Type exit to quite the program.\n", s)
		}
	}
}

func inputSanity(inp string) string {
	lowerinp := strings.ToLower(inp)
	return lowerinp
}
