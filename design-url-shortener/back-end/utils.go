package main

import "github.com/aidarkhanov/nanoid/v2"

func GenerateURL() (string, error) {
	key, err := nanoid.GenerateString(Alphabet, NumberOfCharacters)
	if err != nil {
		return "", err
	}

	return key, nil
}
