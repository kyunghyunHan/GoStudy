package Func

import (
	"fmt"
	"io"
)

func AddOnd(nums []int) {
	for i := range nums {
		nums[i]++
	}
}
func ExampleAddone() {
	n := []int{1, 2, 3, 4}
	AddOnd(n)
	fmt.Println(n)

}
func writeTo(w io.Writer, lines []string) (n int64, err error) {
	for _, line := range lines {
		var nw int
		nw, err = fmt.Fprintln(w, line)
		n += int64(nw)
		if err != nil {
			return
		}
	}
	return
}
func Result() {
	ExampleAddone()
}
