package main

func merge(A []int, m int, B []int, n int) {
	pointerA, pointerB, pointerM := m-1, n-1, m+n-1
	for pointerA >= 0 && pointerB >= 0 {
		if A[pointerA] > B[pointerB] {
			A[pointerM] = A[pointerA]
			pointerA--
		} else {
			A[pointerM] = B[pointerB]
			pointerB--
		}
		pointerM--
	}

	for pointerB >= 0 {
		A[pointerM] = A[pointerB]
		pointerB--
	}
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
