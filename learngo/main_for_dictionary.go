package main

import (
	"fmt"

	"github.com/hoondal6970/learngo/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greetings"
	dictionary.Add(word, definition)
	dictionary.Update(word, "modified")
	dictionary.Delete(word)
	found, err := dictionary.Search(word)
	fmt.Println(found, err)

}
