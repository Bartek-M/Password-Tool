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
	"unicode"
)

func generate(n int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_|/\\'\""
	password := make([]byte, n)

	for i := 0; i < n; i++ {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}

func validate(passw string) int {
	score := 0

	if len(passw) >= 16 {
		score += 2
	} else if len(passw) >= 8 {
		score += 1
	}

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, char := range passw {
		if unicode.IsNumber(char) {
			hasDigit = true
		} else if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		} else {
			hasSpecial = true
		}
	}

	if hasUpper {
		score += 2
	}
	if hasLower {
		score += 2
	}
	if hasDigit {
		score += 2
	}
	if hasSpecial {
		score += 2
	}

	return score
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
		fmt.Println("Validated password score:", validate(text), "/ 10")
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
