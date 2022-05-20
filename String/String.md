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