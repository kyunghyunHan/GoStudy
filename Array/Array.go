package Array

import (
	"fmt"
	"strconv"
	"strings"
)

func Example_array() {
	fruits := [3]string{"사과", "바나나", "토마토"}
	for _, fruit := range fruits {
		fmt.Printf("%s 는 맛있다.\n", fruit)
	}
}
func Example_Sliceing() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3], "nums[1:3]")
	fmt.Println(nums[2:], "nums[2:]")
	fmt.Println(nums[3:], "nums[3:]")
}
func Example_appned() {
	f1 := []string{"사과", "바나나", "토마토"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...)     //이어붙이기
	f4 := append(f1[:2], f2...) //토마토를 제외하고 이어붙이기

	fmt.Println(f1, "f1")
	fmt.Println(f2, "f2")
	fmt.Println(f3, "f3")
	fmt.Println(f4, "f4")
}
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
func Example_sliceCopy() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
	for i := range src {
		dest[i] = src[i]
	}
	fmt.Println(dest)
}
func addddd() {
	nums := make([]int, 5)
	nums = []int{1, 2, 3, 4, 5}
	// nums = nums[:3]
	fmt.Println(nums[1:], "tmffkdltm")
}
func Eval(expr string) int {
	var ops []string
	var nums []int
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}
	reduce := func(higher string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if strings.Index(higher, op) < 0 {
				//목록에 없는 연산자 이므로 종료
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				//괄호룰 제거햇으므로 종료
				return
			}
			b, a := pop(), pop()
			switch op {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a-b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)

			}

		}
	}
	for _, token := range strings.Split(expr, " ") {
		switch token {
		case "(":
			ops = append(ops, token)
		case "+", "-":
			//덧샘과 뺼셈 이상의 우선순위를 가진 사칙연산 적용
			reduce("+-*/")
			ops = append(ops, token)
		case "*", "/":
			//곱셈과 나눗셈 이상의 우선순위를 가진것은 둘뿐
			reduce("*/")
			ops = append(ops, token)
		case ")":
			//다는 괄호는 여는 괄호까지 계산하고 제거
			reduce("+-*/(")
		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	reduce("+-*/")
	return nums[0]
}

// func ExampleReadFrom() {
// 	r := strings.NewReader("bill\ntom\njane\n")
// 	var lines []string
// 	if err := ReadFrom(r, &lines); err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(lines)
// }
func Result() {
	Example_array()
	Example_Sliceing()
	Example_appned()
	Example_SliceCap()
	Example_sliceCopy()
	addddd()
	fmt.Println(Eval("5"))
	fmt.Println(Eval("1 + 2"))
	fmt.Println(Eval("1 - 2 + 3"))
	fmt.Println(Eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))
}
