package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
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
	credential.Password = GeneratePassword()
	return credential
}

func waitForInput(scanner *bufio.Scanner) string {
	fmt.Print(">")
	scanner.Scan()
	return scanner.Text()
}
