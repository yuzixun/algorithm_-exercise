package sort

import (
	"fmt"
)

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

func QuickSort(array []int, left, right int) []int {
	if left < right {
		q := partition(array, left, right)

		QuickSort(array, left, q-1)
		QuickSort(array, q+1, right)
	}
	return array
}

func partition(array []int, left, right int) int {
	x := array[right]
	i := left - 1

	for j := left; j <= right-1; j++ {
		if array[j] < x {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	i++
	array[i], array[right] = array[right], array[i]
	return i
}

func HeapSort(array []int) {
	buildMaxHeap(&array)
	fmt.Println(array)
	size := len(array)
	for i := size - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		size--
		heapify(&array, 0, size)
	}
	fmt.Println("array is ", array)
}

func buildMaxHeap(pArray *[]int) {
	// array := *pArray
	size := len(*pArray)

	for i := size / 2; i >= 0; i-- {
		heapify(pArray, i, size)
	}
}

func heapify(pArray *[]int, c, size int) {
	array := (*pArray)
	root := c
	left := c*2 + 1
	right := c*2 + 2

	if left < size && array[left] > array[root] {
		root = left
	}

	if right < size && array[right] > array[root] {
		root = right
	}

	if root != c {
		array[c], array[root] = array[root], array[c]
		heapify(pArray, root, size)
	}
}

func CountingSort(array []int) {
	size := len(array)
	max := getMax(array)
	countArray := make([]int, max+1)
	result := make([]int, size)

	for _, element := range array {
		countArray[element]++
	}

	for i := 1; i <= max; i++ {
		countArray[i] += countArray[i-1]
	}

	for _, v := range array {
		countArray[v]--
		result[countArray[v]] = v
	}
	fmt.Println(result)
}

func getMax(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return max
}

func BucketSort(array []int) {
	num := len(array)
	max := getMax(array)
	buckets := make([][]int, num)

	index := 0
	for i := 0; i < num; i++ {
		index = array[i] * (num - 1) / max
		buckets[index] = append(buckets[index], array[i])
	}

	pos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])

		if bucketLen > 0 {
			insertionSortInBucket(buckets[i])
			copy(array[pos:], buckets[i])
			pos += bucketLen
		}
	}
	fmt.Println(array)
}

func insertionSortInBucket(array []int) {
	size := len(array)
	for i := 1; i < size; i++ {
		j := i - 1
		current := array[i]

		for j >= 0 && array[j] > current {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = current
	}
}

func RadixSort(array []int) {
	size := len(array)
	max := getMax(array)
	divider, reminder := 10, 1

	for bit := 1; max/bit > 0; bit *= 10 {
		bucket := make([][]int, 10)
		for j := 0; j < size; j++ {
			index := (array[j] % divider) / reminder
			bucket[index] = append(bucket[index], array[j])
		}

		curPos := 0
		for j := 0; j < len(bucket); j++ {
			for k := 0; k < len(bucket[j]); k++ {
				array[curPos] = bucket[j][k]
				curPos++
			}
		}

		divider *= 10
		reminder *= 10
	}
	fmt.Println("array is ", array)
}
