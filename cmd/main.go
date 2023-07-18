package main

import (
	"blockchain/utility"
	"fmt"
)

func main() {
	fmt.Println(utility.GetHost())
	// fmt.Println(utility.FindNeighbors("127.0.0.1", 5001, 0, 3, 5000, 5003))
}
