package main

type CQueue struct {
	primaryStack   []int
	secondaryStack []int
}

func Constructor() CQueue {
	return CQueue{
		primaryStack:   []int{},
		secondaryStack: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.primaryStack = append(this.primaryStack, value)
}

func (this *CQueue) DeleteHead() int {
	size1, size2 := len(this.primaryStack), len(this.secondaryStack)
	if size1 == 0 && size2 == 0 {
		return -1
	}

	if size2 == 0 {
		for size1 > 0 {
			size1--
			this.secondaryStack = append(this.secondaryStack, this.primaryStack[size1])
			size2++
		}

		this.primaryStack = this.primaryStack[:0]
	}

	size2--
	result := this.secondaryStack[size2]
	this.secondaryStack = this.secondaryStack[:size2]

	return result
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {

}
