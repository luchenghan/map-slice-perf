package mapsliceperf

import (
	"fmt"
	"testing"
)

const sliceItemCount = 10 // Placeholder for the number of slice elements

func insertXSlice(itemsCount int, b *testing.B) {
	testSlice := []int{}

	for i := 0; i < itemsCount; i += 1 {
		testSlice = append(testSlice, i)
		fmt.Println(testSlice)
	}
}

func createArray(arraySize int) []int {
	testArray := make([]int, arraySize+1)

	for i := 1; i <= arraySize; i += 1 {
		testArray[i] = i
	}

	return testArray
}

func BenchmarkCallSliceValues(b *testing.B) {
	sliceIndexes := [...]int{10, 40, 152, 201, 3232, 7512, 10012, 23412, 64123, 655123}
	testSlice := createArray(655123)
	for _, v := range sliceIndexes {

		b.Run(fmt.Sprintf("index_%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, tmpValue := range testSlice {
					if tmpValue == v {
						break
					}
				}
			}
		})
	}
}

func BenchmarkInsertXSlice(b *testing.B) {
	b.Run(fmt.Sprintf("item_size_%d", sliceItemCount), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			insertXSlice(sliceItemCount, b)
		}
	})
}
