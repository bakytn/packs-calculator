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
		packSizes      []int
	}{
		{
			name:           "Empty pack sizes",
			itemsOrdered:   1,
			expectedResult: nil,
			packSizes:      []int{},
		},
		{
			name:           "Nil pack sizes",
			itemsOrdered:   1,
			expectedResult: nil,
			packSizes:      nil,
		},
		{
			name:         "Order 1",
			itemsOrdered: 1,
			expectedResult: []PackCount{
				{PackSize: 250, Count: 1},
			},
			packSizes: []int{10000011, 5000, 2000, 1000, 500, 250},
		},
		{
			name:         "Order 250 default sizes",
			itemsOrdered: 250,
			expectedResult: []PackCount{
				{PackSize: 250, Count: 1},
			},
			packSizes: WithDefaultPackSizes(),
		},
		{
			name:         "Order 251 default sizes",
			itemsOrdered: 251,
			expectedResult: []PackCount{
				{PackSize: 500, Count: 1},
			},
			packSizes: WithDefaultPackSizes(),
		},
		{
			name:         "Order 501",
			itemsOrdered: 501,
			expectedResult: []PackCount{
				{PackSize: 250, Count: 1},
				{PackSize: 500, Count: 1},
			},
			packSizes: []int{10000011, 5000, 2000, 1000, 500, 250},
		},
		{
			name:         "Order 12001",
			itemsOrdered: 12001,
			expectedResult: []PackCount{
				{PackSize: 250, Count: 1},
				{PackSize: 2000, Count: 1},
				{PackSize: 5000, Count: 2},
			},
			packSizes: []int{10000011, 5000, 2000, 1000, 500, 250},
		},
		{
			name:           "Invalid",
			itemsOrdered:   0,
			expectedResult: nil,
			packSizes:      []int{10000011, 5000, 2000, 1000, 500, 250},
		},
		{
			name:         "Ten million orders",
			itemsOrdered: 10000000,
			expectedResult: []PackCount{
				{PackSize: 5000, Count: 2000},
			},
			packSizes: []int{10000011, 5000, 2000, 1000, 500, 250},
		},
		{
			name:         "More than 10 million orders",
			itemsOrdered: 10000001,
			expectedResult: []PackCount{
				{PackSize: 250, Count: 1},
				{PackSize: 5000, Count: 2000},
			},
			packSizes: []int{10000011, 5000, 2000, 1000, 500, 250},
		},
		{
			name:         "10000011 million orders",
			itemsOrdered: 10000012,
			expectedResult: []PackCount{
				{PackSize: 250, Count: 1},
				{PackSize: 10000011, Count: 1},
			},
			packSizes: []int{10000011, 5000, 2000, 1000, 500, 250},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			result := CalculatePacks(tt.itemsOrdered, tt.packSizes)

			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("CalculatePacks(%d) = %v; want %v",
					tt.itemsOrdered, result, tt.expectedResult)
			}
		})
	}
}
