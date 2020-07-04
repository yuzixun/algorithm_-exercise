package main

import (
	"container/heap"
	"sort"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	members := make([][]int, n)
	for i := 0; i < n; i++ {
		members[i] = []int{efficiency[i], speed[i]}
	}
	sort.Slice(members[:], func(i, j int) bool {
		for x := range members[i] {

			if members[i][x] == members[j][x] {
				continue
			}
			return members[i][x] > members[j][x]
		}
		return false
	})

	h := &IntHeap{}
	heap.Init(h)

	var result, cur int64
	speedSum := 0
	for i := 0; i < n; i++ {
		speedSum += members[i][1]
		heap.Push(h, members[i][1])

		if h.Len() > k {
			speedSum -= heap.Pop(h).(int)
		}

		cur = int64(speedSum) * int64(members[i][0])
		if result < cur {
			result = cur
		}
	}

	// fmt.Println(result)
	return int(result % 1000000007)
}

func main() {
	maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 2)
	maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 3)
	maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 4)
	maxPerformance(5, []int{10, 10, 7, 9, 8}, []int{9, 8, 3, 6, 9}, 1)
}

// # @param {Integer} n

// # @param {Integer[]} speed
// # @param {Integer[]} efficiency
// # @param {Integer} k
// # @return {Integer}
// require "rubygems"

// require "algorithms"

// include Containers
// Member = Struct.new(:speed, :efficiency)

// def max_performance(n, speed, efficiency, k)
// 	members = []
// 	n.times do |i|
// 		members.push(Member.new(speed[i], efficiency[i]))
// 	end

// 	members.sort! { |a, b| a.efficiency == b.efficiency ? b.speed <=> a.speed : b.efficiency <=> a.efficiency }

// 	result = 0
// 	speedArr = MinHeap.new
// 	speedSum = 0
// 	members.each do |member|
// 		speedArr.push(member.speed)
// 		speedSum += member.speed
// 		if speedArr.size > k
// 			min = speedArr.pop
// 			speedSum -= min
// 		end
// 		result = [result, speedSum*member.efficiency].max
// 		# puts "#{result}, #{speedSum}, #{speedArr} "
// 	end

// 	result%(10**9 + 7)
// end

// puts max_performance(6, [2, 10, 3, 1, 5, 8], [5, 4, 3, 9, 7, 2], 2)
// puts max_performance(6, [2, 10, 3, 1, 5, 8], [5, 4, 3, 9, 7, 2], 3)
// puts max_performance(6, [2, 10, 3, 1, 5, 8], [5, 4, 3, 9, 7, 2], 4)
// puts max_performance(5, [10,10,7,9,8], [9,8,3,6,9], 1)
