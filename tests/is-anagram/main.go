// https://neetcode.io/problems/is-anagram/

package main

func count(s string) (res [26]uint) {
	for _, c := range s {
		res[c-'a']++
	}

	return
}

func isAnagram(s string, t string) bool {
	return count(s) == count(t)
}

func main() {
	println(isAnagram("racecar", "carrace"))
	println(isAnagram("jar", "jam"))
	println(isAnagram("x", "x"))
	println(isAnagram("", ""))
}
