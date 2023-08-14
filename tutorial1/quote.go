package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println(quote.Opt())
}

//  go mod tidy to find and install required module
