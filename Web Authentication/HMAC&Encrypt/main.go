package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

var key = []byte{}

func main() {

	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

	msg := "We can attack them at dawn at 23:00 Hrs"
	sig, err := signMessage(msg)
	if err != nil {
		fmt.Println(err)
	}
	flag, err := checkSig(msg, sig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(flag)
	// hashed, err := passwordMake("12345678")
	// if err != nil {
	// 	fmt.Println("Geneatring pass")
	// }

	// err = compare(hashed, "12345678")
	// if err != nil {
	// 	fmt.Println("Not logged in")
	// }
	// fmt.Println("Logged in")
}

//bcrypt

func passwordMake(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func compare(hashed []byte, password string) error {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(password)); err != nil {
		return err
	}
	return nil
}

//hmac is a crypto signing algo

func signMessage(msg string) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	if _, err := h.Write([]byte(msg)); err != nil {
		return nil, err
	}
	signature := h.Sum(nil)
	return signature, nil
}

func checkSig(msg string, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, err
	}

	status := hmac.Equal(newSig, sig)

	return status, nil

}
