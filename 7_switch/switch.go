package main
import "fmt"
import "time"
func main() {
	//simple switch

	i:=5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("Other")
	}


	//multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its, weekend")
	default:
		fmt.Println("It's workday")
	}

	//type switch

	whoAmI := func(i interface{}){
		switch i.(type){
		case int: 
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Println("other")
	}
}
whoAmI(56)
}

