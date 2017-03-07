package main

import (
	"fmt"
	"os"

	"./unicode"
)

func main() {

	expression := "nord og sør"
	language := os.Args[1]
	//language := "is"
	trans := unicode.Translate(expression, language)
	fmt.Printf("%s", trans)
}
