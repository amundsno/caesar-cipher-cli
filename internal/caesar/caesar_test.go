package caesar

import "testing"

func TestShift(t *testing.T) {
	tests := []struct {
		name, input, expected string
		n                     int
	}{
		{
			name:     "shift alphabet by +1 (encrypt)",
			input:    string(Alphabet),
			n:        1,
			expected: string(Alphabet[1:]) + "a",
		},
		{
			name:     "shift alphabet by -1 (decrypt)",
			input:    string(Alphabet),
			n:        -1,
			expected: "Å" + string(Alphabet[:len(Alphabet)-1]),
		},
		{
			name:     "shifting preserves other symbols",
			input:    "abc123!#$%&/(){}=^*@,.-><",
			n:        1,
			expected: "bcd123!#$%&/(){}=^*@,.-><",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := shift(test.input, test.n)
			if result != test.expected {
				t.Errorf("expected %q, got %q\n", test.expected, result)
			}
		})
	}
}

func TestEncryptDecrypt(t *testing.T) {
	tests := []struct {
		name, input string
		key         int
	}{
		{name: "ciphering preserves capitalization", input: "HeLlO, wOrlD!", key: 3},
		{name: "cipher by key > 28 supported", input: "hello, world!", key: 300},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			encrypted := Encrypt(test.input, test.key)
			decrypted := Decrypt(encrypted, test.key)
			if decrypted != test.input {
				t.Errorf("expected %q, got %q\n", test.input, decrypted)
			}
		})
	}
}
