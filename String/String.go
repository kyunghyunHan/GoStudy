package String

import "fmt"

var(
	start = rune(44032)//"가의 유니코드 포인트"
	end = rune(55204)//"힣"다음 글자의 유니코드 포인트
)
func HasConsonanSuffix(s string) bool{
	numEnds :=28
	result := false
	for _,r := range s{
		if start <= r && r< end {
			index:= int(r-start)
			result = index%numEnds != 0
		}
	}
	return result
}
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
//x는 16진수 숫자형식으로
func Example_printBytes(){
	s:= "가나다"
	for i := 0; i<len(s); i++{
		fmt.Printf("%x:",s[i])
	}
	fmt.Println()
}
func Example_printBytes2(){
	s:= "가나다"
	fmt.Printf("%x\n",s)
	fmt.Printf("% x\n",s)
}
func Example_modifyBytes(){
	b:= []byte("가나다")
	b[2]++
	fmt.Println(string(b))
}
func Result(){
	Example_printBytes()
	Example_printBytes2()
	Example_modifyBytes()
	String()
	fmt.Println(HasConsonanSuffix("안녕"))
}