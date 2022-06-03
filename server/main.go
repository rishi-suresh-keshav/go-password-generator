package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	gpg "github.com/rishi-suresh-keshav/go-password-generator/lib"
)

var (
	ContentTypeHeader = "Content-Type"
)

func main() {
	serveMux := mux.NewRouter()
	serveMux.HandleFunc("/go-password-generator/generate", generatePassword).Methods(http.MethodPost)

	fmt.Println("starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", serveMux))
}

func generatePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(ContentTypeHeader, "application/json")

	requestBody, _ := ioutil.ReadAll(r.Body)
	passwordRequest := &PasswordRequest{}
	_ = json.Unmarshal(requestBody, passwordRequest)

	pg := gpg.NewPasswordGenerator()
	if passwordRequest.Length > 0 {
		pg.WithLength(passwordRequest.Length)
	}

	password, err := pg.Generate()
	if err != nil {
		fmt.Println("error: ", err)
		errResponse, _ := json.Marshal(ErrorResponse{
			Error: err.Error(),
		})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errResponse)
		return
	}

	response, _ := json.Marshal(PasswordResponse{
		Password: password,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

type PasswordRequest struct {
	Length int32 `json:"length,omitempty"`
}

type PasswordResponse struct {
	Password string `json:"password"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
