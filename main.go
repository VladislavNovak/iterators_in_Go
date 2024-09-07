package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[int]string{
		1: "first",
		2: "second",
		3: "third",
		4: "fourth",
	}

	for k, v := range maps.All(m) {
		fmt.Printf("%d: %s\n", k, v)
	}
}
