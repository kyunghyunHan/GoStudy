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
func Example_printBytes(){
	s:= "가나다"
	for i := 0; i<len(s); i++{
		fmt.Printf("%x:",s[i])
	}
	fmt.Println()
}
func Result(){
	Example_printBytes()
	String()
	fmt.Println(HasConsonanSuffix("안녕"))
}