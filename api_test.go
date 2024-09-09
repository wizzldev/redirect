package main

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type svc struct{}

func (*svc) GetRedirectURL(url string) (string, error) {
	if url == "/discord" {
		return "https://discord.com", nil
	}
	return "", errors.New("invalid path")
}

func getErrFile() ([]byte, error) {
	file, err := os.Open("./error.html")
	if err != nil {
		return nil, err
	}
	return io.ReadAll(file)
}

func TestApiServerGetDiscord(t *testing.T) {
	req, err := http.NewRequest("GET", "/discord", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	api := NewApiServer(&svc{})

	api.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusPermanentRedirect, rr.Code, "Status code should be permanent redirect")
	assert.Equal(t, "https://discord.com", rr.Header().Get("Location"), "Location should be https://discord.com")
}

func TestApiBadMethod(t *testing.T) {
	req, err := http.NewRequest("POST", "/discord", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	api := NewApiServer(&svc{})

	api.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "Status code should be bad request")

	content, err := getErrFile()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, string(content), rr.Body.String(), "Body should be error.html content")
}

func TestApiNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/instagram", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	api := NewApiServer(&svc{})

	api.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code, "Status code should be not found")

	content, err := getErrFile()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, string(content), rr.Body.String(), "Body should be error.html content")
}
