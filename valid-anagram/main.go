package main

import "fmt"

func main() {
	fmt.Println(are_anagram("anagram", "nagaram"))
	fmt.Println(are_anagram("car", "rat"))
}

func are_anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	h1 := map[string]int{} // Hash map table 1
	h2 := map[string]int{} // Hash map table 2
	for _, s := range s1 {
		if _, ok := h1[string(s)]; ok {
			h1[string(s)] += 1
		} else {
			h1[string(s)] = 1
		}
	}
	for _, s := range s2 {
		if _, ok := h2[string(s)]; ok {
			h2[string(s)] += 1
		} else {
			h2[string(s)] = 1
		}

	}
	for k := range h1 {
		if _, ok := h2[string(k)]; !ok || h1[k] != h2[k] {
			return false
		}
	}
	return true
}
