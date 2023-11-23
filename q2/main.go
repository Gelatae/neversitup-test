package main

import "fmt"

func main() {
	str := "abbbba"
	perms := GetPermutation(str)
	fmt.Println(len(perms))
	fmt.Println(perms)
}

func GetPermutation(str string) []string {
	// base case
	if len(str) == 0 {
		return []string{""}
	}
	// create empty array to store the answer
	res := []string{}
	// create map for exist string
	exist := make(map[string]bool)

	// find permutation of each character
	for i, s := range str {
		// exclude current character
		optString := str[:i] + str[i+1:]
		// get permutation of remaining string
		permutations := GetPermutation(optString)
		// concat all permutation to the current character
		for _, perm := range permutations {
			newStr := string(s) + perm
			if _, ok := exist[newStr]; !ok { // if string exist then skip adding
				exist[newStr] = true
				res = append(res, newStr)
			}
		}
	}
	return res
}
