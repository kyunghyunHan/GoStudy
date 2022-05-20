package String

import "fmt"

func String(){
	for i ,r := range "가나다" {
		fmt.Println(i,r)
	}
	fmt.Println(len("가나다"))
	fmt.Println("문자열")
	for _,r := range "가가히" {
		fmt.Println(string(r),r)
	}
}

func Result(){

	String()
}