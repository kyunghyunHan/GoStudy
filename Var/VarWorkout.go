package Var
import "fmt"

func Q1(num int){
	
	for i := 1;i<num+1 ;i++{
		num2:=i*10
		num3:=num2+10
		fmt.Printf("타잔이 %v 팬티를 입고%v원짜리 칼을 차고 노래를한다\n",num2,num3)
	}
//다음 노래를 N줄 출력하는 함수를 작성
//타잔이 10원짜리 팬티를 입고 20원짜리 칼을 차고 노래를한다
//타잔이 20원짜리 팬티를 입고 30원짜리 칼을 차고 노래를 한다

}
// func Q2(){
// 	//하노이의 탑 

// }


func Q2(n int) int {
	//파보나치 수열
	var index int
	var last1 int
	var last2 int
	var result int

	if n <= 2 {
		return 1
	}
	last1 = 1
	last2 = 1
	for index = 2; index < n; index++ {
		result = last1 + last2
		last1 = last2
		last2 = result
	}
	return result
}
func Result2(){
	Q1(2)
	fmt.Println(Q2(3))
}