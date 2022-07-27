package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to userinput"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rating for the program")

	// comma ok or err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for reading", input)
	fmt.Printf("type of this is %T", input)
}
