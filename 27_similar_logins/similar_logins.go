package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var in *bufio.Reader
var out *bufio.Writer

func main() {
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	processTest()
}

func processTest() {
	var usersCount int
	fmt.Fscan(in, &usersCount)
	users := make([]string, usersCount)
	for i := 0; i < usersCount; i++ {
		fmt.Fscan(in, &users[i])
	}

	var newUsersCount int
	fmt.Fscan(in, &newUsersCount)
	newUsers := make([]string, newUsersCount)
	for i := 0; i < newUsersCount; i++ {
		fmt.Fscan(in, &newUsers[i])
	}

	loginMap := mapUsersByHash(users)

	for _, newUser := range newUsers {
		hashed := hash(newUser)
		fmt.Fprintln(out, isAvailable(newUser, loginMap[hashed]))
	}
}

func hash(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func mapUsersByHash(users []string) map[string][]string {
	loginMap := make(map[string][]string)
	for _, user := range users {
		hashed := hash(user)
		loginMap[hashed] = append(loginMap[hashed], user)
	}
	return loginMap
}

func isAvailable(newUser string, existingUsers []string) int {
	for _, user := range existingUsers {
		if isSimilar(user, newUser) {
			return 1
		}
	}
	return 0
}

func isSimilar(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	diffIndex1, diffIndex2 := -1, -1
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if diffIndex1 == -1 {
				diffIndex1 = i
			} else if diffIndex2 == -1 {
				diffIndex2 = i
			} else {
				return false
			}
		}
	}
	if diffIndex1 == -1 {
		return true
	}
	if diffIndex2 == -1 {
		return false
	}
	return diffIndex1+1 == diffIndex2 && s1[diffIndex1] == s2[diffIndex2] && s1[diffIndex2] == s2[diffIndex1]
}
