package sumlist



func SumList (list []int) int {
	var sumTotal int 
 	for _, val := range list {
		sumTotal += val
	}
	return sumTotal

}

// using recursion

func SumRecursion (list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + SumRecursion(list[1:])

}

// exmaple: SumRecursion(1,2,3,4) will execute like this,
// 1 + SumRecursion(2,3,4) => 2 SumRecursion(3,4) => 3 + SumRecursion(4) => 4 + SumRecursion()
