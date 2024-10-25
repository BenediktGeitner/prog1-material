package findlist

import "fmt"

func ExampleFindElementLinear() {
	list := []int{1, 3, 5, 7, 9}

	fmt.Println(FindElementLinear(list, 5))

	//Output:
	//2
}

func ExampleFindElementBinary() {
	list := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}

	fmt.Println(FindElementBinary(list, 5))
	fmt.Println(FindElementBinary(list, 99))
	//Output:
	//2
	//-1
}
