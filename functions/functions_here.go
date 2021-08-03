package functions

import (
"fmt"

)

func FizzBuzz (input int) string {

	if input == 100{

		return "\n"
	}
	if input % 15 == 0 {
		return "\nFizzBuzz\n"
	}
	if input % 3 == 0 {
		return "\nFizz\n"
	}
	if input % 5 == 0 {
		return "\nBuzz\n"
	}
	
	return fmt.Sprintf("\n%d\n" ,input)
}