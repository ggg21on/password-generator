package main

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+{}[];:,.<>/?|\\"

func main() {
	rand.Seed(time.Now().UnixNano())
	length := 16
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
    }
    fmt.Println(string(password))
}