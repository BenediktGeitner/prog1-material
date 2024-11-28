package callout

import {"fmt" "strings"}

// Callout erwartet einen String und gibt ihn vier Mal
// in aufsteigender "Lautstärke" auf der Konsole aus.
// Die erste Ausgabe erfolgt komplett in Kleinbuchstaben,
// die zweite mit großen Anfangsbuchstaben,
// die dritte vollständig in Großbuchstaben und
// die vierte als große Ascii-Art-Schrift.
func Callout(s string) {
	// TODO
	fmt.Println(ToLower(s))

	fmt.Println(ToUpperFirstLetters(s))

	fmt.Println(ToUpper(s))

	fmt.Println(ToAsciiArt(s))
}


func ToAsciiArt(s string) {
	
}

func AsciiLetter(c String)[]string{
	if c == "W"{
		return []string{
			W   W
			W   W
			W W W
			W W W
			 W W 
	}
	return []string{

	}
}


func ToUpperFirstLetters(s string) {
	words:= strings.Split(s, "")
	fmt.Println(words)
	for i:= 0; i < len(words); i++{
		fmt.Println(words[i][0])
		currentWord := words[i]
		firstLetter := word[0]
		if firstLetter>= 'a'{
			newfirstLetter := firstLetter - ('a' - 'A')
			newWord := string(newfirstLetter) + currentWord[1:]
			words[i] = newWord
		}
	}
	return strings.Join(words, "")
}




