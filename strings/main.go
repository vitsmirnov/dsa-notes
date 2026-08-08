package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"time"
)

// todo:
// Knuth-Morris-Pratt
func KMP(text, target string) []int    { return nil }
func TwoWay(text, target string) []int { return nil }

const mod int = 1e9 + 7
const mod2 int = 1e9 + 9
const primeBase = 257 // 31

func CustomHash(s string, base int, mod int) int {
	// time: O(n), space: O(1)

	hash := 0
	for i := range s {
		hash = (hash*base + int(s[i]) + 1) % mod
	}
	return hash
}
func CustomHash2(s string, base int, mod int, ord func(c byte) int) int {
	// time: O(n), space: O(1)

	hash := 0
	for i := range s {
		hash = (hash*base + ord(s[i])) % mod
	}
	return hash
}
func Hash(s string) int {
	// time: O(n), space: O(1)

	return CustomHash(s, primeBase, mod)
}

func IntToDecStr(n int) string { return IntToStr(n, 10, 0) }
func IntToBinStr(n int) string { return IntToStr(n, 2, 0) }
func IntToOctStr(n int, upper bool) string {
	if upper {
		return IntToStr(n, 8, 'A')
	}
	return IntToStr(n, 8, 'a')
}
func IntToHexStr(n int, upper bool) string {
	if upper {
		return IntToStr(n, 16, 'A')
	}
	return IntToStr(n, 16, 'a')
}
func IntToStr(n int, base int, startLetter byte) string {
	// time: O(log n), space: O(log n)

	if n == 0 {
		return "0"
	}

	res := []byte{}
	for ; n != 0; n /= base {
		digit := n % base
		if digit < 10 {
			res = append(res, byte(digit)+'0')
		} else {
			res = append(res, byte(digit-10)+startLetter)
		}
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}

const LowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
const UppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Digits = "0123456789"
const Symbols = " .,!?;:-_+=<>*&@$#%()[]{}/^|~`'" // " .,!?;:-_+=<>*&@$#%№()[]{}\\/^|~`'\""
const AllLetters = LowercaseLetters + UppercaseLetters
const LettersAndDigits = AllLetters + Digits
const AllChars = LettersAndDigits + Symbols

func GenerateRandomStringFromChars(length int, chars string) []byte {
	charCount := len(chars)
	s := make([]byte, length)
	for i := range s {
		s[i] = chars[rand.IntN(charCount)]
	}
	return s
}

func GenerateRandomString(length int) []byte {
	return GenerateRandomStringFromChars(length, AllChars)
}

func FindSubstrings(text, target string) []int {
	textLen, targetLen := len(text), len(target)
	if targetLen > textLen {
		return []int{}
	}

	indices := []int{}
	for i := range textLen - targetLen + 1 {
		if text[i:i+targetLen] == target {
			indices = append(indices, i)
		}
	}
	return indices
}

func AreStringsEqual(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func RabinKarp(text, target string) []int {
	textLen, targetLen := len(text), len(target)
	if targetLen > textLen {
		return []int{}
	}

	targetHash, textHash := 0, 0
	pow := 1
	for i := range target {
		targetHash = (targetHash*primeBase + int(target[i])) % mod
		textHash = (textHash*primeBase + int(text[i])) % mod
		pow = (pow * primeBase) % mod
	}
	indices := []int{}
	if targetHash == textHash && target == text[:targetLen] {
		indices = append(indices, 0)
	}
	for i := targetLen; i < textLen; i++ {
		textHash = (textHash*primeBase + int(text[i])) % mod
		textHash = (textHash - (int(text[i-targetLen])*pow)%mod + mod) % mod
		if targetHash == textHash && target == text[i-targetLen+1:i+1] {
			indices = append(indices, i-targetLen+1)
		}
	}
	return indices
}

// simple regular expression matching
// pattern may contain:
// '?' - matches any single character​​
// '*' - matches zero or more of character​​s
func isMatch(s string, pattern string) bool {
	// time: O(n^2 * m), space: O(n*m)
	// n - len(s)
	// m - len(pattern)

	const (
		unknown = iota
		match
		noMatch
	)
	p := make([]byte, 0, len(pattern))
	for i := range pattern {
		if pattern[i] != '*' || i == 0 || pattern[i-1] != '*' {
			p = append(p, pattern[i])
		}
	}
	sLen, pLen := len(s), len(p)
	memo := make([][]int, sLen)
	for i := range memo {
		memo[i] = make([]int, pLen)
	}

	var isMatch func(i, j int) bool // suffix matching
	isMatch = func(i, j int) bool {
		if i == sLen || j == pLen {
			return i == sLen && // s is done
				(j == pLen || (pLen-j == 1 && p[j] == '*')) // p is done or p == '*'
		}
		if memo[i][j] != unknown {
			return memo[i][j] == match
		}

		res := false
		if p[j] == '*' {
			for _i, _j := i, j+1; _i <= sLen && !res; _i++ {
				res = isMatch(_i, _j)
			}
		} else {
			res = (p[j] == '?' || s[i] == p[j]) && isMatch(i+1, j+1)
		}
		if res {
			memo[i][j] = match
		} else {
			memo[i][j] = noMatch
		}
		return res
	}

	return isMatch(0, 0)
}

func testFindSubstrings() {
	const textLen int = 1e8
	const targetLen int = 1000
	text := GenerateRandomStringFromChars(textLen, LowercaseLetters)
	target := GenerateRandomStringFromChars(targetLen, LowercaseLetters)
	for range 10000 {
		start := rand.IntN(textLen - targetLen + 1)
		for i := range targetLen {
			text[i+start] = target[i]
		}
	}
	t := time.Now()
	indices1 := FindSubstrings(string(text), string(target))
	d1 := time.Since(t)
	t = time.Now()
	indices2 := RabinKarp(string(text), string(target))
	d2 := time.Since(t)
	fmt.Printf("FindSubstrings: time: %v, occurrences found: %v\n", d1, len(indices1))
	fmt.Printf("RabinKarp: time: %v, occurrences found: %v\n", d2, len(indices2))
	fmt.Printf("indeces are equal: %v\n", slices.Equal(indices1, indices2))
}

func testIntToStr() {
	test := func(n int, base int, format string) {
		r1 := IntToStr(n, base, 'a')
		r2 := fmt.Sprintf(format, n)
		if r1 != r2 {
			fmt.Printf("IntToStr failed: (n = %v, base = %v) %q != %q\n", n, base, r1, r2)
		}
	}

	t := time.Now()
	for n := range int(1e5) {
		for _, baseFormat := range [][]any{{10, "%v"}, {2, "%b"}, {8, "%o"}, {16, "%x"}} {
			test(n, baseFormat[0].(int), baseFormat[1].(string))
		}

		// fmt.Printf("%d\n%v\n", n, IntToStr(n, 10))
		// fmt.Printf("%b\n%v\n", n, IntToStr(n, 2))
		// fmt.Printf("%o\n%v\n", n, IntToStr(n, 8))
		// fmt.Printf("%x\n%v\n", n, IntToStr(n, 16))
	}
	fmt.Printf("testIntToStr time: %v\n", time.Since(t))
}

func main() {
	// for c := range 256 {
	// 	fmt.Printf("%c", c)
	// }
	// fmt.Println()
	// fmt.Printf("%q\n", GenerateRandomStringFromChars(50, LowercaseLetters))
	// fmt.Printf("%q\n", GenerateRandomStringFromChars(50, UppercaseLetters))
	// fmt.Printf("%q\n", GenerateRandomStringFromChars(50, AllLetters))
	// fmt.Printf("%q\n", GenerateRandomStringFromChars(50, Digits))
	// fmt.Printf("%q\n", GenerateRandomStringFromChars(50, LettersAndDigits))
	// fmt.Printf("%q\n", GenerateRandomStringFromChars(50, Symbols))
	// fmt.Printf("%q\n", GenerateRandomStringFromChars(50, AllChars))
	// fmt.Printf("%q\n", GenerateRandomStringFromChars(50, LowercaseLetters+" .,"))
	// fmt.Printf("%q\n", GenerateRandomStringFromChars(50, Digits+"+-*/=()"))

	testFindSubstrings()
	return

	testIntToStr()
}
