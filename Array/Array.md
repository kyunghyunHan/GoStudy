# 배열과 슬라이스
- 배열과 슬라이스 모두 연속된 메모리 공간을 순차적으로 이용하는 자료구조

## 베열
- 
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

## 슬라이스
- 기본적으로 빈 슬라이스에는 nil값이 들어감
```go
var fruits []string
```
- 빈스트링을 n개 가지고 있는 슬라이드
```go
fruits :=make([]string ,n)
```
- 반복문을 이용하여 i번째 문자열에 접근
```go
fruist := make([]string,n)

fruist[0]="사과"
fruist[1]="바나나"
fruist[2]="토마토"
```
- make로 크기를 잡고 만든 슬라이스에는 자료형의 기본값이 들어간다
- 슬라이스를 자를떄는
```go
func Example_Sliceing(){
	nums := []int{1,2,3,4,5}
	fmt.Println(nums)
	fmt.Println(nums[1:3],"nums[1:3]")
	fmt.Println(nums[2:],"nums[2:]")
	fmt.Println(nums[3:],"nums[3:]")
}
```
## 슬라이스 덧붙이기
- append
```go
fruits= append(fruits,"포도")
```
- 여러개 덧붙이기
```go
fruits= append(fruits,"포도","딸기")
```
- 두슬라이스 이어붙이기
```go
func Example_appned(){
	f1:= []string{"사과","바나나","토마토"}
	f2:= []string{"포도","딸기"}
	f3:= append(f1,f2...) //이어붙이기
	f4:= append(f1[:2],f2...)//토마토를 제외하고 이어붙이기

	fmt.Println(f1,"f1")
	fmt.Println(f2,"f2")
	fmt.Println(f3,"f3")
	fmt.Println(f4,"f4")
}
```
## 슬라이스용량
- 슬라이스는 연속된메모리 공간을 활용하는 것이라서 용량에 제한
- make([]int,5)와 같이 다섯개의 빈공간으 미리 할당하거나 []int{0,0,0,0,0}과 같이 다섯개의 정로 최기화 한 경우 길이뿐만 아니라 용량도 5
```go
func Example_SliceCap() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(nums)
	fmt.Println("lens", len(nums))
	fmt.Println("caps:", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("lens", len(sliced1))
	fmt.Println("caps:", cap(sliced1))
	fmt.Println()

	sliced2 := nums[2:]
	fmt.Println(sliced2)
	fmt.Println("lens", len(sliced2))
	fmt.Println("caps:", cap(sliced2))
	fmt.Println()

	sliced3 := nums[:4]
	fmt.Println(sliced3)
	fmt.Println("lens", len(sliced3))
	fmt.Println("caps:", cap(sliced3))
	fmt.Println()

	nums[2] = 100
	fmt.Println(nums, sliced1, sliced2, sliced3)
}
```
- 용량을 미리 지정해서 생성
- 길이는 3이지만 용량은 6
```go
nums := make([]int,3,5)

```
## 슬라이스의 내부 구현
- 슬라이스는 배열을 가리키고 있는 구조체
- 슬라이스는 시작주소, 길이 , 용량

```go
nums = append(nums,10)
```
## 슬라이스 복사

```go
func Example_sliceCopy() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
	for i := range src {
		dest[i] = src[i]
	}
	fmt.Println(dest)
}

```
- copy
```go
copy(dest,src)
```
- dest 공간이 충분하지 않은지 확인
```go
if n:= copy(dest,src); n!= len(src){
	fmt.Println("복사가 덜 됫습니다")
}
```
- copy로 복사하려면 다음과 같은
```go
src := []int{30, 20, 50, 10, 40}
dest:= make([]int,len(src))
copy(dest,src)
```
- append 이용
```go 
src := []int{30, 20, 50, 10, 40}
dest:= append([]int(nil),src...)
```

## 슬라이스 삽입 및 삭제

- GO에서는 따 제공하지 않음
- 삭제

```go
a[i]= a[len(a)-1]
a= a[:len(a)-1]
```
- k개를 삭제
```go
start := len(a)-k
if i+k >start {
	i+k
}
copy(a[i:i+k],a[start:])
a= a[:len(a)-k]
```
- 순서를 유지하면서 i 번쨰를 삭제
```go
copy(a[i:],a[i+1])
a[len(a)-1]= nil//생략시 메모리 누수 위험
a=a[:len(a)-1]
```
- 한꺼번에 연속된 k개를 지우는 경우에는 마지막 k개를 지우고 다스 슬라이싱
```go
copy(a[i:],a[i+k:])
for i := 0; i <k ; i++{
	a[len(a)-1-i]=nil

}
a= a[:len(a)-k]
```
## 스택
- 후입선출 구조를 가지고 있는 자료구조
- 함수의 호출과 반환 역시 내부적으로 스택 구현
- 스택을 이용 사칙연산과 괄호를 이용한 간단한 정수 수식 계산기
## 맵

- go언어에서 맵은 해시테이블로 구현
```go
var m map[keyType]vauleType
```
- 빈맵으로 취급대어 맵을 읽을수 있지만 변경할수가 없음
- nil값을 가지고 있는 슬라이스에 appen로 덧붙일수 있는 것과 달리 맵은 일단 생성이 되어야 추가가능
```go
m:= make(map[keyType]valueType)
혹은
m:= map[keyType]valueType
```
- 맵에서 읽을떄는 두가지 방법
- m[key]를 이용하여 맵의 값을 읽을수 있음
```go
value,ok:=m[key]
```
- 맵 사용
```go
func count(s string,codeCount map[rune]int){
	for _,r:= range s {
		codeCount[r]++
	}
}


```
- 맵을 range로 반복하게 되면 키와 값이 값이 나오게 댄다 하나의 변수로만 받게되면 키만 나온다
```go
for _,v:= range m{
	//
}
```
```go
func hasDupeRuns(s string) bool {
	runeSet := map[rune]bool{}
	for _, r := range s {
		if runeSet[r] {
			return true
		}
		runeSet[r] = true
	}
	return false
}
//struct은 빈구조 자료향

func hasDupeRuns2(s string) bool {
	runeSet := map[rune]struct{}{}
	for _, r := range s {
		if _, exits := runeSet[r]; exits {
			return true
		}
		runeSet[r] = struct{}{}
	}
	return false

}

```
- 맵 삭제
```go
delete(m,key)
```
- 집합 (문자열 중에 중복되어 들어가 있는 글자가 있는지 검사하는 함수)
## 입출력
- 입출력은 io.Reader와 io.Wruter인터페이스와 파생된 다른 인터페이스들을 이용

## 파일읽기
- os.Open은 반환값이 둘
```go
f,err := os.Open(filename)
if err != nil{
	return err

}
defer f.Close()
var num int
if _,err := fmt.Fscanf(f,"$d\n",&num);err ==nil{
	//읽은 num값 사용
}
```

## 파일쓰기
 ```go
 f,err := os.Create(filename)
 if err != nil{
	 return nil//혹은 다른 에러처리
 }
 defer f.Close()
 if _,err := fmt.Fprintf(f,"%d\n",num);err!= nil{
	 return err
 }
 ```

 ## 텍스트 리스트 읽고 쓰기
 - 단순히 문자열 슬라이스를 라인별로 출력하는 함수
 ```go
 func WriteTo(w io.Writer, lines []string)error{
	 for _,line := range lines{
		 if_,err := fmt.Fprintln(w,line); err != nil{
			 return err
		 }
	 }
	 return nil
 }
 ```
 