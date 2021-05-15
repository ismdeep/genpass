package main

// IsFuzzyChar Is fuzzy char
func IsFuzzyChar(str string) bool {
	if str == "0" || str == "o" || str == "O" || str == "I" || str == "1" || str == "l" || str == "|" {
		return true
	}
	return false
}

// RemoveFuzzy Remove Fuzzy from string
func RemoveFuzzy(str string) string {
	result := ""
	for _, chInt := range str {
		if !IsFuzzyChar(string(chInt)) {
			result += string(chInt)
		}
	}
	return result
}
