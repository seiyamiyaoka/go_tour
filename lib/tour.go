package lib

import (
	"fmt"
	"strconv"
	"strings"
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

func RangeContinue(nums int) [][]uint8 {
	// uint8はbyteと同じ2進数の8個なので 2 ** 8が最大(256まで)
	binary := make([][]uint8, nums) // uint8のからの状態を作成
	for i := range binary {
		binary[i] = make([]uint8, nums) // 0arrayにする
	}
	for i, row := range binary {
		for j := range row {
			// 行数が全て0では面白くないので数字を作成
			row[j] = uint8(i * j)
		}
	}
	return binary
}

type Vertex struct {
	X float64
}

func MapVertex() map[string]Vertex {
	m := make(map[string]Vertex)

	m["address"] = Vertex{
		40.68433,
	}
	return m
}

func MapLiteral(contries []string) map[string]Vertex {
	m := make(map[string]Vertex)
	for i, v := range contries {
		m[v] = Vertex{
			12.11111 * float64(i+1),
		}
	}
	return m
}

func mapInt() map[string]int {
	var m map[string]int
	return m
}

func MapMutating() map[string]int {
	m := mapInt()
	m["consumer"] = 1
	return m
}

type Consumer struct {
	money int
}

func MapHasKey(data map[string]Consumer, key string) bool {
	_, ok := data[key]
	return ok
}

func WordCounter(words string) map[string]int {
	counter := make(map[string]int)
	for _, word := range strings.Fields(words) {
		_, ok := counter[word]
		if ok == true {
			counter[word] += 1
		} else {
			counter[word] = 1
		}
	}
	return counter
}

func Compute(fn func(x, y float64) float64) float64 {
	return fn(3, 4) // 関数は可変で数値が固定
}

func CloujureSum() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		fmt.Println(sum)
		return sum
	}
}

func FibNat() func() int {
	first := 0
	second := 1
	index := -1
	result := 0
	return func() int {
		index += 1
		if index == 0 {
			return 0
		} else if index == 1 {
			return 1
		}
		result = first + second
		first = second
		second = result
		return result
	}
}

func (u User) GetName() string {
	return u.Name
}

func (u *User) UpdateName(name string) {
	u.Name = name
}

type MiyaokaNum int

func (m MiyaokaNum) AddTen() int {
	return int(m) + 10
}

func (m MiyaokaNum) AddNum(num int) int {
	return int(m) + num
}

func UpdateAge(u *User, age int) {
	u.Age = age
}

type UserInterface interface {
	AddTen() int
	AddNum(num int) int
}
