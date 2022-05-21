package String

import "fmt"
import "strconv"

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
func Example_StrCat(){
	s:= "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
}
func StrNum(){
	var i int
	// var k int64
	// var f float64
	// var s string
	var err error
	i, err= strconv.Atoi("350")
	// k,  err= strconv.ParseInt**("cc7fdd",16,32)
	// k, err = strconv.ParseInt("0xcc7fdd",0,32)
	// f, err = strconv.ParseInt***("3.14",64)
	// s= strconv.Itoa****(340)
	// s= strconv.FormatInt****(13402077,16)
    fmt.Println(i,err)
}
func Result(){
	Example_printBytes()
	Example_printBytes2()
	Example_modifyBytes()
	Example_StrCat()
	StrNum()
	String()
	fmt.Println(HasConsonanSuffix("안녕"))
}