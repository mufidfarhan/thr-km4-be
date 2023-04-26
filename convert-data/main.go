package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	var user User

	tokens := strings.Split(data, ",")

	user.Name = tokens[0]
	user.Age, _ = strconv.Atoi(tokens[1])
	user.Address = tokens[2]

	return user
}

func main() {
	fmt.Printf("%+v\n", ConvertData("Edo,20,Jakarta"))
	fmt.Printf("%+v\n", ConvertData("Budi,30,Bandung"))
}
