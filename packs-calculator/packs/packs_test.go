package packs

import (
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	defaultPackSizes := []int{250, 500, 1000, 2000, 5000}

	// Test cases
	tests := []struct {
		name           string
		itemsOrdered   int
		packSizes      []int
		expectedResult map[int]int
	}{
		{
			name:         "Order 1",
			itemsOrdered: 1,
			packSizes:    defaultPackSizes,
			expectedResult: map[int]int{
				250: 1,
			},
		},
		{
			name:         "Order 250",
			itemsOrdered: 250,
			packSizes:    defaultPackSizes,
			expectedResult: map[int]int{
				250: 1,
			},
		},
		{
			name:         "Order 251",
			itemsOrdered: 251,
			packSizes:    defaultPackSizes,
			expectedResult: map[int]int{
				500: 1,
			},
		},
		{
			name:         "Order 501",
			itemsOrdered: 501,
			packSizes:    defaultPackSizes,
			expectedResult: map[int]int{
				500: 1,
				250: 1,
			},
		},
		{
			name:         "Order 12001",
			itemsOrdered: 12001,
			packSizes:    defaultPackSizes,
			expectedResult: map[int]int{
				5000: 2,
				2000: 1,
				250:  1,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			result := CalculatePacks(tt.itemsOrdered, tt.packSizes)

			// Compare the results
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("CalculatePacks(%d, %v) = %v; want %v",
					tt.itemsOrdered, tt.packSizes, result, tt.expectedResult)
			}
		})
	}
}
