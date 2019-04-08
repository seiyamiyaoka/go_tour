package main

import (
	"fmt"
	"math"
	"testing"

	L "./lib"
)

func TestAbs(t *testing.T) {
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs(-1) = %f; wand 1", got)
	}
}

func TestSwap(t *testing.T) {
	result_s, result_i := L.Swap(5, "tarou")

	fmt.Println(result_s)
	fmt.Println(result_i)
}

func TestVariable(t *testing.T) {
	result := L.Variable()
	if result != 0 {
		t.Errorf("L.variable = %d; want 0", result)
	}
}

func TestConvert(t *testing.T) {
	num_of_string := "12"
	result := L.Convert(num_of_string)
	if result != 12 {
		t.Errorf("L.convert = %d; want int 12", result)
	}
}

func TestSumNum(t *testing.T) {
	result := L.Sum(10)
	if result != 55 {
		t.Errorf("L.Sum = %d; want int 55", result)
	}
}

func TestSwitchButton(t *testing.T) {
	result := L.SwitchButton("neko")
	if result != "cat button" {
		t.Errorf("L.SwitchButton = %s; want string cat button", result)
	}
	result = L.SwitchButton("cofee")
	if result != "dob button" {
		t.Errorf("L.SwitchButton = %s; want string dob button", result)
	}
}

func TestLibPointer(t *testing.T) {
	result := L.GetPointer()
	if result != 5 {
		t.Errorf("L.getPointer = %d; want to int 5", result)
	}
}

func TestLibUserStruct(t *testing.T) {
	result := L.User{"tarou", 12}

	if result.Name != "tarou" {
		t.Errorf("User struct has attributes of name; target: %s", result.Name)
	}
	v := &result
	v.Name = "neko"
	if v.Name != "neko" {
		t.Errorf("stuructのpointerが理解できていません")
	}
	(*v).Name = "kujira"
	if (*v).Name != "kujira" {
		t.Errorf("省略形の書き方がが理解できていません")
	}
}

func TestArrayCode(t *testing.T) {
	result := L.ArrayValue()
	if len(result) != 5 {
		t.Errorf("長さが違うよ; %d", len(result))
	}
}

func TestCopyArrayForSlice(t *testing.T) {
	result := L.CopyArray()
	if result[0] != 5 {
		t.Errorf("一番目はidである必要があります; %d", result[0])
	}
}

func TestGetSummaryOfData(t *testing.T) {
	nums := []int{1, 2, 3}
	result_i_1, result_i_2, result_ary := L.GetSummary(nums)
	if result_i_1 != 3 && result_i_2 != 3 && result_ary[0] != 1 {
		t.Errorf("要素数が正しくありません")
	}
}
func TestSlicingAry(t *testing.T) {
	nums := []int{1, 2, 3}
	// slicingする slicingしたものをL.GetSummaryに入れる
	result_i_1, result_i_2, result_ary := L.GetSummary(nums)
	if result_i_1 != 3 && result_i_2 != 3 && result_ary[0] != 1 {
		t.Errorf("要素数が正しくありません")
	}
}

func TestGoMake(t *testing.T) {
	result := L.GoMake()
	fmt.Println(result)
	if cap(result) != 5 {
		t.Errorf("makeの使い方間違っているかも")
	}
}

func TestCreateTicTocToeBoard(t *testing.T) {
	result := L.CreateTicTocToeBoard()
	if len(result) != 3 {
		t.Errorf("それはtictactoeじゃないよ")
	}
}

func TestRangePow(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	expect_num := []int{2, 4, 6, 8, 10, 12}
	result := L.RangePow(nums)
	if result[len(result)-1] != expect_num[len(result)-1] {
		t.Errorf("返される要素数があっていませんよ; %v", result)
	}
}

func TestRangeContinue(t *testing.T) {
	nums := 256
	result := L.RangeContinue(nums)
	// fmt.Println(result)
	if len(result) != 256 && len(result[0]) != 256 {
		t.Errorf("rangeの省略がうまくできていません; return value is %v", result)
	}
}

func TestMapVertex(t *testing.T) {
	result := L.MapVertex()
	// fmt.Println(result)
	if result["address"].X != 40.68433 {
		t.Errorf("住所が正しくありません")
	}
}

func TestMapLiteral(t *testing.T) {
	contries := []string{"jp", "en"}
	result := L.MapLiteral(contries)
	fmt.Println(result)
	if len(result) != 2 {
		t.Errorf("リテラルの要素数が正常ではありません")
	}
}

func TestWordCounter(t *testing.T) {
	word := "hi i am tarou"

	result := L.WordCounter(word)
	if result["hi"] != 1 {
		t.Errorf("単語の数間違っているよ")
	}
}

func TestLibCompute(t *testing.T) {
	result := L.Compute(math.Pow)
	if result != 81 {
		t.Errorf("関数に関数を渡した結果が間違っているよ")
	}
}

func TestCloujureSum(t *testing.T) {
	var result int
	add := L.CloujureSum()
	for i := 0; i < 5; i++ {
		result = add(i)
	}

	if result != 10 {
		t.Errorf("クロージャを使って関数が正しく定義できていません")
	}
}

func TestFibNat(t *testing.T) {
	var result int
	fib := L.FibNat()
	for i := 0; i < 10; i++ {
		result = fib()
	}
	if result != 34 {
		t.Errorf("フィボナッチ数列の計算があっていないよ; got: %d", result)
	}
}

func TestUserName(t *testing.T) {
	tarou := L.User{"tarou", 34}
	if tarou.GetName() != "tarou" {
		t.Errorf("メソッドのレシーバが正しくセットできていません")
	}
}

func TestNumReceiver(t *testing.T) {
	miya_num := L.MiyaokaNum(10)
	if miya_num.AddTen() != 20 {
		t.Errorf("レシーバが間違っているかも; got: %d", miya_num)
	}
}

func TestUpdateName(t *testing.T) {
	tarou := L.User{"tarou", 12}
	tarou.UpdateName("tarou_second")
	if tarou.GetName() != "tarou_second" {
		t.Errorf("ポインタレシーバの使い方が間違っているかも; got: %v", tarou.GetName())
	}
}

func TestUpdateAge(t *testing.T) {
	// tarou := &L.User("tarou", 10)
	// L.UpdateAge(tarou, 20)
	// これでも可能
	tarou := L.User{"tarou", 19}
	L.UpdateAge(&tarou, 20)
	if tarou.Age != 20 {
		t.Errorf("ポインタの使い方間違ってるかも; got: %v", tarou.Age)
	}
}

func TestUserInterface(t *testing.T) {
	var miyaoka_interface L.UserInterface

	new_nums := L.MiyaokaNum(10)

	miyaoka_interface = new_nums

	if miyaoka_interface.AddTen() != 20 {
		t.Errorf("正しくインターフェースがセットできていません;")
	}

	if miyaoka_interface.AddNum(15) != 25 {
		t.Errorf("正しくインターフェースに引数が渡せてないかも; %d", miyaoka_interface.AddNum(15))
	}
}
