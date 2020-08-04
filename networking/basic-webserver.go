package main

import ("fmt"
		"net/http")

func main() {

	// register handlers
	http.HandleFunc("/", handleFunc1)

	//  
	fmt.Println("running server...")
	http.ListenAndServe(":8080", nil)

}

// override handle for a pattern
func handleFunc1(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Ok you are somehow seeing this")
} 