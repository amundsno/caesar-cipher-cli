# Caesar Cipher
A simple CLI implementing the [Caesar cipher](https://en.wikipedia.org/wiki/Caesar_cipher) for the Norwegian alphabet (i.e. including the letters `æ`, `ø`, and `å` following `z`).

## Quickstart
To build the binary, run:

```go
go build -o caesar cmd/main.go
```

To use it natively, move it to a directory included in `$PATH` or extend `$PATH` to include this directory.

```shell
export PATH="$(pwd):$PATH"
```

## Examples
```shell
# Input piped from STDIN, encrypted output to STDOUT
cat file.txt | caesar -e 5

# Specify input file to decrypt and output file
caesar -d -i encrypted.txt -o decrypted.txt 4
```
