package account

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		expectedOutput Account
		shouldErr      bool
	}{
		{
			name:  "base case",
			input: []byte("email=foo@bar.com&uid=10&role=user"),
			expectedOutput: Account{
				email: []byte("foo@bar.com"),
				uid:   10,
				role:  []byte("user"),
			},
			shouldErr: false,
		}, {
			name:           "invalid input",
			input:          []byte(")["),
			expectedOutput: Account{},
			shouldErr:      true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := Parse(test.input)
			if test.shouldErr {
				if err == nil {
					t.Error("expected: error, got: nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected: nil error, got: %s", err)
				}
				if res != &test.expectedOutput {
					t.Errorf("expected: %x, got %x", test.expectedOutput, *res)
				}
			}
		})
	}
}
