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

// func TestStackDefer(t *testing.T) {
// 	result := L.StackDefer()
// 	expect := [3]string{"tarou", "jiro", "saburo"}
// 	if result != expect {
// 		t.Errorf("L.StackDefer = %v; want to array ", result)
// 	}
// }
