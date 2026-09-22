package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand/v2"
	"slices"
	"time"
)

type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

type Number interface{ Integer | float32 | float64 }

func Pow(base, exp int) int {
	// time: O(log exp), space: O(log exp)

	return PowMod(base, exp, math.MaxInt)

	// if exp == 0 {
	// 	return 1
	// } else if base == 0 || base == 1 {
	// 	return base
	// }

	// res := Pow(base*base, exp>>1)
	// if exp&1 == 1 {
	// 	res *= base
	// }
	// return res
}

func Pow2(base, exp int) int {
	// time: O(log exp), space: O(1)

	res := 1
	for exp != 0 {
		if exp&1 == 1 {
			res *= base
		}
		base *= base
		exp >>= 1
	}
	return res
}

func PowMod(base, exp, mod int) int {
	// time: O(log exp), space: O(log exp)

	if exp == 0 {
		return 1 % mod
	} else if base == 0 || base == 1 {
		return base % mod
	}

	res := PowMod((base*base)%mod, exp>>1, mod)
	if exp&1 == 1 {
		res = (res * base) % mod
	}
	return res
}

func PowMod2(base, exp, mod int) int {
	// time: O(log exp), space: O(1)

	res := 1
	for exp != 0 {
		if exp&1 == 1 {
			res = (res * base) % mod
		}
		base = (base * base) % mod
		exp >>= 1
	}
	return res % mod
}

func Sqrt(n int) int {
	// time: O(log n), space: O(1)

	if n <= 1 { // ~
		return n
	}

	left, right := 0, n // 46_341 // sqrt(2^31-1)
	for left <= right {
		mid := left + (right-left)/2
		if sqr := mid * mid; sqr > n {
			right = mid - 1
		} else if sqr < n {
			left = mid + 1
		} else {
			return mid
		}
	}
	return right
}

func IsPrime(n int) bool {
	if n <= 2 || n%2 == 0 {
		return n == 2
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func GetPrimeFactors(n int) []int {
	// time: O(sqrt(n)), space: O(1)

	primes := []int{}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			primes = append(primes, i)
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		primes = append(primes, n)
	}
	return primes
}

func GetPrimeFactors2(n int) []int {
	// time: O(sqrt(n)), space: O(1)

	primes := []int{}
	if n&1 == 0 {
		primes = append(primes, 2)
		for n&1 == 0 {
			n >>= 1
		}
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			primes = append(primes, i)
			for n%i == 0 {
				n /= i
			}
		}
	}
	if n > 1 {
		primes = append(primes, n)
	}
	return primes
}

// sieve of Eratosthenes
func Sieve(n int) []bool {
	// time: O(n log log n), space: O(n)

	primeTags := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primeTags[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if primeTags[i] {
			for j := i * i; j <= n; j += i {
				primeTags[j] = false
			}
		}
	}
	return primeTags
}

func Sieve2(n int) []bool {
	// time: O(n log log n), space: O(n)
	// ?

	if n < 2 {
		return make([]bool, n)
	}

	primeTags := make([]bool, n+1)
	primeTags[2] = true
	for i := 3; i <= n; i += 2 {
		primeTags[i] = true
	}
	for i := 3; i*i <= n; i += 2 {
		if primeTags[i] {
			for j := i * i; j <= n; j += i {
				primeTags[j] = false
			}
		}
	}
	return primeTags
}

func GetLeastPrimeFactors(maxNum int) []int {
	// time: O(n), space: O(n)
	// hasn't been tested!

	leastPrimeFactors := make([]int, maxNum+1)
	primes := []int{}
	for i := 2; i <= maxNum; i++ {
		if leastPrimeFactors[i] == 0 {
			leastPrimeFactors[i] = i
			primes = append(primes, i)
		}
		for j := 0; primes[j]*i <= maxNum && primes[j] <= leastPrimeFactors[i]; j++ {
			leastPrimeFactors[primes[j]*i] = primes[j]
		}
	}
	return leastPrimeFactors
}

func GetPrimeFactorsViaLPF(num int, leastPrimeFactors []int) []int {
	// time: O(log n), space: O(1) + ..
	// hasn't been tested!

	primes := []int{}
	for num > 1 {
		prime := leastPrimeFactors[num]
		if len(primes) == 0 || primes[len(primes)-1] != prime {
			primes = append(primes, prime)
		}
		num /= prime
	}
	return primes
}

func testLPF() {
	// todo
}

