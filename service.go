package main

import (
	"errors"
	"strings"
)

type Service interface {
	GetRedirectURL(path string) (string, error)
}

type RedirectService struct {
	redirects map[string]string
}

func NewRedirectService(redirects map[string]string) Service {
	return &RedirectService{
		redirects: redirects,
	}
}

func (s *RedirectService) GetRedirectURL(path string) (string, error) {
	path = strings.TrimSuffix(path, "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	url, ok := s.redirects[path]
	if !ok {
		return "", errors.New("unknown path")
	}
	return url, nil
}
