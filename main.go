package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"math/rand"
)

type Credential struct {
	service   string
	username  string
	password  string
	createdAt time.Time
}

type CredentialList struct {
	credentials []Credential
}

func (c *CredentialList) add(credential Credential) {
	c.credentials = append(c.credentials, credential)
}

func (c *CredentialList) print() {
	for _, credential := range c.credentials {
		string_credential := fmt.Sprintf("|| %s ||  %s  ||  %s  || [%s]", credential.service, credential.username, credential.password, credential.createdAt.Format("2006-01-02 15:04:05"))
		fmt.Println(strings.Repeat("-", len(string_credential)))
		fmt.Println(string_credential)
		fmt.Println(strings.Repeat("-", len(string_credential)))
	}
}

func main() {
	fmt.Println("Hello, World!")
	credentials := CredentialList{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		credential := addCredential(scanner)
		credentials.add(credential)
		credentials.print()
	}
}

func addCredential(scanner *bufio.Scanner) Credential {
	credential := Credential{}
	fmt.Println("Generate credential")
	fmt.Println("Insert service name")
	service := waitForInput(scanner)
	credential.service = service
	fmt.Println("Insert user/email")
	username := waitForInput(scanner)
	credential.username = username
	credential.createdAt = time.Now()
	credential.password = generatePassword()
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

func saveCredential(credential Credential) {
	fmt.Println("saving credential....")
	fmt.Println(credential)
}
