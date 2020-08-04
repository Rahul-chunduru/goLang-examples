package main

import "fmt"

func main() {

	zeta := map[string] map[string] string {
		"hmm" : map[string] string {
		    "0" : "ok" },
		"zen" : map[string] string {
			"1" : "peace" },
	}

	fmt.Println(zeta)
}