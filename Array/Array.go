package Array

import "fmt"

func Example_array(){
    fruits := [3]string{"사과","바나나","토마토"}
    for _,fruit := range fruits{
        fmt.Printf("%s 는 맛있다.\n",fruit)
    }
}

func Result(){
	Example_array()
}