package num_in_list

import (
	"fmt"
	"strconv"
	"strings"
)





func NumInList(list []int, num int) bool{
	for i, number := range list {
		fmt.Printf("\n This is my list at i , %d\n This is the number input: %d\n", number, num)
		if list[i] == num {
			fmt.Printf("\nThis number: %d is in this list: %s\n", num, SplitToString(list, " ,"))
			return true
		}
i++
	}
return false

}

func SplitToString(a []int, sep string) string {
    if len(a) == 0 {
        return ""
    }

    b := make([]string, len(a))
    for i, v := range a {
        b[i] = strconv.Itoa(v)
    }
    return strings.Join(b, sep)
}

