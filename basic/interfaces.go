package main

import "fmt"

func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func giveType(v interface{}) {
  fmt.Printf("%T\n", v) ; 
}


func main() {

	// all types
	fmt.Println("interface{} matches with ...")

	giveType(1)
	giveType(3.0)
	giveType("ok")
	giveType('c')
	giveType(true)

	// needs conversion to use
	names := []string{"stanley", "david", "oscar"}
    vals := make([]interface{}, len(names))
    for i, v := range names {
        vals[i] = v
    }
    PrintAll(vals)

}