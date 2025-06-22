package main
import "fmt"
func main(){
	var nums[4] int
	nums[0] = 1
	fmt.Println(nums[0])
	//array length
	fmt.Println(len(nums))

	//since the array is of type int all the unassigned values will be zero
	fmt.Println(nums)
	//since the array is of type bool all the unassigned values will be false
	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)
	//since tohe array is of type string all the unassigned values will be empty string
	var name [3]string
	name[0] = "golang"
	fmt.Println(name)


	numbers := [3]int{1,2,3} // len datatype and elements 
	fmt.Println(numbers)

	//how to define 2 dimensional array

	numbers2 := [2][2]int{{3,4}, {5,6}}

	fmt.Println(numbers2)

	//fixed size, that is predictable
	// memory optimization
	// constant time access
}