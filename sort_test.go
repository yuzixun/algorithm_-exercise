package main

import (
	"math/rand"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	list := randList()
	InsertionSort(&list)
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

func randList() []int {
	var randList []int

	// rand.Seed(time.Now().UnixNano())

	for index := 0; index < 100000; index++ {
		randList = append(randList, rand.Intn(10000))
	}

	return randList
}
