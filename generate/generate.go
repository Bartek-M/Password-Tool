package generate

import "math/rand"


func Generate(n int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_|/\\'\""
	password := make([]byte, n)

	for i := 0; i < n; i++ {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}
