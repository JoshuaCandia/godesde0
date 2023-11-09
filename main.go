package main

import (
	"fmt"
	"runtime"
)

func main() {
	/* 	state, text := variables.ConvertToText(1500)
	   	fmt.Println(state)
	   	fmt.Println(text) */
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("This is not windows")
	} else {
		fmt.Println("This is windows")
	}

	switch os := runtime.GOOS; os {

	case "linux":
		fmt.Println("This is linux")

	case "darwin":
		fmt.Println("This is darwin")
	default: //%s == string
		fmt.Printf("%s \n", os)
	}
}
