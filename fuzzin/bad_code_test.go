package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateNilUsername(t *testing.T) {
	naiveInput := UserCredentials{
		{
			Username: nil,
		},
	}
	validCredentials, err := UserCredentials(naiveInput).FilterValid()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(validCredentials))
}

func TestValidateNilCredentials(t *testing.T) {
	naiveInput := UserCredentials{
		nil,
	}
	validCredentials, err := UserCredentials(naiveInput).FilterValid()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(validCredentials))
}

func TestValidateEmptyPassword(t *testing.T) {
	emptyUsername := ""
	emptyPassword := ""
	naiveInput := UserCredentials{
		{
			Username: &emptyUsername,
			Password: &emptyPassword,
		},
	}
	validCredentials, err := UserCredentials(naiveInput).FilterValid()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(validCredentials))
}

func TestValidateNormalCredentials(t *testing.T) {
	username := "_拑沿"
	password := "N鰌ƚ`H埏欚ǡ耬#珶浃ō`轟LVOȺ"
	naiveInput := UserCredentials{
		{
			Username: &username,
			Password: &password,
		},
	}
	validCredentials, err := UserCredentials(naiveInput).FilterValid()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(validCredentials))
}

func TestValidateAdmin(t *testing.T) {
	username := "fail@kofe.com"
	password := "epic pass"
	naiveInput := UserCredentials{
		{
			Username: &username,
			Password: &password,
		},
	}
	validCredentials, err := UserCredentials(naiveInput).FilterValid()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(validCredentials))
}
