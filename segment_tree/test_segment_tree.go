package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"slices"
	"time"
)

type STTest struct {
	nums           []int
	stree          *SegmentTree[MinMaxSum, int]
	maxNum, minNum int
}

type MinMaxSum struct {
	min, max int
	sum      int
}

func MakeSTTest(numsSize int, minNum, maxNum int) *STTest {
	nums := GenerateNums(numsSize, minNum, maxNum)
	initNode := func(val int) MinMaxSum {
		return MinMaxSum{min: val, max: val, sum: val}
	}
	buildNode := func(leftChild, rightChild MinMaxSum, left, mid, right int) MinMaxSum {
		return MinMaxSum{
			min: min(leftChild.min, rightChild.min),
			max: max(leftChild.max, rightChild.max),
			sum: leftChild.sum + rightChild.sum}
	}
	return &STTest{
		nums:   nums,
		stree:  MakeSegmentTree(nums, initNode, buildNode),
		maxNum: maxNum,
		minNum: minNum}
}

func (t *STTest) TestQuery(left, right int) bool {
	node := t.stree.Query(left, right)
	minNum, maxNum := math.MaxInt, math.MinInt
	sum := 0
	for _, num := range t.nums[left : right+1] {
		minNum = min(minNum, num)
		maxNum = max(maxNum, num)
		sum += num
	}
	return minNum == node.min && maxNum == node.max && sum == node.sum
	// return MinRange(t.nums, left, right) == node.min &&
	// 	MaxRange(t.nums, left, right) == node.max &&
	// 	Sum(t.nums, left, right) == node.sum
}

func (t *STTest) Update(index int, value int) {
	t.nums[index] = value
	t.stree.Update(index, value)
}

// todo: cleanup

func testSegmentTree() {
	const numsLen = 1_000_000
	const minNum = -1_000_000
	const maxNum = 1_000_000
	const loopCount = 1_000

	t := time.Now()
	test := MakeSTTest(numsLen, minNum, maxNum)
	if !slices.Equal(test.nums, test.stree.Items(func(node MinMaxSum) int { return node.sum })) {
		fmt.Println("nums and ST are not equal")
	}
	for range loopCount {
		left := rand.IntN(numsLen)
		right := left + rand.IntN(numsLen-left)
		if !test.TestQuery(left, right) {
			fmt.Println("test query failed")
		}

		index := rand.IntN(numsLen)
		value := rand.IntN(maxNum-minNum+1) + minNum
		test.Update(index, value)
		if !test.TestQuery(index, index) {
			fmt.Println("test query failed after update")
		}
	}
	fmt.Printf("testSegmentTree() time: %v\n", time.Since(t))
}
