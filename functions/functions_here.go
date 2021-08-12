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

func FindTwoThatSum (list[] int, sum int) (int,int){
	for i, val := range list {
		for j, val2 := range list {
			if i == j  {
				continue
			}
			if val + val2 == sum{
				return i,j
			}
		}

	}
	return -1,-1

}
func Fibionacci(n int) int {
	if n <= 1 {
		return n
	} else {

		

		return Fibionacci(n - 1) + Fibionacci(n - 2);
		
	}
	
}

func  FibNoRecursion(n int) int {
	
	if n <= 1 {
		return n
	}
	nMinus2 := 0
	nMinus1 := 1
	var cur int
	for i := 2; i <=n; i++ {
		cur = nMinus2 + nMinus1
		nMinus2 = nMinus1
		nMinus1 =cur
	}
	return cur
}