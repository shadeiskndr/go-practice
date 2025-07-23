package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"dog": "bark",
		"cat": "meow",
        "bird": "chirp",
	}
	changeMap(m)
	fmt.Println(m)

	for key, value := range m {
		fmt.Println(value, "is the sound of", key)
	}
}

func changeMap(m map[string]string) {
	m["duck"] = "quack"
}