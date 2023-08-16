package main // A standalone program (as opposed to a library) is always in package main. If this code were a library,
// the package would be hello instead

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

// go mod init example/hello
// go run . to run this program
