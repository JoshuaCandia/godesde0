package main

import (
	"fmt"

	"github.com/JoshuaCandia/godesde0/variables"
)

func main(){
	state,text := variables.ConvertToText(1500)
	fmt.Println(state)
	fmt.Println(text)
}
