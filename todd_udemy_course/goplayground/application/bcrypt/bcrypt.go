// bcrypt is a password hashing function designed by Niels Provos and David Mazières,
// based on the Blowfish cipher. Refer this Wiki page for more information https://en.wikipedia.org/wiki/Bcrypt
package main

import (
	"fmt"
	// `go get` below package since this is not a standard package
	"golang.org/x/crypto/bcrypt"
)

func main() {
	userPassword := `1q2w3e4r`

	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	bs, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("User password:", userPassword)
	fmt.Printf("%v\n", bs)

	loginPassword1 := `1q2w3e4r`
	// loginPassword2 := `q1w2e3r4`

	// func CompareHashAndPassword(hashedPassword, password []byte) error
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword1))
	// err = bcrypt.CompareHashAndPassword(bs, []byte(loginPassword2))
	if err != nil {
		fmt.Println("Incorrect Password. Failed to login!!")
		return
	}
	fmt.Println("Login Successful!")
}
