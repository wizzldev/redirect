package main

import (
	"io"
	"net/http"
	"os"
)

type ApiServer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (s *ApiServer) Start(listenAddr string) error {
	return http.ListenAndServe(listenAddr, s)
}

func (s *ApiServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusBadRequest)
		return
	}

	redirectURL, err := s.svc.GetRedirectURL(r.URL.Path)
	if err != nil {
		writeError(w, http.StatusNotFound)
		return
	}

	w.Header().Add("Location", redirectURL)
	w.WriteHeader(http.StatusPermanentRedirect)
}

func writeError(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	file, err := os.Open("./error.html")
	if err != nil {
		http.Error(w, "Failed to display error page.", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	io.Copy(w, file)
}
