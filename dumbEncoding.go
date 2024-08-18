package main

import (
    "fmt"
)

func SimpleEncode(input string) string {
	encoded := ""
	for _, char := range input {
		encoded += fmt.Sprintf("%08b", char)
	}
	return encoded
}
