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
		expectedResult []PackCount
		errorExpected  bool
	}{
		{
			name:         "Order 1",
			itemsOrdered: 1,
			expectedResult: []PackCount{
				{PackSize: 250, Count: 1},
			},
		},
		{
			name:         "Order 250",
			itemsOrdered: 250,
			expectedResult: []PackCount{
				{PackSize: 250, Count: 1},
			},
		},
		{
			name:         "Order 251",
			itemsOrdered: 251,
			expectedResult: []PackCount{
				{PackSize: 500, Count: 1},
			},
		},
		{
			name:         "Order 501",
			itemsOrdered: 501,
			expectedResult: []PackCount{
				{PackSize: 250, Count: 1},
				{PackSize: 500, Count: 1},
			},
		},
		{
			name:         "Order 12001",
			itemsOrdered: 12001,
			expectedResult: []PackCount{
				{PackSize: 250, Count: 1},
				{PackSize: 2000, Count: 1},
				{PackSize: 5000, Count: 2},
			},
		},
		{
			name:           "Invalid",
			itemsOrdered:   0,
			expectedResult: nil,
		},
		{
			name:         "Ten million orders",
			itemsOrdered: 10000000,
			expectedResult: []PackCount{
				{PackSize: 5000, Count: 2000},
			},
		},
		{
			name:         "More than 10 million orders",
			itemsOrdered: 10000001,
			expectedResult: []PackCount{
				{PackSize: 250, Count: 1},
				{PackSize: 5000, Count: 2000},
			},
		},
		{
			name:         "10000011 million orders",
			itemsOrdered: 10000012,
			expectedResult: []PackCount{
				{PackSize: 250, Count: 1},
				{PackSize: 10000011, Count: 1},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			result := CalculatePacks(tt.itemsOrdered, []int{10000011, 5000, 2000, 1000, 500, 250})

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("CalculatePacks(%d) = %v; want %v",
					tt.itemsOrdered, result, tt.expectedResult)
			}
		})
	}
}
