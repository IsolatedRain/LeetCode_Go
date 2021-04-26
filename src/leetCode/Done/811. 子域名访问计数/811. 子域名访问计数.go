package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	res := []string{}
	d := map[string]int{}
	for _, s := range cpdomains {
		words := strings.Split(s, " ")
		t, _ := strconv.Atoi(words[0])
		d[words[1]] += t
		for _, k := range dots(words[1]) {
			d[words[1][k+1:]] += t
		}
	}
	for k, v := range d {
		// res = append(res, strconv.Itoa(d[k])+" "+k)
		res = append(res, fmt.Sprintf("%d %s", v, k))
	}
	return res
}

func dots(s string) []int {
	res := []int{}
	for i := range s {
		if s[i] == '.' {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
