package main

import (
	"fmt"
	"math/bits"
	"math/rand/v2"
	"time"
)

func MakeMinSparseTable(nums []int) *SparseTable {
	return MakeSparseTable(nums, func(a, b int) int { return min(a, b) })
}
func MakeMaxSparseTable(nums []int) *SparseTable {
	return MakeSparseTable(nums, func(a, b int) int { return max(a, b) })
}
func MakeOrSparseTable(nums []int) *SparseTable {
	return MakeSparseTable(nums, func(a, b int) int { return a | b })
}

type SparseTable struct {
	items [][]int
	f     func(int, int) int
}

func MakeSparseTable(nums []int, f func(int, int) int) *SparseTable {
	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	}

	// levelCount := logs2[numsLen] + 1
	levelCount := bits.Len(uint(numsLen))
	items := make([][]int, levelCount)
	items[0] = make([]int, numsLen)
	copy(items[0], nums)
	for level := 1; level < levelCount; level++ {
		items[level] = make([]int, numsLen-(1<<level)+1)
		prevLevel := level - 1
		prevLevelLen := 1 << prevLevel
		for i := range items[level] {
			items[level][i] = f(items[prevLevel][i], items[prevLevel][i+prevLevelLen])
		}
	}
	return &SparseTable{
		items: items,
		f:     f}
}

// idempotent query (min/max, or, etc)
func (st *SparseTable) IdempQuery(left, right int) int {
	// l := st.logs2[right-left+1]
	l := bits.Len(uint(right-left+1)) - 1
	return st.f(st.items[l][left], st.items[l][right-(1<<l)+1])
}

// non-idempotent query (sum, xor etc)
func (st *SparseTable) NonIdempQuery(left, right int) int {
	sum := 0
	length := right - left + 1
	for level := 0; length != 0; level++ {
		if length&1 == 1 {
			sum = st.f(sum, st.items[level][left])
			left += 1 << level
		}
		length >>= 1
	}
	return sum
}

func calcLogs2(maxNum int) []int {
	logs2 := make([]int, maxNum+1)
	logs2[1] = 0
	for i := 2; i <= maxNum; i++ {
		logs2[i] = logs2[i>>1] + 1
	}
	return logs2
}

func bitLength(n int) int {
	if n == 0 {
		return 1
	}

	count := 0
	for n != 0 {
		count++
		n >>= 1
	}
	return count
}

func testSparseTableIdempQuery() {
	const minNum int = -1e5
	const maxNum int = 1e5
	const numsRange = maxNum - minNum + 1
	const numsLen int = 1e5
	const testCount int = 100
	const queryCount int = 100
	funcs := []func(int, int) int{
		func(a, b int) int { return min(a, b) },
		func(a, b int) int { return max(a, b) },
		func(a, b int) int { return a | b }}
	nums := make([]int, numsLen)
	t := time.Now()
	for range testCount {
		for i := range nums {
			nums[i] = rand.IntN(numsRange) + minNum
		}
		f := funcs[rand.IntN(len(funcs))]
		st := MakeSparseTable(nums, f)
		for range queryCount {
			length := rand.IntN(numsLen-1) + 1
			left := rand.IntN(numsLen)
			right := min(left+length-1, numsLen-1)
			res1 := st.IdempQuery(left, right)
			res2 := agr(nums[left:right+1], f)
			if res1 != res2 {
				fmt.Printf("%v != %v: [%v, %v]\n", res1, res2, left, right)
			}
		}
	}
	fmt.Printf("testSparseTableIdempQuery time: %v\n", time.Since(t))
}

func testSparseTableNonIdempQuery() {
	const minNum int = -1e2
	const maxNum int = 1e2
	const numsRange = maxNum - minNum + 1
	const numsLen int = 1e5
	const testCount int = 100
	const queryCount int = 500
	nums := make([]int, numsLen)
	sum := func(a, b int) int { return a + b }
	t := time.Now()
	for range testCount {
		for i := range nums {
			nums[i] = rand.IntN(numsRange) + minNum
		}
		st := MakeSparseTable(nums, sum)
		for range queryCount {
			length := rand.IntN(numsLen-1) + 1
			left := rand.IntN(numsLen)
			right := min(left+length-1, numsLen-1)
			res1 := st.NonIdempQuery(left, right)
			res2 := agr(nums[left:right+1], sum)
			if res1 != res2 {
				fmt.Printf("%v != %v: [%v, %v]\n", res1, res2, left, right)
			}
		}
	}
	fmt.Printf("testSparseTableNonIdempQuery time: %v\n", time.Since(t))
}

func agr(nums []int, f func(int, int) int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		res = f(res, num)
	}
	return res
}

func main() {
	// logs := make([]int, 22)
	// logs[1] = 0
	// for i := 2; i < len(logs); i++ {
	// 	logs[i] = logs[i>>1] + 1
	// }
	// for i, log := range logs {
	// 	fmt.Printf("%v: %v\n", i, log)
	// }

	// nums := []int{1, 2, 3, 4, 5, 6, 7}
	// numsLen := len(nums)
	// st := MakeSparseTable(nums, func(a, b int) int { return min(a, b) })
	// fmt.Println(nums)
	// fmt.Println(st.Query(0, numsLen-1))
	// return

	testSparseTableIdempQuery()
	testSparseTableNonIdempQuery()
}

//            l             r
//         0  1  2  3  4  5 6
// 0 (1):  1  2  3  4  5  6 7
// 1 (2):  3  5  7  9 11 13
// 2 (4): 10 14 18 22
// 3 (8): X

//   l       r
// 1 2 3 4 5 6 7
// length: 5
// 101

// ln
// 1: 0
// 2: 1
// 3: 1
// 4: 2
// 5: 2
// 6: 2
// 7: 2
// 8: 3
// 9: 3
// 10: 3
// 11: 3
// 12: 3
// 13: 3
// 14: 3
// 15: 3
// 16: 4

// 0..2^3-1
// 0..7
// 0..2^(3-1)-1 + 2^(3-1)..2^3-1
// 0..2^2-1 + 2^2..2^3-1
// 0..3 + 4..7

// 0..2^4-1
// 0..2^3-1 + 2^3..2^4-1

// 0..7
// 0..1 + 2..3 + 4..7

// 0..15
// 0..1 + 2..3 + 4..7 + 8..15

// 22 = 2 + 2^2 + 2^4 = 2 + 4 + 16 = 22
// 10110

// 2^i
// 2^(i-1)

// 011
// 101
// 110
