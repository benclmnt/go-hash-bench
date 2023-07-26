package main

import (
	"log"

	"github.com/alexedwards/argon2id"
	scrypt "github.com/elithrar/simple-scrypt"
	"golang.org/x/crypto/bcrypt"
)

func Bcrypt(password []byte, cost int) []byte {
	hash, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		log.Println(err)
	}
	return hash
}

func Argon2(password []byte, params *argon2id.Params) []byte {
	hash, err := argon2id.CreateHash(string(password), params)
	if err != nil {
		log.Println(err)
	}
	return []byte(hash)
}

func Scrypt(password []byte, params scrypt.Params) []byte {
	hash, err := scrypt.GenerateFromPassword(password, params)
	if err != nil {
		log.Println(err)
	}
	return hash
}