// greatest common divisor
func GCD(a, b int) int {
	// time: O(log min(a,b)), space: O(1)

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func GCD2(a, b int) int {
	// time: O(log min(a,b)), space: O(log min(a,b))

	if b == 0 {
		return a
	}
	return GCD2(b, a%b)
}

// least common multiple
func LCM(a, b int) int        { return Abs(a*b) / GCD(a, b) } // a / GCD(a, b) * b
func DivCeil(a, b int) int    { return (a + b - 1) / b }
func SumTo(n int) int         { return n * (n + 1) / 2 } // n*(n-1)/2 - sum to (n-1)
func IsPowerOfTwo(x int) bool { return x > 0 && x&(x-1) == 0 }

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Calculate(expression string) float64 { return 0 }

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

func Factorial(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

func MulMatrices(a, b [][]int, mod int) [][]int {
	ah, aw := len(a), len(a[0])
	bw := len(b[0])
	// handle bad input
	res := make([][]int, ah)
	for i := range res {
		res[i] = make([]int, bw)
		for j := range res[i] {
			sum := 0
			for p := range aw {
				sum += a[i][p] * b[p][j]
			}
			res[i][j] = sum % mod
		}
	}
	return res
}

func StrNumToDigits(num string, base int) []int {
	// time: O(n^2), space: O(1) + O(n)
	// n - len(num)

	if num == "0" {
		return []int{0}
	}

	digits := []int{}
	var digit int
	for num != "0" {
		num, digit = StrNumDivMod(num, base)
		digits = append(digits, digit)
	}
	for l, r := 0, len(digits)-1; l < r; l, r = l+1, r-1 {
		digits[l], digits[r] = digits[r], digits[l]
	}
	return digits
}

func StrNumDivMod(num string, div int) (quotient string, remainder int) {
	// time: O(n), space: O(1) + O(n)
	// n - len(num)

	digitCount := len(num)
	res := []byte{}
	n := 0
	for i := 0; i < digitCount; {
		n = n*10 + int(num[i]-'0')
		i++
		for i < digitCount && n < div {
			n = n*10 + int(num[i]-'0')
			if len(res) != 0 {
				res = append(res, '0')
			}
			i++
		}
		res = append(res, byte(n/div)+'0')
		n %= div
	}
	return string(res), n
}

func testDivMod() {
	const aLen = 500
	const maxB = 1000
	a, b, c, m := new(big.Int), new(big.Int), new(big.Int), new(big.Int)
	t := time.Now()
	for range 10000 {
		num := GenerateRandomLongNum(aLen)
		div := rand.IntN(maxB) + 1
		a.SetString(num, 10)
		b.SetInt64(int64(div))
		c.DivMod(a, b, m)
		res, rem := StrNumDivMod(num, div)
		if res != c.String() || rem != int(m.Int64()) {
			fmt.Printf("divMod failed: %v / %v = %v (%v), should be: %v (%v)\n",
				num, div, res, rem, c, m)
		}
	}
	fmt.Printf("testDivMod() time: %v\n", time.Since(t))
}

func GenerateRandomLongNum(length int) string {
	num := make([]byte, length)
	for i := range num {
		num[i] = byte(rand.IntN(10)) + '0'
	}
	return string(num)
}

func testSieve() {
	const maxNum int = 1e5
	const loopCount = 1000

	countPrimes := func(n int) int {
		count := 0
		for i := range n + 1 {
			if IsPrime(i) {
				count++
			}
		}
		return count
	}
	countPrimesWithSieve := func(n int) int {
		count := 0
		s1 := Sieve(n)
		s2 := Sieve2(n)
		if !slices.Equal(s1, s2) {
			fmt.Printf("Sieve & Sieve2 yield different results for n = %v\n", n)
		}
		for _, isPrime := range s2 {
			if isPrime {
				count++
			}
		}
		return count
	}

	t := time.Now()
	var d1, d2 time.Duration
	for range 1 { //loopCount {
		n := int(1e7) //rand.IntN(maxNum)
		_t := time.Now()
		count1 := countPrimes(n)
		d1 += time.Since(_t)
		_t = time.Now()
		count2 := countPrimesWithSieve(n)
		d2 += time.Since(_t)
		if count1 != count2 {
			fmt.Printf("Count primes failed for %v: %v != %v (sieve)\n", n, count1, count2)
		}
	}
	fmt.Printf("Count primes test time: %v (%v, %v (sieve))\n", time.Since(t), d1, d2)
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

func testPow() {
	pow := func(base, exp int) int {
		// if exp == 0 {
		// 	return 1
		// }
		res := 1
		for range exp {
			res *= base
		}
		return res
	}

	t := time.Now()
	for base := range 13 {
		for exp := range 17 {
			// res1 := Pow(base, exp)
			res1 := PowMod(base, exp, math.MaxInt)
			// res2 := int(math.Pow(float64(base), float64(exp)))
			res2 := pow(base, exp)
			if res1 != res2 {
				fmt.Printf("testPow unexpexted result: %v^%v = %v insted of %v\n", base, exp, res1, res2)
			}

			res1 = Pow(base, exp)
			res2 = Pow2(base, exp)
			if res1 != res2 {
				fmt.Printf("testPow unexpexted result 2: %v^%v = %v insted of %v\n", base, exp, res1, res2)
			}
		}
	}
	fmt.Printf("testPow() time: %v\n", time.Since(t))
}

func main() {
	testSieve()
	testPow()
	testDivMod()
	return

	// x := 3710
	// fmt.Printf("%b\n", x)
	// fmt.Println(CountBits(x))
	// fmt.Println(CountSetBits(x))
	// return

	n := 3
	for _, perm := range Permutations(n) {
		fmt.Println(perm)
	}
	fmt.Println()
	for _, perm := range Permutations0(n) {
		fmt.Println(perm)
	}
	// fmt.Println(len(Permutations(n)))
	// fmt.Println(len(Permutations0(n)))
	return

	// for i := range 10 {
	// 	fmt.Printf("%v: %v\n", i, Factorial(i))
	// }
	// return

	testCombinations()
	return

	// t := time.Now()
	// for i := range int(1e5) {
	// 	PrimeFactorization(i)
	// }
	// fmt.Println(time.Since(t))

	fmt.Println(GetPrimeFactors(2 * 3 * 5 * 7 * 11 * 13))
	fmt.Println(GetPrimeFactors(2))
	return

	// k := 0
	// for i := range int(1e6 + 1) {
	// 	if IsPrime(i) {
	// 		k++
	// 	}
	// }
	// fmt.Println(k)
}

// 2^6 = 2^(2*3) = (2^2)^3
// 4^3 = (4^2)
// 2^3 = 2^1 * 2^1 * 2^1
// 2^3 = 2^2 * 2^1
