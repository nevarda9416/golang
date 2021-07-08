// Slices are a key data type in Go, giving a more powerful interface to sequences than arrays
/**
Unlike arrays, slices are typed only by the elements they contain (not the number of elements). To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued)
We can set and get just like with arrays
len returns the length of the slice as expected
*/
package main
import "fmt"
func main() {
	s:=make([]string,3)
	fmt.Println("emp:",s)
	s[0]="a"
	s[1]="b"
	s[2]="c"
	fmt.Println("set:",s)
	fmt.Println("get:",s[2])
	fmt.Println("len:",len(s))
	s=append(s,"d") // = is for assignment only
	s=append(s,"e","f") // = is for assignment only
	fmt.Println("apd:",s)
	c:=make([]string,len(s)) // := is for declaration + assignment
	copy(c,s)
	fmt.Println("cpy:",c)
	l:=s[2:5] // := is for declaration + assignment
	fmt.Println("sl1:",l)
	l=s[:5] // = is for assignment only
	fmt.Println("sl2:",l)
	l=s[2:] // = is for assignment only
	fmt.Println("sl3:",l)
	t:=[]string{"g","h","i"} // := is for declaration + assignment
	fmt.Println("dcl:",t)
	twoD:=make([][]int,3) // := is for declaration + assignment
	for i:=0;i<3;i++ {
		innerLen:=i+1
		twoD[i]=make([]int, innerLen)
		for j:=0;j<innerLen;j++ {
			twoD[i][j]=i+j
		}
	}
	fmt.Println("2d: ",twoD)
}