package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"

func generatePassword(length int) string {
	password := make([]byte, length)
	for i := range password {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[index.Int64()]
	}
	return string(password)
}

func main() {
	fmt.Println("Gerador de Senhas")
	fmt.Println("Digite o comprimento da senha:")
	var length int
	fmt.Scanln(&length)

	if length < 1 {
		fmt.Println("Comprimento deve ser maior que zero!")
		return
	}

	password := generatePassword(length)
	fmt.Printf("Senha gerada: %s\n", password)
}
