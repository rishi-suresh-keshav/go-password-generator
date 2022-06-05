package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rishi-suresh-keshav/go-password-generator/lib/controller"

	"github.com/gorilla/mux"
)

func main() {
	config := &Config{
		BindAddress: ":8000",
	}
	server := NewServer(config)

	err := server.Start()
	defer server.Stop()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Config struct {
	BindAddress string
}

type Server struct {
	Config *Config
}

func NewServer(config *Config) *Server {
	return &Server{
		config,
	}
}

func (s *Server) Start() error {
	serveMux := mux.NewRouter()
	serveMux.HandleFunc("/go-password-generator/generate", controller.GeneratePassword).Methods(http.MethodPost)

	fmt.Println("starting server at port 8000")
	return http.ListenAndServe(s.Config.BindAddress, serveMux)
}

func (s *Server) Stop() {

}
