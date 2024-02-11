package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Initializing code health checker...")
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Current directory: ", wd)
	os.Exit(0)
}
