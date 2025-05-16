package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Credential struct {
	Service  string `json:"service"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CredentialStore struct {
	Credentials []Credential `json:"credentials"`
	filepath    string
}

func newCredentialStore() *CredentialStore {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return &CredentialStore{filepath: "credentials.json"}
	}

	dataDir := filepath.Join(homeDir, ".guardatuti")
	err = os.MkdirAll(dataDir, 0700)
	if err != nil {
		fmt.Println("Error creating data directory:", err)
		return &CredentialStore{filepath: "credentials.json"}
	}

	filepath := filepath.Join(dataDir, "credentials.json")
	store := &CredentialStore{filepath: filepath}
	store.load()
	return store
}

func (cs *CredentialStore) load() {
	data, err := ioutil.ReadFile(cs.filepath)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Error reading credentials file:", err)
		}
		return
	}

	err = json.Unmarshal(data, cs)
	if err != nil {
		fmt.Println("Error parsing credentials file:", err)
	}
}

func (cs *CredentialStore) save() {
	data, err := json.MarshalIndent(cs, "", "  ")
	if err != nil {
		fmt.Println("Error serializing credentials:", err)
		return
	}

	err = ioutil.WriteFile(cs.filepath, data, 0600)
	if err != nil {
		fmt.Println("Error writing credentials file:", err)
	}
}

func (cs *CredentialStore) add(service, username, password string) {
	cred := Credential{
		Service:  service,
		Username: username,
		Password: password,
	}
	cs.Credentials = append(cs.Credentials, cred)
	cs.save()
	fmt.Println("Credentials saved successfully!")
}

func (cs *CredentialStore) list() {
	if len(cs.Credentials) == 0 {
		fmt.Println("No credentials saved yet.")
		return
	}

	fmt.Println("\nStored credentials:")
	fmt.Println("-------------------")
	for i, cred := range cs.Credentials {
		fmt.Printf("%d. Service: %s, Username: %s\n", i+1, cred.Service, cred.Username)
	}
}

func (cs *CredentialStore) get(service string) {
	for _, cred := range cs.Credentials {
		if cred.Service == service {
			fmt.Printf("\nService: %s\nUsername: %s\nPassword: %s\n",
				cred.Service, cred.Username, cred.Password)
			return
		}
	}
	fmt.Printf("No credentials found for service: %s\n", service)
}

func (cs *CredentialStore) delete(service string) {
	for i, cred := range cs.Credentials {
		if cred.Service == service {
			// Remove the credential at index i
			cs.Credentials = append(cs.Credentials[:i], cs.Credentials[i+1:]...)
			cs.save()
			fmt.Printf("Credentials for %s deleted.\n", service)
			return
		}
	}
	fmt.Printf("No credentials found for service: %s\n", service)
}

var store *CredentialStore

func main() {
	store = newCredentialStore()
	fmt.Println("Welcome to GuardaTuti Password Manager!")

	for {
		displayMenu()
		input := waitForInput()
		executeCommand(input)
	}
}

func waitForInput() string {
	fmt.Print("> ")
	var input string
	fmt.Scanln(&input)
	return input
}

func executeCommand(command string) {
	switch command {
	case "1":
		addCredential()
	case "2":
		listCredentials()
	case "3":
		getCredential()
	case "4":
		deleteCredential()
	case "5":
		exit()
	default:
		fmt.Println("Unknown command. Please try again.")
	}
}

func displayMenu() {
	fmt.Println("\nPASSWORD MANAGER MENU")
	fmt.Println("--------------------")
	fmt.Println("1. Add new credentials")
	fmt.Println("2. List all services")
	fmt.Println("3. Get password for a service")
	fmt.Println("4. Delete credentials")
	fmt.Println("5. Exit")
}

func addCredential() {
	fmt.Print("Enter service name: ")
	var service string
	fmt.Scanln(&service)

	fmt.Print("Enter username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Print("Enter password: ")
	var password string
	fmt.Scanln(&password)

	store.add(service, username, password)
}

func listCredentials() {
	store.list()
}

func getCredential() {
	fmt.Print("Enter service name: ")
	var service string
	fmt.Scanln(&service)

	store.get(service)
}

func deleteCredential() {
	fmt.Print("Enter service name to delete: ")
	var service string
	fmt.Scanln(&service)

	store.delete(service)
}

func exit() {
	fmt.Println("Exiting...")
	os.Exit(0)
}
