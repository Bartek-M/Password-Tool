package main

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func generate(n int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_|/\\'\""
	password := make([]byte, n)

	for i := 0; i < n; i++ {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}

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

		fmt.Println(generate(num))
	} else if option == "val" {
		fmt.Println("Validate password")
	} else if option == "hash" {
		fmt.Println("Hashed password:", hash(text))
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

		if option == "quit" || option == "quit()" || !err {
			break
		}

		menu(option)
	}
}
