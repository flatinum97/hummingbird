package hummingbird

import (
	"fmt"
)

const version = "0.1.0"

// PrintVersion prints hummingbird version
func PrintVersion() {
	fmt.Printf("hummingbird %s\n", version)
}
