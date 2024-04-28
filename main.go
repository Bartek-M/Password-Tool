package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Choose an option:\n" +
		"  - gen [generate password]\n" +
		"  - val [validate password]\n" +
		"  - hash [hash password]\n" +
		"  - quit [quit loop]\n")

	for {
		fmt.Printf("\n> ")
		err := scanner.Scan()
		option := strings.ToLower(scanner.Text())

		if option == "quit" || option == "quit()" || !err {
			break
		}

		menu(option)
	}
}
