package cli

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Encrypt    bool
	Decrypt    bool
	InputPath  string
	OutputPath string
	Key        int
}

func ParseArgs() (Config, error) {
	var config Config
	flag.BoolVar(&config.Encrypt, "e", false, "encrypt the input")
	flag.BoolVar(&config.Decrypt, "d", false, "decrypt the input")
	flag.StringVar(&config.InputPath, "i", "", "input file (defaults to STDIN)")
	flag.StringVar(&config.OutputPath, "o", "", "output file (defaults to STDIN)")
	flag.Parse()

	if config.Encrypt == config.Decrypt {
		return config, fmt.Errorf(
			"either -e (encrypt) or -d (decrypt) flag must be set")
	}

	if flag.NArg() != 1 {
		return config, fmt.Errorf(
			"expected exactly 1 argument, got %v (%s)", flag.NArg(), flag.Args())
	}

	key, err := strconv.Atoi(flag.Arg(0))
	if err != nil || key < 1 || key > 28 {
		return config, fmt.Errorf(
			"expected argument (key) to be an integer in range 1-28, got %s", flag.Arg(0))
	}
	config.Key = key

	return config, nil
}

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: caesar -e|-d [-i file] [-o file] <key>\n\nOptions:")
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stderr, "  -%s: %s\n", f.Name, f.Usage)
		})
		fmt.Fprintln(os.Stderr, "\nArguments:\n  key: integer in range 1-28")

		fmt.Fprintln(os.Stderr, "\nExamples:")
		examples := []string{
			"cat file.txt | caesar -e 5 > encrypted.txt",
			"caesar -d -i encrypted.txt -o decrypted.txt 5",
		}
		for _, ex := range examples {
			fmt.Fprintln(os.Stderr, "  "+ex)
		}
	}
}
