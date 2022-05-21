# 문자열
- 바이트들이 연속적으로 나ㄹ

## 유니코드 처리
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