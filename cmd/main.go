package main

import (
	"fmt"
	"os"

	gpg "github.com/rishi-suresh-keshav/go-password-generator/lib"
)

func main() {

	pg := gpg.NewPasswordGenerator().
		WithLength(11)

	password, err := pg.Generate()
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	fmt.Println("Here is your random password: ", password)
}
