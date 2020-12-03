package aes128

import "testing"

func TestGenerate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "base case",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			key, err := GenerateKey()

			if err != nil {
				t.Errorf("expected: nil error, got: %s", err)
			}
			if len(key) != 16 {
				t.Errorf("expected: key length to be 16, got: %d", len(key))
			}
		})
	}
}
