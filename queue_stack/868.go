package main

type MovingAverage struct {
	size int
	arr  []int
	sum  int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {

	return MovingAverage{
		size: size,
		arr:  []int{},
		sum:  0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.arr = append(this.arr, val)
	this.sum += val

	if len(this.arr) > this.size {
		this.sum -= this.arr[0]
		this.arr = this.arr[1:]
	}

	return float64(this.sum) / float64(len(this.arr))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

func main() {

}
