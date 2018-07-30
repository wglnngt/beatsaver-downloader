package main

import (
	"fmt"
)

func main() {
	response := fetch(0)
	fmt.Printf("%v", len(response.Songs))
}
