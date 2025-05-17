package main

import (
	"encoding/json"
	"fmt"
	"os"
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
