package iterators

import (
	"fmt"
)

func Iterator() {
	for i := 10; i >= 0; i-- {
		if i == 3 {
			break
		}
		if i == 8 {
			continue
		}
		fmt.Println(i)
	}
}
