package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var leyend string
var err error

func EntryNumbers() {
	fmt.Println("Entry number 1: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("The input is incorrect" + err.Error())
		}
	}

	fmt.Println("Entry number 2: ")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("The input is incorrect" + err.Error())
		}
	}
	fmt.Println("Entry leyend: ")
	if scanner.Scan() {
		leyend = scanner.Text()
		if err != nil {
			panic("The input is incorrect")
		}
	}
	fmt.Println(leyend, num1*num2)
}
