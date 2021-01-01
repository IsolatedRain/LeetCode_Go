package main

import (
	"fmt"
	"strconv"
	"strings"
)

func trulyMostPopular(names []string, synonyms []string) []string {
	nameFre := map[string]int{}
	p := map[string]string{}

	getRoot := func(name string) string {
		if _, ok := p[name]; !ok {
			return ""
		}
		for p[name] != name {
			name = p[name]
		}
		return name
	}

	union := func(name1, name2 string) {
		name1Root, name2Root := getRoot(name1), getRoot(name2)
		if name1Root != "" && name2Root != "" && name1Root != name2Root {
			if name1Root < name2Root {
				p[name2Root] = name1Root
				nameFre[name1Root] += nameFre[name2Root]
				nameFre[name2Root] = 0
			} else {
				p[name1Root] = name2Root
				nameFre[name2Root] += nameFre[name1Root]
				nameFre[name1Root] = 0
			}
		}
	}

	for _, s := range names {
		i := strings.Index(s, "(")
		name := s[:i]
		p[name] = name
		nameFre[name], _ = strconv.Atoi(s[i+1 : len(s)-1])
	}

	for _, syn := range synonyms {
		i := strings.Index(syn, ",")
		name1 := syn[1:i]
		name2 := syn[i+1 : len(syn)-1]
		union(name1, name2)
	}

	res := []string{}
	for name := range nameFre {
		if nameFre[name] != 0 {
			res = append(res, name+"("+strconv.Itoa(nameFre[name])+")")
		}
	}
	return res
}

func main() {
	names := []string{"John(15)", "Jon(12)", "Chris(13)", "Kris(4)", "Christopher(19)"}
	synonyms := []string{"(Jon,John)", "(John,Johnny)", "(Chris,Kris)", "(Chris,Christopher)"}
	r := trulyMostPopular(names, synonyms)
	fmt.Println(r)
}
