package Func
import "fmt"
func AddOnd(nums []int){
for i := range nums{
	nums[i]++
}
}
func ExampleAddone(){
	n:= []int{1,2,3,4}
	AddOnd(n)
	fmt.Println(n)
	
}

func Result(){
   ExampleAddone()
}