package aggregate_test

import (
	"testing"

	"github.com/aceino/belajar-ddd/src/domain/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type TestCase struct {
		test        string
		name        string
		email       string
		expectedErr error
	}

	test := []TestCase{

		{
			test:        "Empty Name",
			name:        "",
			email:       "",
			expectedErr: aggregate.ErrInvalidCustomer,
		},
		{
			test:        "Name",
			name:        "john Doe",
			email:       "johdoe@example.com",
			expectedErr: nil,
		},
	}

	for _, tc := range test {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name, tc.email)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v: got %v", tc.expectedErr, err)
			}
		})
	}
}
