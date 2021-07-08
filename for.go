// for is Go's only looping construct. Here are some basic types of for loops
/**
The most basic type, with a single condition
A classic initial/condition/after for loop
*/
package main
import "fmt"
func main() {
	i:=1
	for i<=3 {
		fmt.Println(i)
		i = i+1
	}
	for j:=7;j<=9;j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("loop")
		break
	}
	for n:=0;n<=5;n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}