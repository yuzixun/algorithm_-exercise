package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func randList() []int {
	var randList []int

	// rand.Seed(time.Now().UnixNano())
	for index := 0; index < 100; index++ {
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
	l := *list
	fmt.Println(l)
	listLen := len(l)

	for gap := (listLen / 2); gap > 0; gap /= 2 {
		for i := gap; i < listLen; i++ {
			j := i
			for j-gap >= 0 && l[j] < l[j-gap] {
				l[j], l[j-gap] = l[j-gap], l[j]
				j -= gap
			}
		}
	}
	fmt.Println(l)
}

func TestMergeSort(t *testing.T) {
	list := randList()
	fmt.Println("xxx", MergeSort(list))

}

func MergeSort(arr []int) []int {
	listLen := len(arr)

	if listLen == 1 {
		return arr
	}
	middle := int(listLen / 2)
	var left = make([]int, middle)
	var right = make([]int, listLen-middle)

	for i := 0; i < listLen; i++ {
		if i < middle {
			left[i] = arr[i]
		} else {
			right[i-middle] = arr[i]
		}
	}
	result := Merge(MergeSort(left), MergeSort(right))
	return result
}

func Merge(left []int, right []int) (result []int) {
	result = make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}
