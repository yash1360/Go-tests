package main

import "fmt"

const assignNameParameter = "Hello, "

func Hello(name string) string{
	if name == "" {
		name = "User"
	}

	return assignNameParameter + name
}

func main (){
	fmt.Println(Hello("user"))
}