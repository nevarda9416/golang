/*
*
Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.
*/
package main

import "fmt"

func ỉntSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	nextInt := ỉntSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	nextInts := ỉntSeq()
	fmt.Println(nextInts())
}
