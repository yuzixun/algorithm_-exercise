package main

import "sort"

func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
	const eps = 1e-6

	sub := func(a, b []int) []int {
		c := make([]int, 2)
		c[0], c[1] = a[0]-b[0], a[1]-b[1]
		return c
	}

	det := func(a, b []int) int {
		return a[0]*b[1] - a[1]*b[0]
	}

	inline := func(s, v, p []int) bool {
		i := 0
		if v[i] == 0 {
			i = 1
		}
		t := float64(p[i]-s[i]) / float64(v[i])
		return -eps <= t && t <= 1+eps
	}

	v1 := sub(end1, start1)
	v2 := sub(end2, start2)
	b := sub(start2, start1)
	deno := det(v1, v2)
	ans := []float64{}

	if deno == 0 {
		if det(v1, b) == 0 {
			tmp := [][]int{start1, end1, start2, end2}
			sort.Slice(tmp, func(i, j int) bool {
				return tmp[i][0] < tmp[j][0] || (tmp[i][0] == tmp[j][0] && tmp[i][1] < tmp[j][1])
			})
			if inline(start1, v1, tmp[1]) && inline(start2, v2, tmp[1]) {
				ans = []float64{float64(tmp[1][0]), float64(tmp[1][1])}
			}
		}

	} else {
		t1 := float64(det(b, v2)) / float64(deno)
		t2 := -float64(det(v1, b)) / float64(deno)

		if 0 <= t1 && t1 <= 1 && 0 <= t2 && t2 <= 1 {
			ans = []float64{
				float64(start1[0]) + float64(v1[0])*t1,
				float64(start1[1]) + float64(v1[1])*t1,
			}
		}
	}
	return ans
}

func main() {

}
