package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right := 0, n
	for {
		mid := left + (right-left)/2
		result := guess(mid)
		if result == 0 {
			return mid
		} else if result == 1 {
			right = mid - 1
		} else if result == -1 {
			left = mid + 1
		}
	}
}
