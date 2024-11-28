package ean

// Is Digit erwartt ein Zeichen c und prüft
// ob dieses eine Ziffer ist
func IsDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func InputOk(code string) bool {
	if len(code) != 13 {
		return false
	}
	for i := 0; i < len(code); i++ {
		if !IsDigit(code[i]) {
			return false
		}
	}
	return true
}

// ToIntList erwarter einen String "code", der nur aus Ziffern besteht
// Liefert eine Liste dieser Ziffern als []int.
func ToIntList(code string) []int {
	var result []int
	for i := 0; i < len(code); i++ {
		num := int(code[i] - '0')
		result = append(result, num)
	}
	return result
}

//EanOk erwartet einen String und liefert true
//genau dann wenn code eine gültige EAN-Nummer ist

func EanOk(code string) bool {
	if !InputOk(code) {
		return false
	}

	digits := ToIntList(code)

	checksum := CheckSum(digits)

	return checksum == 0
}

// Checksum erwartet eine Liste von Ziffern
// berechnet deren EAN-Prüfsumme
func CheckSum(digits []int) int {

	//TODO
	result := 0
	for i, n := range digits {
		if i%2 == 0 {
			result += n
		} else {
			result += 3 * n
		}
	}
	return (10 - (result % 10)) % 10
}
