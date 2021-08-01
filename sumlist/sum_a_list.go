package sumlist

import(



)


func sumList (list []int) int{
	var sumTotal int 
 	for _, val := range list {
		sumTotal += val
	}
	return sumTotal

}