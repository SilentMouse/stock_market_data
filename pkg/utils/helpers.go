package utils

import (
	"fmt"
	"crypto/rand"
)

func RemoveDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func RemoveFindElement(elements []string, element string) []string {

	result := []string{}

	for _,v := range elements {
		if v != element {
			result = append(result, v)
		}
	}

	return result
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func RandToken() string {
	b := make([]byte, 30)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}


func RandToken5() string {
	b := make([]byte, 5)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}