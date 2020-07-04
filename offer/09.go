package main

type CQueue struct {
	stack1, stack2 []int
}

// func Constructor() CQueue {
// 	return CQueue{}
// }

func (this *CQueue) AppendTail(value int) {
	this.stack1 = append(this.stack1, value)
}

func (this *CQueue) DeleteHead() int {
	len1, len2 := len(this.stack1), len(this.stack2)

	if len1 == 0 && len2 == 0 {
		return -1
	}

	if len2 == 0 {
		for len1 > 0 {
			len1--
			this.stack2 = append(this.stack2, this.stack1[len1])
			len2++
		}
		this.stack1 = this.stack1[:len1]
	}
	len2--
	result := this.stack2[len2]
	this.stack2 = this.stack2[:len2]

	return result
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

// func main() {
// 	obj := Constructor()
// 	fmt.Println(obj.DeleteHead())
// 	obj.AppendTail(5)
// 	obj.AppendTail(2)
// 	fmt.Println(obj.DeleteHead())
// 	fmt.Println(obj.DeleteHead())

// }
