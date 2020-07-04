package main

type StackOfPlates struct {
	cap    int
	stacks [][]int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{
		cap:    cap,
		stacks: make([][]int, 0),
	}

}

func (this *StackOfPlates) Push(val int) {
	if this.cap == 0 {
		return
	}

	if len(this.stacks) == 0 || len(this.stacks[len(this.stacks)-1]) >= this.cap {
		this.stacks = append(this.stacks, []int{val})
	} else {
		this.stacks[len(this.stacks)-1] = append(this.stacks[len(this.stacks)-1], val)
	}
}

func (this *StackOfPlates) Pop() int {
	if len(this.stacks) == 0 {
		return -1
	}
	lastStack := this.stacks[len(this.stacks)-1]
	if len(lastStack) == 0 {
		return -1
	}
	// fmt.Println("last stack", lastStack)
	result := lastStack[len(lastStack)-1]
	if len(lastStack) == 1 {
		this.stacks = this.stacks[:len(this.stacks)-1]
	} else {
		this.stacks[len(this.stacks)-1] = lastStack[:len(lastStack)-1]
	}
	// fmt.Println("pop ", this.stacks, lastStack[len(lastStack)-1])
	return result
}

func (this *StackOfPlates) PopAt(index int) int {
	if len(this.stacks) == 0 || index > len(this.stacks)-1 {
		return -1
	}

	currentStack := this.stacks[index]
	if len(currentStack) == 0 {
		return -1
	}
	result := currentStack[len(currentStack)-1]
	if len(currentStack) == 1 {
		this.stacks = append(this.stacks[:index], this.stacks[index+1:]...)
	} else {
		this.stacks[index] = currentStack[:len(currentStack)-1]
	}

	return result
}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
