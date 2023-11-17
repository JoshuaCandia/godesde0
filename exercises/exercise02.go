package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num int

func Multiply() {
	for {
		fmt.Println("Entry number: ")
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			num, err := strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Invalid number, try again")
			} else {
				for i := 0; i <= 10; i++ {
					fmt.Println(num, "x", i, "=", num*i)
				}
				break
			}
		}
	}
}
