package functions

import (
	"fmt"
	"strconv"
	
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

func DecToBase (dec int, base int ) string {
	var res string 
	
	for  dec != 0 {
		remainder := dec % base
		res = strconv.Itoa(remainder) + res
		dec = dec / base

	}

	return res
	

}