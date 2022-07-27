package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to math in GO :p")

	//var mynumone int = 2
	//var mynumtwo float64 = 4

	//not valid
	//fmt.Println("The sum is: ", mynumone+int(mynumtwo))

	//random number
	//	rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(5) + 1)

	//using crypto
	myrannum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myrannum)

}
