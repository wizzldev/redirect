package main

import (
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLogger(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetRedirectURL(path string) (url string, err error) {
	defer func(start time.Time) {
		fmt.Printf("path=%s to=%s err=%v took=%v\n", path, url, err, time.Since(start))
	}(time.Now())
	return s.next.GetRedirectURL(path)
}
