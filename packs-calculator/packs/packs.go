package packs

import (
	"sort"
)

var defaultPackSizes = []int{250, 2000, 1000, 500, 5000}

func WithDefaultPackSizes() []int {
	return defaultPackSizes
}

// PackCount represents the count of each pack size to send
type PackCount struct {
	PackSize int `json:"pack_size"`
	Count    int `json:"count"`
}

func sortDescPackSizes(packSizes []int) []int {
	if packSizes == nil {
		return nil
	}

	sortedPackSizesDesc := make([]int, len(packSizes))
	copy(sortedPackSizesDesc, packSizes)
	sort.Slice(sortedPackSizesDesc, func(i, j int) bool {
		return sortedPackSizesDesc[i] > sortedPackSizesDesc[j]
	})

	return sortedPackSizesDesc
}

func CalculatePacks(order int, packSizes []int) []PackCount {
	if order == 0 {
		return nil
	}

	packCounts := getPacks(order, sortDescPackSizes(packSizes))

	var packsToSend []PackCount
	for packSize, count := range packCounts {
		packsToSend = append(packsToSend, PackCount{PackSize: packSize, Count: count})
	}

	sort.Slice(packsToSend, func(i, j int) bool {
		return packsToSend[i].PackSize < packsToSend[j].PackSize
	})

	return packsToSend
}

func getPacks(order int, packSizes []int) map[int]int {
	if packSizes == nil {
		return nil
	}

	if order == 0 {
		return nil
	}

	result := make(map[int]int)

	for i, pack := range packSizes {
		packs := order / pack
		remainder := order % pack

		if packs < 1 && remainder == order {
			// if no packs next we just process current and return the result
			if i+1 == len(packSizes) {
				result[pack] += 1

				return result
			}

			// we need next pack to determine if we can assign current pack
			// instead of multiple smaller ones
			nextPack := packSizes[i+1]
			if remainder > nextPack {
				// peak next to assign current least biggest pack size
				// to avoid assigning multiple smaller ones in favor of bigger one
				if i+2 == len(packSizes) {
					// we assign current pack
					result[pack] += 1

					return result
				}
			}

			continue
		}

		if packs > 0 && order > remainder && order > 0 {
			result[pack] += packs

			r := getPacks(remainder, packSizes[i+1:])
			mergePacks(r, result)

			return result
		}

		// if we have packs and it divided perfectly we just return the result
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
