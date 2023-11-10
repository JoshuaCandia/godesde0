package exercises

import "strconv"

func StrToInt(str string) (int, string) {
	integer, err := strconv.Atoi(str)
	if err != nil {
		return integer, "Has an error" + err.Error()
	}

	if integer > 100 {
		return integer, "Is higher than 100"
	} else {
		return integer, "is lower than 100"
	}
}
