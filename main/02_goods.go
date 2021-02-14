package main

import (
	"fmt"
	"runtime"
	"os"
)

const (
	//自动推断
	//a := 50
	//b := false
	//d,e,f = 5,7,"abc"
)

func main()  {

	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s", goos)
	path := os.Getenv("PATH")
	fmt.Printf("path is %s\n", path)
}

