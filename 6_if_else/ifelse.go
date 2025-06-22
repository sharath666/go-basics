package main 

import "fmt"

func main(){
	age :=18

	// if age >=18 {
	// 	fmt.Println("person is an adult")
	// } else {
	// 	fmt.Println("person is not an adult")
	// }


	//implementing elseif

	age =16
	if age >=18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is teeanger")
	} else{
		fmt.Println("person is a kid")
	}


	//multi condition 

	var role = "admin"
	var hasPermissions = true

	if role == "admin" || hasPermissions{
		fmt.Println("yes")
	}

	//we can declare the variable inside the if condition

	if age:= 20; age >=18{
		fmt.Println("person is an adult", age)
	} else if age>=12 {
		fmt.Println("person is teenager", age)
	}

}
