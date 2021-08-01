package main

import(
		"fmt"
		// "errors"
		// "math"
		// "github.com/James-Muralles/Go-Practice/num_in_list"
		// "github.com/James-Muralles/Go-Practice/sumlist"
		"github.com/James-Muralles/Go-Practice/strings"
		



)

// a struct is a collection of fields that you can group together to make a more logical type. done outside of the main

// type person struct {
// 	name string
// 	age int
// }

func main () {


// p := person{name: "Jake", age: 23}
// fmt.Println(p) // gives you the struct (like an object)
// fmt.Println(p.name) // gives you field name (like properties )
// fmt.Println(p.age)// gives you the field age

// initializing a variable
	// x := 1
	// fmt.Println(x)
	// fmt.Println(&x) // gives you the memory adress of a variable


// if else statement
	// x := 1

	// if x > 6 {
	// 	 fmt.Println("More than 6")
	// } else if x < 2 {
	// 	fmt.Println("Less than 2")
	
	// } else {
	// 	fmt.Println("Number is in between 2 and 6")
	// }

	
// how to write arrays	
	// var a[4]int
	// a[3] = 5
	//removing the 6 from the [] makes it a slice of ints and allows append() reutrning a new array
		// a := []int{1,2,3,4,5,6}
		// // a[0] = 2
		// a = append(a, 15)
		// fmt.Println(a)

//maps key value pairs

// v := make(map[string]int)

// v["triangle"] = 2
// v["square"] = 3
// v["dodecagon"] = 12

// fmt.Println(v)
// fmt.Println(v["triangle"]) //this gives you the value of the key 
// delete(v, "square")//this deletes square
// fmt.Println(v)

// only type of loop in go is a for loop

// for i := 0; i < 5; i++ {
// 	fmt.Println(i)
// }

// for loop doubles as a while loop like this

// i := 0
// for i < 5 {
// 	fmt.Println(i)
// 	i ++
// }


// looping through an array

// arr := []string{"a", "b", "c"}

// arr2 := []int{10,20,30,40}


// for index, value := range arr{

// 	fmt.Println("index:", index, "value:", value)


// }

// looping through a map

// m := make(map[string]string) 

// m["a"] = "alpha"
// m["b"] = "beta"
// m["c"] = "charlie"

// for key, value := range m{

// 	fmt.Println("key:", key, "value:", value)


// }

// result := sum(4,5)
// fmt.Println(result)


//go has no exceptions it uses errors. if error is not equal to nil it will return the error, if else it returns the positive result

// result, err := sqrt(-16)

// if err != nil {
// 	fmt.Println(err)
// } else {
// 	fmt.Println(result)
// }

// i :=  7
// inc(&i) // points to the memory space at i
// fmt.Println(i)	

// }

// making a function outside of the main

// func sum(x int, y int ) int {
// return x + y
// }


// //go can return more than 1 type. In this case it returns a float64 if the number is greater than 0 or it return and error message if it is negative

// func sqrt(x float64) (float64, error) {
// if x < 0{

// 	//for an error message return 0  because there is no calc done, and retirn the error message
// 	return 0, errors.New("Undefined for negative numbers")
// }

// // needs two returns so it returns the calculation if greater than 0 and nil
// return math.Sqrt(x), nil


// }


// pointers
// func inc(x *int){
// 	*x++ //derefrencing 



// fmt.Println(num_in_list.NumInList([]int{1,2,3,4},4))
// fmt.Println(sumlist.SumList(arr2))
fmt.Println(strings.Reverse("hello world"))

}
