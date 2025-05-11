package main

import (
	"fmt"
	"time"

	"github.com/kumarvmsathish/mystrings"
	tinytime "github.com/wagslane/go-tinytime"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Printf("Reverse string: %v\n", mystrings.Reverse("Hello World!"))

	fmt.Println("-------------------------------------------------")
	fmt.Println("Import usign <go get module> and use")

	tt := tinytime.New(1746961439)
	tt = tt.Add(time.Hour * 48)
	fmt.Println(tt)
	fmt.Println("-------------------------------------------------")
}
