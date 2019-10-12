package main

func trap(height []int) int {
	size := len(height)
	if size == 0 {
		return 0
	}

	result := 0
	leftMaxArr := make([]int, size, size)
	rightMaxArr := make([]int, size, size)

	for l := 0; l < size; l++ {
		leftMaxArr[l] = height[l]
		if l > 0 && leftMaxArr[l] < leftMaxArr[l-1] {
			leftMaxArr[l] = leftMaxArr[l-1]
		}
	}

	for r := size - 1; r >= 0; r-- {
		rightMaxArr[r] = height[r]
		if r < size-1 && rightMaxArr[r] < rightMaxArr[r+1] {
			rightMaxArr[r] = rightMaxArr[r+1]
		}
	}

	for i := 1; i < size; i++ {
		var currentHight int
		if rightMaxArr[i] > leftMaxArr[i] {
			currentHight = leftMaxArr[i]
		} else {
			currentHight = rightMaxArr[i]
		}
		// fmt.Println("i is ", i, "height is ", currentHight, height[i])
		if currentHight < height[i] {
			continue
		}
		result += (currentHight - height[i])
	}

	return result
}
