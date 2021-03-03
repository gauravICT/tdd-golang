package main

import "fmt"

const helloPrefix = "Hello, "

// Hello :
func Hello(name string) string {
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("Gaurav"))
}
