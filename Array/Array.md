# 배열과 슬라이스
- 배열과 슬라이스 모두 연속된 메모리 공간을 순차적으로 이용하는 자료구조

## 베열

```go 
func Example_array(){
    fruits := [3]string{"사과","바나나","토마토"}
    for _,fruit := range fruits{
        fmt.Println("%s는 맛있다.\n",fruit)
    }
}
//Output
//사과 는 맛있다.
//바나나 는 맛있다.
//토마토 는 맛있다.
```