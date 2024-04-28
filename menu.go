package main

import (
	"bufio"
	"fmt"
	"main/password-tool/generate"
	"main/password-tool/hash"
	"main/password-tool/validate"
	"os"
	"strconv"
)

func menu(option string) {
	if option != "gen" && option != "val" && option != "hash" {
		fmt.Println("Invalid option")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	checkText := "Type password: "

	if option == "gen" {
		checkText = "Password length: "
	}

	fmt.Print(checkText)
	scanner.Scan()
	text := scanner.Text()

	if option == "gen" {
		num, err := strconv.Atoi(text)
		if err != nil || num < 8 {
			fmt.Println("Invalid password length! Pass a number >= 8")
			return
		}

		fmt.Println(generate.Generate(num))
	} else if option == "val" {
		fmt.Println("Validated password score:", validate.Validate(text), "/ 10")
	} else if option == "hash" {
		fmt.Println("Hashed password:", hash.Hash(text))
	}
}
