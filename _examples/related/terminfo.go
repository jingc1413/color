package main

import (
	"fmt"

	"github.com/xo/terminfo"
)

// go run ./related/terminfo.go
func main() {
	// ti, err := terminfo.LoadFromEnv()
	// if err != nil {
	// 	panic(err)
	// }

	lv, err := terminfo.ColorLevelFromEnv()
	if err != nil {
		panic(err)
	}

	fmt.Println(lv.String())
}
