package main

import (
	"fmt"
	"os"
)

func main() {
	p := NewInitialModel()
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
