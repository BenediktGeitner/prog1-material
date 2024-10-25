package findlist

//FindElement erwartet eine Liste von Zahlen und eine weitere Zahl n
//Die Funktion liefert die Position von n in der Liste
//Falls n nicht in der Liste vorkommt, liefert die Funktion

func FindElementLinear(list []int, n int) int {
	for i := 0; i < len(list); i++ {
		if list[i] == n {
			return i
		}
	}
	return -1
}

//FindElement erwartet eine Liste von Zahlen und eine weitere Zahl n
//Die Funktion liefert die Position von n in der Liste
//Falls n nicht in der Liste vorkommt, liefert die Funktion
//Vorraussetzung Liste muss sortiert sein

func FindElementBinary(list []int, n int) int {
	if len(list) == 0 {
		return -1
	}
	m := len(list) / 2
	if list[m] == n {
		return m
	}

	if n < list[m] {
		return FindElementBinary(list[:m], n)
	} else {
		result := FindElementBinary(list[m+1:], n) + m + 1
		if result == -1 {
			return -1
		}
		return result
	}

}
