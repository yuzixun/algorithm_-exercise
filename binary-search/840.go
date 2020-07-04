package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left, right := 0, n
	for left != right {
		mid := (left + right) >> 1
		if !isBadVersion(mid) {
			if isBadVersion(mid + 1) {
				return mid
			}
			left = mid + 1
		} else {
			right = mid
		}
	}

	return -1
}
