package ean

import "fmt"

func ExampleInputOk() {
	fmt.Println(InputOk("1231231231231"))
	fmt.Println(InputOk("123123"))
	fmt.Println(InputOk("123123123123123A"))

	//Output:
	//true
	//false
	//false
}

func ExampleIsDigit() {
	fmt.Println(IsDigit('0'))
	fmt.Println(IsDigit('9'))
	fmt.Println(IsDigit('a'))

	//Output:
	//true
	//true
	//false

}

func ExampleToIntList() {
	fmt.Println(ToIntList("1234"))

	//Output:
	//[1 2 3 4]

}

func ExampleCheckSum() {
	fmt.Println(CheckSum([]int{2, 3, 5, 4}))

	//Output:
	//0
}

func ExampleEanOk() {
	fmt.Println(EanOk("2400"))

	//Output:
	//false
}
