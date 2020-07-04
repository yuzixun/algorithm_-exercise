package main

func isUnique(astr string) bool {
	helper := 0

	for _, item := range astr {
		value := 1 << uint(item-'a')
		if helper&value != 0 {
			return false
		} else {
			helper |= value
		}
	}

	return true
}
