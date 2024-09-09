package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRedirectService(t *testing.T) {
	svc := NewRedirectService(map[string]string{
		"/discord":   "https://discord.com",
		"/instagram": "https://instagram.com",
	})

	url, err := svc.GetRedirectURL("/discord")
	assert.Nil(t, err, "Error should be nil, discord exists")
	assert.Equal(t, "https://discord.com", url, "/discord should be redirected to https://discord.com")

	url, err = svc.GetRedirectURL("/instagram/")
	assert.NoError(t, err, "Error should be nil, instagram should work with slash at the end")
	assert.Equal(t, "https://instagram.com", url, "/instagram/ should be redirected to https://instagram.com")

	url, err = svc.GetRedirectURL("/facebook")
	assert.Error(t, err, "Error should occur, /facebook does not exists")
	assert.Equal(t, "", url, "URL should be empty, /facebook does not exists")
}
