package main

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func randList() []int {
	var randList []int

	// rand.Seed(time.Now().UnixNano())
	for index := 0; index < 1000; index++ {
		randList = append(randList, rand.Intn(10000))
	}

	return randList
}

func BenchmarkInsertionSort(b *testing.B) {
	b.StopTimer()
	list := randList()
	b.StartTimer()
	InsertionSort(&list)
}

func InsertionSort(list *[]int) {
	listLen := len(*list)
	for i := 1; i < listLen; i++ {
		current := (*list)[i]
		j := i - 1
		for j >= 0 && (*list)[j] > current {
			(*list)[j+1] = (*list)[j]
			j--
		}
		(*list)[j+1] = current
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	b.StopTimer()
	list := randList()
	b.StartTimer()
	BubbleSort(&list)
}

func TestBubbleSort(t *testing.T) {
	// list := randList()
	// BubbleSort(&list)
}

func BubbleSort(list *[]int) {
	listLen := len(*list)
	fmt.Println(*list)
	for i := 0; i < listLen; i++ {
		for j := i; j < listLen; j++ {
			if (*list)[j] < (*list)[i] {
				(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
			}
		}
	}
	fmt.Println(*list)
}

func TestSelectionSort(t *testing.T) {
	list := randList()
	SelectionSort(&list)
}

func BenchmarkSelectionSort(b *testing.B) {
	b.StopTimer()
	list := randList()
	b.StartTimer()
	SelectionSort(&list)
}

func SelectionSort(list *[]int) {
	listLen := len(*list)

	for i := 0; i < listLen; i++ {
		minIndex := i
		for j := i + 1; j < listLen; j++ {
			if (*list)[minIndex] > (*list)[j] {
				minIndex = j
			}
		}
		(*list)[i], (*list)[minIndex] = (*list)[minIndex], (*list)[i]
	}
}

func TestShellSort(t *testing.T) {
	list := randList()
	ShellSort(&list)
}

func BenchmarkShellSort(b *testing.B) {
	b.StopTimer()
	list := randList()
	b.StartTimer()
	ShellSort(&list)
}

func ShellSort(list *[]int) {
	listLen := len(*list)

	for gap := math.Floor(float64(listLen / 2)); gap > 0; gap = math.Floor(gap / 2) {
		for i := int(gap); i < listLen; i++ {
			j := i
			current := (*list)[i]
			for j-int(gap) >= 0 && current < (*list)[j-int(gap)] {
				(*list)[j] = (*list)[j-int(gap)]
				j = j - int(gap)
			}
			(*list)[j] = current
		}
	}
}
