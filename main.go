package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

var wg sync.WaitGroup

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	numbers := make([]int, size)
	for i := range numbers {
		numbers[i] = rand.Int()
	}
	return numbers
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]

	for i := range data {
		if max < data[i] {
			max = data[i]
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var maxData int
	sizeData := len(data) / CHUNKS
	if sizeData == 0{
		maxData = maximum(data)
		return maxData
	}
	chanksData := make([]int, CHUNKS)
	for i := range chanksData {
		wg.Add(1)
		start := i * sizeData
		end := start + sizeData
		if len(chanksData) == i+1 {
			end = len(data)
		}
		elements := data[start:end]
		go func(data []int) {
			defer wg.Done()
			max := maximum(data)
			chanksData[i] = max
		}(elements)
	}
	wg.Wait()

	maxData = maximum(chanksData)
	return maxData
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	s := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now()
	max := maximum(s)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)

	start = time.Now()
	max = maxChunks(s)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
