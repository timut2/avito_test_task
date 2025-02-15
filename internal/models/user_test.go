package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_HashPassword(t *testing.T) {
	user := &User{
		Password: "testpassword",
	}

	err := user.HashPassword()
	assert.NoError(t, err)
	assert.NotEqual(t, "testpassword", user.Password)
}

func TestUser_CheckPassword_CorrectPassword(t *testing.T) {
	user := &User{
		Password: "testpassword",
	}

	err := user.HashPassword()
	assert.NoError(t, err)

	err = user.CheckPassword("testpassword")
	assert.NoError(t, err)
}

func TestUser_CheckPassword_IncorrectPassword(t *testing.T) {
	user := &User{
		Password: "testpassword",
	}

	err := user.HashPassword()
	assert.NoError(t, err)

	err = user.CheckPassword("wrongpassword")
	assert.Error(t, err)
}
