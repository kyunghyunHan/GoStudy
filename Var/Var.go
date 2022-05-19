package Var
import "fmt"


func fac(n int)int{
	if n <=0 {
		return 1
	}
	return n * fac(n-1)
}

func fucItr (n int)int{
	result :=1
	for n>0{
		result *= n
		n--
	}
	return result
}

func facItr2(n int)int{
result := 1
for i :=2; i<=n; i++{
	result *= i
}
return result
}
func Result(){
	fmt.Println(facItr2(5))
     fmt.Println(fac(5))
	 fmt.Println(fucItr(5))
	fmt.Println("hello world")
}