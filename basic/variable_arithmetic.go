package main

import "fmt"

func main() {

	var a int = 3 ; 
	b := 5 ; 
	var (
		c = 6 ; 
		d = 7 ; 
	)

	// printing works !
	fmt.Println("I see", a, a + b, a + b * c, a + b * c / d)
	// gives type.
	fmt.Printf("%T\n", d)

}