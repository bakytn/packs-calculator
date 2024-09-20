package packs

import (
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	// Test cases
	tests := []struct {
		name           string
		itemsOrdered   int
		expectedResult map[int]int
		errorExpected  bool
	}{
		{
			name:         "Order 1",
			itemsOrdered: 1,
			expectedResult: map[int]int{
				250: 1,
			},
			errorExpected: false,
		},
		{
			name:         "Order 250",
			itemsOrdered: 250,
			expectedResult: map[int]int{
				250: 1,
			},
			errorExpected: false,
		},
		{
			name:         "Order 251",
			itemsOrdered: 251,
			expectedResult: map[int]int{
				500: 1,
			},
			errorExpected: false,
		},
		{
			name:         "Order 501",
			itemsOrdered: 501,
			expectedResult: map[int]int{
				500: 1,
				250: 1,
			},
			errorExpected: false,
		},
		{
			name:         "Order 12001",
			itemsOrdered: 12001,
			expectedResult: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
			errorExpected: false,
		},
		{
			name:           "Invalid",
			itemsOrdered:   0,
			expectedResult: nil,
			errorExpected:  true,
		},
		{
			name:         "Ten million orders",
			itemsOrdered: 10000000,
			expectedResult: map[int]int{
				5000: 2000,
			},
			errorExpected: false,
		},
		{
			name:           "More than 10 million orders",
			itemsOrdered:   10000001,
			expectedResult: nil,
			errorExpected:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			result, err := CalculatePacks(tt.itemsOrdered)

			if tt.errorExpected {
				if err == nil {
					t.Error("Error expected")
				}
			}

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("CalculatePacks(%d) = %v; want %v",
					tt.itemsOrdered, result, tt.expectedResult)
			}
		})
	}
}
