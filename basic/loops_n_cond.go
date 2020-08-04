package main 
import "fmt"

func main() {

	n := 100
	for i := 0 ; i < n ; i++ {
		if i % 10 == 0 {
			fmt.Println(i)			
		}

	} 
}