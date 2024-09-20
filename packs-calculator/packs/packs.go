package packs

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var defaultPackSizes = []int{250, 500, 1000, 2000, 5000}

type Packs []int

func (p Packs) Len() int           { return len(p) }
func (p Packs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Packs) Less(i, j int) bool { return p[i] < p[j] }

// PackCount represents the count of each pack size to send
type PackCount struct {
	PackSize int `json:"pack_size"`
	Count    int `json:"count"`
}

// Define a struct to store combination details
type Combination struct {
	TotalItems int
	TotalPacks int
	PackCounts map[int]int
}

// ParsePackSizes parses the comma-separated pack sizes string into a sorted slice of integers
func ParsePackSizes(packs string) ([]int, error) {
	packStrs := strings.Split(packs, ",")
	packInts := make([]int, len(packStrs))
	for i, s := range packStrs {
		packSize, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil || packSize <= 0 {
			return nil, fmt.Errorf("invalid pack size '%s'", s)
		}
		packInts[i] = packSize
	}
	// Sort the pack sizes in ascending order
	sort.Ints(packInts)
	return packInts, nil
}

// CalculatePacks calculates the packs needed to fulfill the order
func CalculatePacks(order int) map[int]int {
	if order <= 0 {
		return map[int]int{}
	}

	maxPackSize := defaultPackSizes[len(defaultPackSizes)-1]
	maxTotalItems := order + maxPackSize

	// Initialize a map to store the best combination for each total number of items
	combinations := make(map[int]*Combination)
	combinations[0] = &Combination{
		TotalItems: 0,
		TotalPacks: 0,
		PackCounts: make(map[int]int),
	}

	// Build combinations
	for totalItems := 1; totalItems <= maxTotalItems; totalItems++ {
		for _, packSize := range defaultPackSizes {
			if packSize > totalItems {
				break
			}
			remaining := totalItems - packSize
			if prevComb, exists := combinations[remaining]; exists {
				// Build new combination
				newPackCounts := make(map[int]int)
				for k, v := range prevComb.PackCounts {
					newPackCounts[k] = v
				}
				newPackCounts[packSize]++

				newComb := &Combination{
					TotalItems: totalItems,
					TotalPacks: prevComb.TotalPacks + 1,
					PackCounts: newPackCounts,
				}

				// Check if this combination is better
				existingComb, exists := combinations[totalItems]
				if !exists || isBetterCombination(newComb, existingComb) {
					combinations[totalItems] = newComb
				}
			}
		}
	}

	// Find the best combination starting from the order quantity
	var bestComb *Combination
	for totalItems := order; totalItems <= maxTotalItems; totalItems++ {
		if comb, exists := combinations[totalItems]; exists {
			if bestComb == nil || isBetterCombination(comb, bestComb) {
				bestComb = comb
			}
		}
	}

	if bestComb != nil {
		return bestComb.PackCounts
	}

	// If no combination found, select the smallest pack size larger than order
	for _, packSize := range defaultPackSizes {
		if packSize >= order {
			return map[int]int{packSize: 1}
		}
	}

	// As a last resort, return the largest pack size
	return map[int]int{defaultPackSizes[len(defaultPackSizes)-1]: 1}
}

// isBetterCombination determines if comb1 is better than comb2
func isBetterCombination(comb1, comb2 *Combination) bool {
	if comb1.TotalItems < comb2.TotalItems {
		return true
	}
	if comb1.TotalItems == comb2.TotalItems && comb1.TotalPacks < comb2.TotalPacks {
		return true
	}
	return false
}

// max returns the maximum value among the provided integers
func max(nums ...int) int {
	maxVal := nums[0]
	for _, n := range nums {
		if n > maxVal {
			maxVal = n
		}
	}
	return maxVal
}
