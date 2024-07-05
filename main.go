package main

import (
	"bufio"
	"fmt"

	//"strings"
	"os"
)

func main() {
off:
	for {
		fmt.Printf("PokeDex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		switch s := scanner.Text(); s {
		case "":
		case "exit":
			break off
		default:
			fmt.Printf("'%v', is an unknown command. \n\n- Type help for help page. \n\n- Type exit to quite the program.\n\n", s)
		}
	}

}
