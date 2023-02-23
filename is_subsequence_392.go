package main

import "fmt"

func main() {
	result := isSubsequence("code", "leetcode")
	fmt.Println("result: ", result)
}

//Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
//A subsequence of a string is a new string that is formed from the original string by deleting
//some (can be none) of the characters without disturbing the relative positions of the remaining characters.
// (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
func isSubsequence(s string, t string) bool {

	tLength := len(t)
	subequence := 0

	for i := 0; i < tLength; i++ {
		if s[subequence] == t[i] {
			subequence++
		}
	}

	return subequence == len(s)
}
