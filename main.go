package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to Guardatuti! Type 'help' for available commands.")
	credentials := CredentialList{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		command := strings.TrimSpace(scanner.Text())

		switch command {
		case "help":
			HelpCommand()
		case "list":
			ListCommand()
		case "generate":
			credential := addCredential(scanner)
			credentials.add(credential)
			credentials.save("credentials.json")
			fmt.Println("Credential generated and saved successfully!")
		case "quit", "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}
}

func HelpCommand() {
	fmt.Println("Available commands:")
	fmt.Println("  help     - Show this help message")
	fmt.Println("  list     - List all saved credentials")
	fmt.Println("  generate - Generate a new credential")
	fmt.Println("  quit     - Exit the program")
}

func ListCommand() {
	file, err := os.Open("credentials.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	credentials := CredentialList{}

	for {
		var credential Credential
		err := decoder.Decode(&credential)
		if err != nil {
			break
		}
		credentials.CredentialList = append(credentials.CredentialList, credential)
	}

	if len(credentials.CredentialList) == 0 {
		fmt.Println("No credentials found")
		return
	}

	for _, credential := range credentials.CredentialList {
		string_credential := fmt.Sprintf("|| %s ||  %s  ||  %s  || [%s]", credential.Service, credential.Username, credential.Password, credential.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Println(strings.Repeat("-", len(string_credential)))
		fmt.Println(string_credential)
		fmt.Println(strings.Repeat("-", len(string_credential)))
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
	fmt.Println("Generated password: " + credential.Password)
	return credential
}

func waitForInput(scanner *bufio.Scanner) string {
	fmt.Print("> ")
	scanner.Scan()
	return scanner.Text()
}
