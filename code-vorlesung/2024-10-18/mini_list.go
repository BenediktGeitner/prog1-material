package lists

import (
	"fmt"
)

func ListDemo() {

	//Definiere eine Liste von int
	//[1 3 5 7 9]
	list1 := []float64{1, 3.5, 5, 7, 9}

	//gib das erste Element der Liste auf der Konsole
	fmt.Println(list1[0])
	//gib das letzte Element der Liste auf der Konsole
	fmt.Println(list1[4])
	fmt.Println()
	//Gib ein Element an ungültiger Stelle aus
	//fmt.Println(list1[5])

	//Gib alle Elemente der Liste aus
	for n := 0; n < len(list1); n++ {
		fmt.Println(list1[n])
	}

	fmt.Println(list1[len(list1)-1])
}
