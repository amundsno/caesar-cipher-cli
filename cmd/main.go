package main

import (
	"fmt"
	"os"

	"github.com/amundsno/caesar-cipher-cli/internal/caesar"
	"github.com/amundsno/caesar-cipher-cli/internal/cli"
	"github.com/amundsno/caesar-cipher-cli/internal/io"
)

func exit(err error) {
	fmt.Fprintf(os.Stderr, "caesar: %s - see 'caesar -h' for usage \n", err)
	os.Exit(1)
}

func main() {
	config, err := cli.ParseArgs()
	if err != nil {
		exit(err)
	}

	inputFile, err := io.OpenInput(config.InputPath)
	if err != nil {
		exit(err)
	}
	defer func() {
		if inputFile != os.Stdin {
			inputFile.Close()
		}
	}()

	outputFile, err := io.OpenOutput(config.OutputPath)
	if err != nil {
		exit(err)
	}
	defer func() {
		if outputFile != os.Stdout {
			outputFile.Close()
		}
	}()

	cipherFunc := caesar.Encrypt
	if config.Decrypt {
		cipherFunc = caesar.Decrypt
	}

	err = io.Transform(inputFile, outputFile, func(s string) string {
		return cipherFunc(s, config.Key)
	})
	if err != nil {
		exit(err)
	}

}
