package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"math/rand"
)

func main() {
	fmt.Println("Hello, World!")
	credentials := CredentialList{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		credential := addCredential(scanner)
		credentials.add(credential)
		credentials.save("credentials.json")
		listCredentials()
	}
}

func addCredential(scanner *bufio.Scanner) Credential {
	credential := Credential{}
	fmt.Println("Generate credential")
	fmt.Println("Insert service name")
	service := waitForInput(scanner)
	credential.Service = service
	fmt.Println("Insert user/email")
	username := waitForInput(scanner)
	credential.Username = username
	credential.CreatedAt = time.Now()
	credential.Password = generatePassword()
	return credential
}

func waitForInput(scanner *bufio.Scanner) string {
	fmt.Print(">")
	scanner.Scan()
	return scanner.Text()
}

func generatePassword() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!.-_")
	b := make([]rune, 15)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
