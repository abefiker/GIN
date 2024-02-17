package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
func randomFormat()string{
	format := []string{
		"Hi , %v Welcome!",
		"Great to see you, %v",
		"Hail , %v Well met!",
	}
	return format[rand.Intn(len(format))]
}