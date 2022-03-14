/*
Created by Jordan
Date: 7.3.22
*/
package lib

import (

	"fmt"
	"github.com/joho/godotenv"

	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	value := os.Getenv(key)
	if len(value) == 0 {
		return "NOT FOUND"
	}
	return value
}

func GenerateRandomString(lenght int, includeSpecial bool) string {

	var letterRunes []rune

	if includeSpecial {
		letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-.,;:_$!+*/")
	} else {
		letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	}

	b := make([]rune, lenght)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
