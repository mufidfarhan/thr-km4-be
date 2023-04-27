package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {

	newArr := []any{}
	if position == "left" {
		newArr = arr[1:]
	} else if position == "right" {
		newArr = arr[:len(arr)-1]
	}

	return newArr
}

func main() {
	fmt.Println(removeLeftRight([]any{1, 2, 3, 4, 5}, "left"))
	fmt.Println(removeLeftRight([]any{1, 2, 3, 4, 5}, "right"))
	fmt.Println(removeLeftRight([]any{"Edo", "Budi", "Joko", "Tono"}, "left"))
	fmt.Println(removeLeftRight([]any{"Edo", "Budi", "Joko", "Tono"}, "right"))
}
