package utils

import (
	"errors"
	"log"
	"math/rand"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateMinMaxRandomNumber() (int, error) {
	min := 111111
	max := 999999
	return min + rand.Intn(max-min), nil
}

func GenerateUserDefinedRandomNumber(min, max int) (int, error) {
	if min < 1 || (max < 1 && max > 9999) {
		return 0, errors.New("Invalid min and/or max params")
	}
	return min + rand.Intn(max-min), nil
}

func GenerateRandomNumber() (int, error) {
	min := 1
	max := 999999
	return min + rand.Intn(max-min), nil
}

func GenerateID() string {
	id, err := gonanoid.New()

	if err != nil {
		log.Println(err.Error())
		return "0"
	}
	return id
}

func GenerateName(size int) string {
	if size < 1 {
		size = 13
	}

	name, err := gonanoid.Generate("abcdefghijklmopqrstuvwxyzACDEFGHIJKLMOPQRSTUVWXYZ", size)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return name
}

func GenerateUID() string {
	uid, err := gonanoid.Generate("0123456789", 27)

	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return uid
}

func GenerateWord(size int) string {
	if size < 1 {
		size = 13
	}

	name, err := gonanoid.Generate("abcdefghijklmopqrstuvwxyzACDEFGHIJKLMOPQRSTUVWXYZ", size)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return name
}

func GenerateRandomString(size int) string {
	if size < 1 {
		size = 13
	}

	name, err := gonanoid.Generate("abcdefghijklmopqrstuvwxyzACDEFGHIJKLMOPQRSTUVWXYZ~!@#$%^&*,.?>}]|", size)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return name
}
