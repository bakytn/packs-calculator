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
	}{
		{
			name:         "Order 1",
			itemsOrdered: 1,
			expectedResult: map[int]int{
				250: 1,
			},
		},
		{
			name:         "Order 250",
			itemsOrdered: 250,
			expectedResult: map[int]int{
				250: 1,
			},
		},
		{
			name:         "Order 251",
			itemsOrdered: 251,
			expectedResult: map[int]int{
				500: 1,
			},
		},
		{
			name:         "Order 501",
			itemsOrdered: 501,
			expectedResult: map[int]int{
				500: 1,
				250: 1,
			},
		},
		{
			name:         "Order 12001",
			itemsOrdered: 12001,
			expectedResult: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
		},
		{
			name:           "Invalid",
			itemsOrdered:   0,
			expectedResult: map[int]int{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			result := CalculatePacks(tt.itemsOrdered)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("CalculatePacks(%d) = %v; want %v",
					tt.itemsOrdered, result, tt.expectedResult)
			}
		})
	}
}
