package main

type TripleInOne struct {
	array     []int
	current   [3]int
	stackSize int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		array:     make([]int, stackSize*3),
		current:   [3]int{-1, -1, -1},
		stackSize: stackSize,
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	cur := this.current[stackNum%3]
	cur++
	if cur >= this.stackSize {
		return
	}
	this.current[stackNum%3] = cur
	//fmt.Println("push ", this.current)
	this.array[cur*3+stackNum%3] = value
}

func (this *TripleInOne) Pop(stackNum int) int {

	cur := this.current[stackNum%3]
	if cur == -1 {
		return -1
	}
	this.current[stackNum%3]--
	//fmt.Println("ppp", cur, this.current, this.array)
	return this.array[cur*3+stackNum%3]
}

func (this *TripleInOne) Peek(stackNum int) int {
	cur := this.current[stackNum%3]
	if cur == -1 {
		return -1
	}
	return this.array[cur*3+stackNum%3]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.current[stackNum%3] == -1
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */

func main() {

}
