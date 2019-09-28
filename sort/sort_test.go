package sort

import (
	"math/rand"
	"testing"
)

func randList() []int {
	var randList []int

	// rand.Seed(time.Now().UnixNano())
	for index := 0; index < 100; index++ {
		randList = append(randList, rand.Intn(1000000))
	}

	return randList
}

// func BenchmarkInsertionSort(b *testing.B) {
// 	b.StopTimer()
// 	list := randList()
// 	b.StartTimer()
// 	InsertionSort(&list)
// }

// func BenchmarkBubbleSort(b *testing.B) {
// 	b.StopTimer()
// 	list := randList()
// 	b.StartTimer()
// 	BubbleSort(&list)
// }

// func TestBubbleSort(t *testing.T) {
// 	list := randList()
// 	BubbleSort(&list)
// }
// func TestSelectionSort(t *testing.T) {
// 	list := randList()
// 	SelectionSort(&list)
// }

// func BenchmarkSelectionSort(b *testing.B) {
// 	b.StopTimer()
// 	list := randList()
// 	b.StartTimer()
// 	SelectionSort(&list)
// }

// func TestShellSort(t *testing.T) {
// 	list := randList()
// 	ShellSort(&list)
// }

// func BenchmarkShellSort(b *testing.B) {
// 	b.StopTimer()
// 	list := randList()
// 	b.StartTimer()
// 	ShellSort(&list)
// }

// func TestMergeSort(t *testing.T) {
// 	list := randList()
// 	MergeSort(list)
// }

// func TestQuickSort(t *testing.T) {
// 	list := randList()
// 	fmt.Println(QuickSort(list, 0, len(list)-1))
// }

// func TestHeapSort(t *testing.T) {
// 	list := randList()
// 	HeapSort(list)
// }

// func TestCountingSort(t *testing.T) {
// 	list := randList()
// 	CountingSort(list)
// }

// func TestBucketSort(t *testing.T) {
// 	list := randList()
// 	BucketSort(list)
// }

func TestRadixSort(t *testing.T) {
	list := randList()
	RadixSort(list)
}
