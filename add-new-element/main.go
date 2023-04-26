package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	//your code here!

	fmt.Println(data, newData, position)

	switch position {
	case "up":
		data = append(data, newData)
	case "down":
		data = append([]int{newData}, data...)
	} 

	return data
}

func main() {
	fmt.Println(AddElement([]int{1, 2, 3, 4, 5}, 6, "up"))
	fmt.Println(AddElement([]int{1, 2, 3, 4, 5}, 6, "down"))
}
