package strings

import (
	"strings"
	"fmt"

)

// to reverse a character that is more than 1 byte use runes
//  ---> for _, r range word {
// 			res = string(r) + res
// 		}

func StringReverse (word string) string {
	var res string

	for i := len(word)-1 ; i>=0; i--{
		res += string(word[i])
	}

	return res


}

func ReverseString (word string) string {
	var sb strings.Builder
	for i := len(word)-1 ; i>=0; i--{
		sb.WriteByte(word[i])
	}

	return sb.String()


}

func Reverse (word string) string {
	var res string

	for i := 0;  i < len(word); i++{
		res = string(word[i]) + res
		fmt.Println(res)

	}

	return res


}
