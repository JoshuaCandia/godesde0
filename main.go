package main

import (
	"fmt"

	"github.com/JoshuaCandia/godesde0/exercises"
)

func main() {
	//variables
	/* 	state, text := variables.ConvertToText(1500)
	   	fmt.Println(state)
	   	fmt.Println(text) */

	//If
	/* if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("This is not windows")
	} else {
		fmt.Println("This is windows")
	} */

	//Switch
	/* switch os := runtime.GOOS; os {

	case "linux":
		fmt.Println("This is linux")

	case "darwin":
		fmt.Println("This is darwin")
	default: //%s == string
		fmt.Printf("%s \n", os)
	} */
	num, str := exercises.StrToInt("90")

	fmt.Println(num, str)
	numb, stri := exercises.StrToInt("110")
	fmt.Println(numb, stri)
}
