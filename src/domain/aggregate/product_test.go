package aggregate_test

import (
	"testing"

	"github.com/aceino/belajar-ddd/src/domain/aggregate"
)

func TestProduct_NewProduct(t *testing.T) {

	type testCase struct {
		test        string
		name        string
		year        int
		color       string
		price       float64
		expectedErr error
	}

	test := []testCase{
		{
			test:        "Invalid",
			name:        "",
			expectedErr: aggregate.ErrMisingValue,
		},
		{
			test:        "Valid",
			name:        "johnDoe",
			year:        2002,
			color:       "red",
			price:       10.23,
			expectedErr: nil,
		},
	}

	for _, tc := range test {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewProduct(tc.name, tc.year, tc.color, tc.price)
			if err != nil {
				t.Errorf("Expected error %v: got %v", tc.expectedErr, err)
			}
		})
	}

}
