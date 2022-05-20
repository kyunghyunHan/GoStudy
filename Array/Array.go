package Array

import "fmt"

func Example_array(){
    fruits := [3]string{"사과","바나나","토마토"}
    for _,fruit := range fruits{
        fmt.Printf("%s 는 맛있다.\n",fruit)
    }
}
func Example_StrCat(){
	s:="abc"
	ps:=&s
	s+="def"
	fmt.Println(s)
	fmt.Println(*ps)
}
func Result(){
	Example_array()
	Example_StrCat()
}