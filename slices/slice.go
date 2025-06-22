package main
import "fmt"
import "slices"
// slice dynamic arrays
// we should not give the length
// most used construct in go
//+ useful methods

func main() {
	// uninitializead slice is nil
	var nums []int
	fmt.Println(nums)

	//printing the length of the slice
	fmt.Println(len(nums))


	var numbers = make([]int, 2, 5)
	 // capacity -> maximum numbers of elements can fit
	 //if we know how many elements needed in the slice then you can send the capacity
	fmt.Println(cap(numbers))
	fmt.Println(numbers == nil)

	numbers = append(numbers,3)
	numbers = append(numbers,4)
	numbers = append(numbers,5)
	numbers = append(numbers,6)
	fmt.Println(cap(numbers))
	fmt.Println(numbers)

	// copy function

	var numbers3 = make([]int, len(numbers))

	fmt.Println(numbers, numbers3)
	copy(numbers3, numbers)
	fmt.Println(numbers3)

	//slice operator

	var number3 = []int{1,2,3}
	fmt.Println(number3[0:2]) // op [1,2]
	fmt.Println(number3[0:1]) // op [1]
	fmt.Println(number3[:1])  // [0]
	fmt.Println(number3[1:])

	//equality check of slices

	//using slices package you can call some of the methods which can be called

	var nums1 = []int{1,2}
	var nums2 = []int{1,2}

	fmt.Println(slices.Equal(nums1, nums2))

	var numbers4 = [][]int{{1,2,3}, {4,5,6}}
	fmt.Println(numbers4)
}