package main

import (
	"fmt"
	"sort"
)

func main() {
	fruits := map[string]int{
		"apple": 1,

		"orange": 2,
		"banana": 3,
	}

	keys := make([]string, 0, len(fruits))
	for key := range fruits {
		keys = append(keys, key)
	}
	fmt.Printf("keys: %v\n", keys)
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("Key: %s, Value: %d\n", key, fruits[key])

	}
}
