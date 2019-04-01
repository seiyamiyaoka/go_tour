package lib

import (
	"fmt"
	"strconv"
)

var i int

func Swap(x int, y string) (string, int) {
	fmt.Println("first return type is %d", x)
	return y, x
}

func Variable() int {
	return i
}

func Convert(num string) int {
	i, _ := strconv.Atoi(num)
	return i
}

func Sum(num int) int {
	result := 0
	for i := 0; i <= num; i++ {
		result += i
	}
	return result
}

func SwitchButton(name string) string {
	switch button := name; button {
	case "neko":
		return "cat button"
	default:
		return "dob button"
	}
}

func StackDefer() [3]string {
	var names [3]string
	users := [3]string{"tarou", "jiro", "saburo"}

	for i := 0; i < len(users); i++ {
		defer fmt.Println(users[i])
	}
	return names
}

func GetPointer() int {
	i := 5
	p := &i
	return *p
}

type User struct {
	Name string
	Age  int
}

func ArrayValue() [5]int {
	var nums [5]int
	for ; i < len(nums); i++ {
		nums[i] = i
	}
	return nums
}

func CopyArray() []int {
	var profile [3]int
	profile[0] = 1
	profile[1] = 5
	profile[2] = 19
	var s []int = profile[1:2]
	return s
}

func GetSummary(ary []int) (int, int, []int) {
	nums := ary
	return len(nums), cap(nums), nums
}

func GoMake() []int {
	nums := make([]int, 0, 5)
	return nums
}

func CreateTicTocToeBoard() [3][]string {
	boards := [3][]string{}
	for i := 0; i < 3; i++ {
		boards[i] = createTicTocLine()
	}
	fmt.Println(boards)
	return boards
}

func createTicTocLine() []string {
	line := []string{
		"_", "_", "_",
	}
	return line
}

func RangePow(nums []int) []int {
	new_nums := nums
	for i, v := range nums {
		new_nums[i] = 2 * v
	}
	return new_nums
}
