package main

import (
	"fmt"
	"math/rand"
	"time"
)

const numberOfRecordsToProcess = 10

func main() {

	// Get a list of user records
	x := list(80)
	fmt.Println(x)
	y := getChunkedNumbers(x)

	start := time.Now()

	// go run those tasks
	processDividentNumbers(y)

	elapsed := time.Since(start)
	fmt.Printf("The process took %s\n", elapsed)
}

func list(x int) []int {
	arr := make([]int, x)
	for i := range arr {
		arr[i] = i + 1
	}

	return arr
}

func getChunkedNumbers(arr []int) [][]int {
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
}

func processDividentNumbers(y [][]int) {

	ch := make(chan string)
	for i := 0; i < len(y); i++ {
		go doThis(y[i], ch)
	}

	// Do this so the function can wait
	for i := 0; i < len(y); i++ {
		<-ch
	}
}

func doThis(arr []int, c chan<- string) {

	// Doing this sleep to simulate real time processing
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	fmt.Println(arr)
	c <- fmt.Sprintf("%d", len(arr))
}
