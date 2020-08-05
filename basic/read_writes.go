package main

import ("fmt"
		"os")

func main() {

	var num int 

	// makes it while true.
	for true {

		// read element
		fmt.Printf("Give me a number ")
		fmt.Scanf("%d", &num)		

		if num == 1729 {
			fmt.Println("Woah")
			break
		}
		fmt.Println("give me something better\nmore like 1729..")

	}

	f, err := os.Create("go-test-file")

	if(err != nil){
		panic(err)
	}

	defer f.Close()

	n3, err := f.WriteString("Hello world!!\n")
    fmt.Printf("wrote %d bytes\n", n3)

    // read from file

    
    // set offset 
    o, err := f.Seek(-8, 2)
    fmt.Printf("new offset %d\n", o)

    buffer := make([]byte, 8)
    f.Read(buffer)
    // "expects last 6 characters"
    fmt.Printf("read %s", string(buffer))
}