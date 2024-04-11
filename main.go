package main

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
)

func hash(passw string) string {
	hasher := sha512.New()
	hasher.Write([]byte(passw))

	hashSum := hasher.Sum(nil)
	hexHash := hex.EncodeToString(hashSum)

	return hexHash
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type a password: ")
	scanner.Scan()

	input := scanner.Text()
	fmt.Println("Hashed password:", hash(input))
}
