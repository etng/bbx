package main

import "fmt"

var commands = []string{
	"password",
	"passwd",
	"snippet",
	"help",
}

func main() {
	fmt.Println("Welcome to use BBX")
	fmt.Println("Available commands are:")
	for _, command := range commands {
		fmt.Printf(" * %s", command)
	}
}
