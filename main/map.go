package main

import "fmt"

func test() {
	m := map[string]int{"apple": 20, "age": 19}
	frults := map[string]int{
		"apple": 1,
		"banna": 2,

		"orange": 3,
	}
	fmt.Printf("frults: %v\n", frults)
	m["banne"] = 3
	delete(m, "apple")
	key, ok := m["apple"]
	fmt.Printf("key: %v\n", key)
	fmt.Printf("ok: %v\n", ok)

	for key, v := range m {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("v: %v\n", v)

	}
}
func main() {
	test()

}
