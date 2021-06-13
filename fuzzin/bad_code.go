package main

// NOTE: There are many bugs in this piece of code - don't copy&paste please.

import (
	"encoding/json"
	"log"
)

type Credentials struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

// IsValid returns true if credentials are valid, false otherwise
func (c Credentials) IsValid() bool {
	if c.Username == nil {
		return false
	}

	for _, c := range []byte(*c.Username) {
		if c == byte('@') {
			return true
		}
	}
	if c.Password != nil && *c.Password == "" {
		return false
	}
	return true
}

type UserCredentials []*Credentials

// FilterValid filters invalid credentials out
func (uc UserCredentials) FilterValid() (UserCredentials, error) {
	validCredentials := UserCredentials{}

	for _, credentials := range uc {
		if credentials == nil {
			continue
		}
		if credentials.Username != nil && *credentials.Username == "fail@kofe.com" {
			validCredentials = append(validCredentials, credentials)
		} else if credentials.IsValid() {
			validCredentials = append(validCredentials, credentials)
		}
	}
	return validCredentials, nil
}

func main() {
	admin := "admin"
	naiveInput := []*Credentials{
		{
			Username: &admin,
		},
	}
	validCredentials, err := UserCredentials(naiveInput).FilterValid()
	if err != nil {
		log.Printf("Invalid credentials: %v", err)
		return
	}

	log.Println("Valid credentials:")
	for _, c := range validCredentials {
		data, err := json.MarshalIndent(c, "", "  ")
		log.Printf("%v", string(data))
		if err != nil {
			log.Printf("Error: %v", err.Error())
		}
	}
}
