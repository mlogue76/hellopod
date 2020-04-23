package main

import (
	"fmt"
	"os"
)

func main() {
	n, _ := os.Hostname()
	fmt.Printf("Hello Pod `%s`\n", n)
}
