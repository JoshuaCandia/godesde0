package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Multiply() string {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for {
		fmt.Println("Enter a number:")
		if scanner.Scan() {
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}

			for i := 1; i <= 10; i++ {
				text += fmt.Sprintf("%d x %d = %d\n", num, i, num*i)
			}
			break
		} else {
			fmt.Println("Error reading input. Exiting.")
			break
		}
	}
	return text
}
