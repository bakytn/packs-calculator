package packs

import (
	"sort"
)

var defaultPackSizes = []int{5000, 2000, 1000, 500, 250}

func WithDefaultPackSizes() []int {
	sort.Slice(defaultPackSizes, func(i, j int) bool {
		return defaultPackSizes[i] > defaultPackSizes[j]
	})

	return defaultPackSizes
}

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

func CalculatePacks(order int, packSizes []int) map[int]int {
	if packSizes == nil {
		return nil
	}

	if order == 0 {
		return nil
	}

	result := make(map[int]int)

	remainingOrder := order
	for i, pack := range packSizes {
		packs := remainingOrder / pack
		remainder := remainingOrder % pack

		if packs < 1 && remainder == remainingOrder {
			if i+1 == len(packSizes) {
				result[pack] += 1

				return result
			}

			nextPack := packSizes[i+1]

			if remainder > nextPack {
				if i+2 == len(packSizes) {
					// we assign current pack
					result[pack] += 1

					return result
				}
			}

			continue
		}

		if packs > 0 && remainingOrder > remainder && remainingOrder > 0 {
			result[pack] += packs

			r := CalculatePacks(remainder, packSizes[i+1:])
			mergePacks(r, result)

			return result
		}

		if packs > 0 && remainder == 0 {
			result[pack] += 1

			return result
		}
	}

	return result
}

func mergePacks(src map[int]int, dst map[int]int) {
	for k, v := range src {
		dst[k] += v
	}
}
