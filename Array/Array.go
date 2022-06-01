package Array

import (
	"fmt"
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
func ExampleReadFrom() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string
	if err := ReadFrom(r, &lines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
}
func Result() {
	Example_array()
	Example_Sliceing()
	Example_appned()
	Example_SliceCap()
}
