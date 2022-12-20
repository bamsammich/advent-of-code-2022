package main

import (
	"encoding/json"
)

func PrettyJSON(in any) string {
	b, _ := json.MarshalIndent(&in, "", "  ")
	return string(b)
}

func sumInts(ints []int) int {
	var sum = 0
	for _, i := range ints {
		sum += i
	}
	return sum
}

func product[V int | float64](vals []V) V {
	var product V = 1
	for _, i := range vals {
		product *= i
	}
	return product
}

func index[T comparable](sl []T, e T) int {
	for i := range sl {
		if sl[i] == e {
			return i
		}
	}
	return -1
}

func removeAll[T comparable](sl []T, e T) []T {
	for i, c := range sl {
		if c == e {
			sl = append(sl[:i], sl[i+1:]...)
		}
	}
	return sl
}

func unique[T comparable](sl []T) []T {
	uniq := make(map[T]bool)
	for _, e := range sl {
		uniq[e] = true
	}
	var out []T
	for k := range uniq {
		out = append(out, k)
	}
	return out
}

func isUnique[T comparable](sl []T) bool {
	uniq := make(map[T]bool)
	for _, e := range sl {
		if ok := uniq[e]; ok {
			return false
		}
		uniq[e] = true
	}
	return true
}
