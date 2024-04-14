package main

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func hash(passw string) string {
	hasher := sha512.New()
	hasher.Write([]byte(passw))

	hashSum := hasher.Sum(nil)
	hexHash := hex.EncodeToString(hashSum)

	return hexHash
}

func menu(option string) {
	if option != "gen" && option != "val" && option != "hash" {
		fmt.Println("Invalid option")
		return
	}

	if option == "gen" {
		fmt.Println("Generate password")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type password: ")
	scanner.Scan()
	password := scanner.Text()

	if option == "val" {
		fmt.Println("Validate password")
	} else if option == "hash" {
		fmt.Println("Hashed password:", hash(password))
	}
}

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

		if option == "quit" || !err {
			break
		}

		menu(option)
	}
}
