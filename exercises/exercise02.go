package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiply() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a number:")
		if scanner.Scan() {
			number, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}

			for i := 1; i <= 10; i++ {
				fmt.Printf("%d x %d = %d\n", number, i, number*i)
			}
			break
		} else {
			fmt.Println("Error reading input. Exiting.")
			break
		}
	}
}
