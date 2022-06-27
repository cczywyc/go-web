package main

import "unicode/utf8"

func main() {
	println(len("你好"))
	println(utf8.RuneCountInString("你好"))
}
