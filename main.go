package main

import (
	"fmt"

	rest "goodgoods/api"
)

func main() {
	fmt.Printf("Start external adapter for Good Goods\n\n")

	// create API
	rest.Create()
}
