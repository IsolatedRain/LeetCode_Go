package main

import (
	"fmt"
	"sort"
)

// type set struct {
// 	vals map[string]struct{}
// }

// func newSet() set {
// 	v := map[string]struct{}{}
// 	return set{v}
// }

func accountsMerge(accounts [][]string) [][]string {
	emailName := map[string]string{}
	p := map[string]string{}
	var getRoot func(x string) string
	getRoot = func(x string) string {
		if _, ok := p[x]; !ok {
			p[x] = x
		}
		if p[x] != p[p[x]] {
			p[x] = getRoot(p[x])
		}
		return p[x]
	}
	for _, acc := range accounts {
		name := acc[0]
		xRoot := getRoot(acc[1])
		emailName[xRoot] = name
		for _, email := range acc[2:] {
			yRoot := getRoot(email)
			if xRoot != yRoot {
				p[yRoot] = xRoot
			}
		}
	}
	emails := map[string][]string{}
	exist := map[string]bool{}
	for email, root := range p {
		k := getRoot(root)
		if _, ok := emails[k]; !ok {
			emails[k] = append(emails[k], k)
			exist[k] = true
		}
		if !exist[email] {
			emails[k] = append(emails[k], email)
			exist[email] = true
		}
	}
	res := [][]string{}
	for k, v := range emails {
		name := emailName[k]
		curEmails := []string{}
		for _, email := range v {
			curEmails = append(curEmails, email)
		}
		sort.Strings(curEmails)
		curEmails = append([]string{name}, curEmails...)
		res = append(res, curEmails)
	}
	return res
}

func main() {
	accounts := [][]string{{"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}}
	fmt.Println(accountsMerge(accounts))
}
