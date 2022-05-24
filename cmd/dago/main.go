package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	buildInfo, _ := debug.ReadBuildInfo()
	fmt.Println(buildInfo)
}
