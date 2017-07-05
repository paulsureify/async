package main

import (
	"fmt"
	"time"
)

const numberOfRecordsToProcess = 10

func main() {

	// Get a list of user records
	x := listx(80)

	// go run those tasks

	fmt.Println(x)
	y := getChunkedNumber(x)

	// ch1 := make(chan bool)
	start := time.Now()
	// processDividentNumbers(y)
	processWithoutMultiThreading(y)
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s\n", elapsed)
	// <-ch1
}

func listx(x int) []int {
	arr := make([]int, x)
	for i := range arr {
		arr[i] = i + 1
	}

	return arr
}

func getChunkedNumber(arr []int) [][]int {
	var divided [][]int
	chunkSize := numberOfRecordsToProcess % (len(arr))

	for i := 0; i < len(arr); i += chunkSize {
		end := i + chunkSize

		if end > len(arr) {
			end = len(arr)
		}

		divided = append(divided, arr[i:end])
	}

	return divided
	// fmt.Printf("%#v\n", divided)
}

func processWithoutMultiThreading(y [][]int) {
	for i := 0; i < len(y); i++ {
		doThisWithoutMultiThreading(y[i])
	}

}

func doThisWithoutMultiThreading(arr []int) {
	// Doing this sleep to simulate real time processing
	time.Sleep(time.Duration(2000 * time.Millisecond))
	fmt.Println(arr)
}
