package main

import "fmt"

func main() {
	fmt.Println("Apache's Counter Utility is Setting Up")
	util := NewUtil()

	util.start()
}
