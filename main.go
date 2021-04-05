package main

import (
	"fmt"

	rest "goodgoods/api"
)

func main() {
	fmt.Printf("Start external adapter for Good Goods")

	// create API
	rest.Create()
}
