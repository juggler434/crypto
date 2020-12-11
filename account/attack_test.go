package account

import (
	"testing"
)

func TestAttack(t *testing.T) {
	tests := []struct {
		name           string
		expectedOutput bool
	}{
		{
			name:           "base case",
			expectedOutput: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			us := NewUserService()
			res := Attack(us)

			isAdmin, err := us.CheckAdminPermission(res)
			if err != nil {
				t.Errorf("unexpected error: %s", err)
			}

			if test.expectedOutput != isAdmin {
				t.Errorf("expected: %t, got: %t", test.expectedOutput, isAdmin)
			}
		})
	}
}
