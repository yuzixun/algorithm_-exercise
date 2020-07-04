package main

import "fmt"

func numOfWays(n int) int {
	mod := 1000000007
	same, diff := 6, 6
	for i := 1; i < n; i++ {
		nextSame := 3*same + 2*diff
		nextDiff := 2*same + 2*diff
		same, diff = (nextSame % mod), (nextDiff % mod)
	}

	return (same + diff) % mod
}

func main() {
	fmt.Println(numOfWays(1))
	fmt.Println(numOfWays(2))
	fmt.Println(numOfWays(3))
	fmt.Println(numOfWays(7))
	fmt.Println(numOfWays(5000))
}
