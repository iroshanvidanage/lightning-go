package main

import (
	"fmt"
	"os"
)

var name string = "I'm Iroshan"

func main() {
	fmt.Println("Hello" + name)
	_, exists := os.LookupEnv("MY_NAME")
	if !exists {
		fmt.Println("MY_NAME var is not set")
		fmt.Println("Setting MY_NAME var....")
		setEnvVar()
	}

	val := os.Getenv("MY_NAME")

	fmt.Println("MY_NAME", val)
}

func setEnvVar() {
	os.Setenv("MY_NAME", "Iroshan")
}
