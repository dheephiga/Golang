package main

import "fmt"

const Logintoken string = "WHJOWROVN"

func main() {
	var username string = "Dheephiga"
	fmt.Println(username)
	fmt.Printf("Variable is of type is: %T \n", username)

	var isloggedon bool = true
	fmt.Println(isloggedon)
	fmt.Printf("Variable is of type is: %T \n", isloggedon)

	var smallval int = 256
	fmt.Println(smallval)
	fmt.Printf("Variable is of type is: %T \n", smallval)

	var smallfloat float64 = 256.4565656466546465633
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type is: %T \n", smallfloat)

	var anothervar int
	fmt.Println(anothervar)
	fmt.Printf("variable is of type: %T", anothervar)

	var anotherstr string
	fmt.Println(anotherstr)
	fmt.Printf("variable is of type: %T", anotherstr)

	//explicit
	var website = "www.eheehahaehe.com"
	fmt.Println(website)
	fmt.Printf("variable type is %T\n", website)

	//no var style

	web := 3000000000
	fmt.Println(web)

	//public thing
	fmt.Println(Logintoken)

}
