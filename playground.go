package main

import "fmt"

type Phone interface {
	call()
}

type iPhone struct {
	name string
}

func (phone iPhone) call() {
	fmt.Println("hello iphone")
}

func (phone iPhone) send_wechat() {
	fmt.Println("hello wechat")
}

func main() {
	a := 10

	ii := interface{}(a)
	switch ii.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
}
