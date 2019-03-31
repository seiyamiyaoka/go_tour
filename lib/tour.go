package lib

import (
	"fmt"
	"strconv"
)

var i int

func Swap(x int, y string) (string, int) {
	fmt.Println("first return type is %s", x)
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
