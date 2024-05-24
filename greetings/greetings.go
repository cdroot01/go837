package greetings 

import (
	"fmt"
	"errors"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func randomFormat() string {

	formats := []string {
		"Hi, %v",
		"Greet to see you, %v",
		"Hola, %v",
	}

	return formats[rand.Intn(len(formats))]
}
