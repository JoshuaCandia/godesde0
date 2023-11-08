package variables

import "fmt"

func ShowIntegers (){
	var commonint int 
	intof32 := int32(10)
	intof64 := int64(100)
	
	fmt.Println("intcomun =",commonint)
	fmt.Println("intof32 =",intof32)
	fmt.Println("intof64 =",intof64)
}