package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//TODO: make a switch for handling different errors
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {

	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	log.Println("Bank running!")

	http.ListenAndServe(s.listenAddr, router)

}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	switch {
	case r.Method == "GET":
		return s.handleGetAccount(w, r)

	case r.Method == "POST":
		return s.handleCreateAccount(w, r)

	case r.Method == "DELETE":
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("Method not allowed")
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
