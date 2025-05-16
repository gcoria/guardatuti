package models

type Credential struct {
	Service  string `json:"service"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CredentialStore struct {
	Credentials []Credential `json:"credentials"`
}
