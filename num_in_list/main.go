package main
import(
"fmt"

)

func main () {
fmt.Println(NumInList([]int{1,2,3,4},3))

}



func NumInList(list []int, num int) bool{
	for i, number := range list {
		fmt.Printf("\n This is my list at i , %d\n This is the number input: %d\n", number, num)
		if list[i] == num {
			
			return true
		}
i++
	}
return false

}