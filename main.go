package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: go run . <in> <out>")
		os.Exit(1)
	}
	input, output := os.Args[1], os.Args[2]

	text, err := readInput(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "lecture :", err)
		os.Exit(1)
	}

	result := Process(text)

	if err := writeOutput(output, result); err != nil {
		fmt.Fprintln(os.Stderr, "écriture :", err)
		os.Exit(1)
	}
}

func readInput(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func writeOutput(path string, text string) error {
	return os.WriteFile(path, []byte(text), 0644)
}
