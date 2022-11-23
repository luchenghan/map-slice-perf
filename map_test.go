package mapsliceperf

import (
	"fmt"
	"testing"
)

const mapItemCount = 10 // Placeholder for the number of map elements

func insertXMap(itemsCount int, b *testing.B) {
	testMap := map[int]int{}

	for i := 0; i < itemsCount; i += 1 {
		testMap[i] = i
	}
}

func createMap(mapSize int) map[int]int {
	testMap := map[int]int{}

	for i := 0; i < mapSize; i += 1 {
		testMap[i] = i
	}

	return testMap
}

func BenchmarkCallMapValues(b *testing.B) {
	mapKeys := []int{10, 40, 152, 201, 3232, 7512, 10012, 23412, 64123, 655123}
	testMap := createMap(mapItemCount)
	for _, v := range mapKeys {

		b.Run(fmt.Sprintf("index_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = testMap[v]
			}
		})
	}
}

func BenchmarkInsertXMap(b *testing.B) {
	b.Run(fmt.Sprintf("item_size_%d", mapItemCount), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			insertXMap(mapItemCount, b)
		}
	})
}
