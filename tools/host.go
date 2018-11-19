package main

import (
	"fmt"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(hostname)
}
