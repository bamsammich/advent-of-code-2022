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
