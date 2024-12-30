package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUSer(t *testing.T) {
	user, err := NewUser("Andre Diniz", "a@gmail.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Andre Diniz", user.Name)
	assert.Equal(t, "a@gmail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Andre Diniz", "a@gmail.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("12345"))
	assert.NotEqual(t, "123456", user.Password)
}
