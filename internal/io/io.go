package io

import (
	"bufio"
	"fmt"
	"os"
)

func isCharDevice(file *os.File) bool {
	info, _ := file.Stat()
	return info.Mode()&os.ModeCharDevice != 0
}

func OpenInput(path string) (*os.File, error) {
	if path == "" {
		file := os.Stdin
		if isCharDevice(file) {
			// STDIN attached to a dynamic 'char device' - e.g. terminal
			return nil, fmt.Errorf("no input file, and no data piped from STDIN")
		}
		return file, nil
	} else {
		file, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("could not open input file: %s", path)
		}
		return file, nil
	}
}

func OpenOutput(path string) (*os.File, error) {
	if path == "" {
		return os.Stdout, nil
	} else {
		file, err := os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("could not open output file: %s", path)
		}
		return file, nil
	}
}

func Transform(inputFile, outputFile *os.File, transformation func(string) string) error {
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	for scanner.Scan() {
		writer.WriteString(
			transformation(scanner.Text()) + "\n",
		)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
