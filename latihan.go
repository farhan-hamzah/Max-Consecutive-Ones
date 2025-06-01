package main
import "fmt"
const NMAX int = 100
type tabInt[NMAX]int
func main(){
	var A tabInt
	var n, i, c, temp int
	fmt.Scan(&n)
	for i = 0; i < n;i++{
		fmt.Scan(&A[i])
	}
	temp = 0
	c = 0
	for i = 0; i < n; i++{
		if A[i] == 1{
			c++
		}else{
			c = 0
		}
		if temp < c{
			temp = c
		}
	}
	fmt.Print(temp)
}