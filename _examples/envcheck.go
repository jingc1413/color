package main

import (
	"fmt"
	"runtime"

	"github.com/gookit/color"
	// "github.com/gookit/goutil/dump"
)

// go run ./_examples/envcheck.go
func main() {
	fmt.Println("OS", runtime.GOOS)

	fmt.Println("TermColorLevel:", color.TermColorLevel())
	fmt.Println("IsSupportColor:", color.IsSupportColor())
	fmt.Println("IsSupport256Color:", color.IsSupport256Color())
	fmt.Println("IsSupportRGBColor:", color.IsSupportRGBColor())

	// dump.P(os.Environ())
	// fmt.Println(os.Environ())
}
