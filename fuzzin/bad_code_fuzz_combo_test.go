// +build gofuzzbeta

package main

import (
	"encoding/json"
	"github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func prettyPrint(creds UserCredentials) string {
	s, _ := json.MarshalIndent(creds, ",", "\t")
	return string(s)
}

func FuzzValidate(f *testing.F) {
	gof := fuzz.New().NilChance(0.2).NumElements(1, 5)
	for i := 0; i < 10000; i++ {
		creds := UserCredentials{}
		gof.Fuzz(&creds)

		sanitizedUserCredentials := UserCredentials{}
		// skip empty input
		for _, el := range creds {
			if el == nil || (*el == Credentials{}) {
				continue
			}
			if el.Username == nil || el.Password == nil {
				continue
			}
			if *el.Username == "" || *el.Password == "" {
				continue
			}
			sanitizedUserCredentials = append(sanitizedUserCredentials, el)
		}

		res, err := json.Marshal(sanitizedUserCredentials)

		assert.Nil(f, err)
		f.Add(res)
	}

	f.Fuzz(func(t *testing.T, rawCredentials []byte) {
		var credentials UserCredentials
		err := json.Unmarshal(rawCredentials, &credentials)
		if err != nil {
			t.Fatalf("Failed unmarshalling input %v: %v", prettyPrint(credentials), err)
		}
		creds, err := credentials.FilterValid()
		if err != nil {
			t.Fatalf("Validate failed to execute for %v: %v", prettyPrint(creds), err)
		}
		if len(credentials) != len(creds) {
			t.Fatalf("Some credentials are not valid %v: %v | %d: %d", prettyPrint(creds), err, len(credentials), len(creds))
		}
	})
}
