/*
Created by Jordan
Date: 7.3.22
functions:
	GoDotEnvVariable
	GenerateRandomString
*/
package lib

import (
	"fmt"

	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
)

//Gets all the environment Variables
//Values: key(DB connection String)
func GoDotEnvVariable(key string) string {

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

//Creates a Random String with normal characters or if needed special characters
//values: length(the length of the demanded string), includeSpecial(if SpecialCharacters are demandeds)
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
