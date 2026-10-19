// combinatorics
package main

import (
	"fmt"
	"slices"
)

// some formulas
func NumberOfSubsets(n int) int { return 1 << n } // 2^n
func NumberOfPairs(n int) int   { return n * (n - 1) / 2 }
func A(n int, k int) int        { return Factorial(n) / Factorial(n-k) }
func C2(n int, k int) int       { return Factorial(n) / (Factorial(k) * Factorial(n-k)) }
func C3(n, k int) int {
	if k > n {
		return 0
	}
	if k == n || k == 0 {
		return 1
	}
	return C3(n-1, k-1) + C3(n-1, k)
}

func C(n int, k int, mod int) int {
	if k > n {
		return 0 // panic(fmt.Sprintf("math.C(): k > n (%v > %v)", k, n))
	}
	if k == n || k == 0 {
		return 1
	}

	// n! / (k! * (n-k)!)
	nFactDivNKFact := 1 // n! / (n-k)!
	for i := n - k + 1; i <= n; i++ {
		nFactDivNKFact = (nFactDivNKFact * i) % mod
	}
	kFact := 1 // k!
	for i := 2; i <= k; i++ {
		kFact = (kFact * i) % mod
	}
	return (nFactDivNKFact * ModInv(kFact, mod)) % mod
}

func PascalsTriangle(n int, mod int) [][]int {
	triangle := make([][]int, n)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j, e := 1, i/2; j <= e; j++ {
			triangle[i][j] = (triangle[i-1][j-1] + triangle[i-1][j]) % mod
			triangle[i][i-j] = triangle[i][j]
		}
	}
	return triangle
}

func Combinations(n int, k int) [][]int {
	// time: O(C(n,k)*k), space: O(n) / O(C(n,k)*k) ?

	if k > n {
		return nil
	}

	combinations := [][]int{}
	combination := make([]int, k)

	var combine func(pos int, start int)
	combine = func(pos int, start int) {
		if pos == k {
			combinations = append(combinations, slices.Clone(combination))
			return
		}

		for num := start; num < n; num++ {
			combination[pos] = num
			combine(pos+1, num+1)
		}
	}

	combine(0, 0)
	return combinations
}

func UniquePermutations(nums []int) [][]int {
	// k-permutations of N (partial permutation)

	numsLen := len(nums)
	permutations := [][]int{}

	var permute func(index int)
	permute = func(index int) {
		if index == numsLen {
			permutations = append(permutations, slices.Clone(nums))
			return
		}

		seen := map[int]bool{}
		for i := index; i < numsLen; i++ {
			if !seen[nums[i]] {
				nums[index], nums[i] = nums[i], nums[index]
				permute(index + 1)
				nums[index], nums[i] = nums[i], nums[index]
				seen[nums[i]] = true
			}
		}
	}

	permute(0)
	return permutations
}

func UniquePermutations0(nums []int) [][]int {
	// k-permutations of N (partial permutation)

	numsLen := len(nums)
	permutations := [][]int{}
	permutation := make([]int, numsLen)
	numsLeft := map[int]int{}
	for _, num := range nums {
		numsLeft[num]++
	}

	var permute func(index int)
	permute = func(index int) {
		if index == numsLen {
			permutations = append(permutations, slices.Clone(permutation))
			return
		}

		for num, left := range numsLeft {
			if left > 0 {
				permutation[index] = num
				numsLeft[num]--
				permute(index + 1)
				numsLeft[num]++
			}
		}
	}

	permute(0)
	return permutations
}

func Permutations(n int) [][]int {
	// time: O(n!*n), space: O(n) ?

	permutations := make([][]int, 0, Factorial(n))
	permutation := make([]int, n)
	for i := range permutation {
		permutation[i] = i
	}

	var permute func(pos int)
	permute = func(pos int) {
		if pos == n {
			permutations = append(permutations, slices.Clone(permutation))
			return
		}

		for i := pos; i < n; i++ {
			permutation[i], permutation[pos] = permutation[pos], permutation[i]
			permute(pos + 1)
			permutation[i], permutation[pos] = permutation[pos], permutation[i]
		}
	}

	permute(0)
	return permutations
}

func Permutations0(n int) [][]int {
	// time: O(n!*n), space: O(n) ?

	permutations := make([][]int, 0, Factorial(n))
	permutation := make([]int, n)
	usedIndices := make([]bool, n)

	var permute func(pos int)
	permute = func(pos int) {
		if pos == n {
			permutations = append(permutations, slices.Clone(permutation))
			return
		}

		for index, used := range usedIndices {
			if !used {
				permutation[pos] = index
				usedIndices[index] = true
				permute(pos + 1)
				usedIndices[index] = false
			}
		}
	}

	permute(0)
	return permutations
}

func NextPermutation(nums []int) {
	// time: O(n), space: O(1)

	numsLen := len(nums)
	i := numsLen - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		left, right := i+1, numsLen-1
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] > nums[i] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		nums[i], nums[right] = nums[right], nums[i]
	}
	for left, right := i+1, numsLen-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
}

func testCombinations() {
	strs := []string{"a", "b", "c"}
	for _, indices := range Combinations(len(strs), 2) {
		for _, index := range indices {
			fmt.Printf("%v ", strs[index])
		}
		fmt.Println()
	}

	return
	combinations := Combinations(10, 3)
	for _, combination := range combinations {
		fmt.Println(combination)
	}
}
