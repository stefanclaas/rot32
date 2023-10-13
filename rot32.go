package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func rot32(input string) string {
	rotated := ""
	for _, char := range input {
		if char >= 33 && char <= 126 {
			char = ((char - 33 + 32) % 94) + 33
		}
		rotated += string(char)
	}
	return rotated
}

func unrot32(input string) string {
	unrotated := ""
	for _, char := range input {
		if char >= 33 && char <= 126 {
			char = ((char - 33 - 32 + 94) % 94) + 33
		}
		unrotated += string(char)
	}
	return unrotated
}

func main() {
	var decode bool
	flag.BoolVar(&decode, "d", false, "Decode the input text")
	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Fprintln(os.Stderr, "Error: Unknown parameters:", flag.Args())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if decode {
			fmt.Println(unrot32(text))
		} else {
			fmt.Println(rot32(text))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		os.Exit(1)
	}
}
