# 문자열
- 바이트들이 연속적으로 나ㄹ

## 유니코드 처리
- 전세계 모든 문지를 일관되게 표현하기 위한 산업 표준
```go
for i ,r := range "가나다" {
    fmt.Println(i,r)
}
fmt.Println(len("가나다"))
```
- 가가히의 유니코드출력
```go
for _,r := range "가가히" {
    fmt.Println(string(r),r)
}
```
```go
//한글이 받침이 있는지 없는지 확인
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
```
## 바이트단위 처리
```go
func Example_printBytes(){
	s:= "가나다"
	for i := 0; i<len(s); i++{
		fmt.Printf("%x:",s[i])
	}
	fmt.Println()
}
```
```go
func Example_printBytes2(){
	s:= "가나다"
	fmt.Printf("%x\n",s)
	fmt.Printf("% x\n",s)
}
```
## 문자열 잇기
- go의 문자열은 사실상 문자열에 대한 포인터와 비슷하기 떄문에 이어붙이면
```go
func Example_StrCat(){
	s:= "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
}
```
## 문자열 숫자로
- strconv패키지의Atoi함수로 문자열을 정수로 변경
- strconv패키지의ParseInt()함수로 64비트 혹은 10진수로 변경
```go
	var i int
	var k int64
	var f float64
	var s string
	var err error
	i, err= strconv.Atoi("350")
	k,  err= strconv.ParseInt**("cc7fdd",16,32)
	k, err = strconv.ParseInt("0xcc7fdd",0,32)
	f, err = strconv.ParseInt***("3.14",64)
	s= strconv.Itoa****(340)
	s= strconv.FormatInt****(13402077,16)
    fmt.Println(i,err)
    ```