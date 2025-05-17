package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Credential struct {
	Service   string    `json:"service"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type CredentialList struct {
	CredentialList []Credential `json:"credentialList"`
}

func (c *CredentialList) add(credential Credential) {
	c.CredentialList = append(c.CredentialList, credential)
}

func listCredentials() {
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

func (c *CredentialList) save(filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if len(c.CredentialList) > 0 {
		credential := c.CredentialList[len(c.CredentialList)-1]
		encoder.Encode(credential)
	}

	return nil
}
