package main

import (
	"fmt"

	"github.com/hoondal6970/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary["hello"] = "Hello"
	fmt.Println(dictionary)
}
