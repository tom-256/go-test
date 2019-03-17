package main

import (
	"fmt"
)

const englishPrefix = "Hello, "
const spanish = "Spanish"
const french = "French"
const spanishHelloPrefix = "Holla, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name ,lang string) string {
	if name =="" {
		name = "world"
	}

	return greetingPrefix(lang) + name

}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishPrefix
		}
	return
}


func main(){
	fmt.Println(Hello("world","English"))
}
