package strutil

import (
	"strings"
)

const DEFAULT_FILELENGTH_LIMIT = 50

var acceptableChars = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-', '_',
}

var translitMap = map[rune]string{
	'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d", 'е': "e", 'ё': "yo", 'ж': "zh",
	'з': "z", 'и': "i", 'й': "y", 'к': "k", 'л': "l", 'м': "m", 'н': "n", 'о': "o",
	'п': "p", 'р': "r", 'с': "s", 'т': "t", 'у': "u", 'ф': "f", 'х': "kh", 'ц': "ts",
	'ч': "ch", 'ш': "sh", 'щ': "shch", 'ъ': "", 'ы': "y", 'ь': "", 'э': "e", 'ю': "yu",
	'я': "ya",
	'А': "A", 'Б': "B", 'В': "V", 'Г': "G", 'Д': "D", 'Е': "E", 'Ё': "Yo", 'Ж': "Zh",
	'З': "Z", 'И': "I", 'Й': "Y", 'К': "K", 'Л': "L", 'М': "M", 'Н': "N", 'О': "O",
	'П': "P", 'Р': "R", 'С': "S", 'Т': "T", 'У': "U", 'Ф': "F", 'Х': "Kh", 'Ц': "Ts",
	'Ч': "Ch", 'Ш': "Sh", 'Щ': "Shch", 'Ъ': "", 'Ы': "Y", 'Ь': "", 'Э': "E", 'Ю': "Yu",
	'Я': "Ya",
}

var charsToTrim = []rune{
	' ', ',', '.',
}

func TrimByChars(input string) string {
	start := 0
	end := len(input)

	for start < end {
		trimmed := false
		for _, char := range charsToTrim {
			if start < end && rune(input[start]) == char {
				start++
				trimmed = true
			}
			if start < end && rune(input[end-1]) == char {
				end--
				trimmed = true
			}
		}
		if !trimmed {
			break
		}
	}

	return input[start:end]
}

func GetLastPartOfURL(input string) string {
	lastSlashIndex := strings.LastIndex(input, "/")
	if lastSlashIndex != -1 {
		input = input[lastSlashIndex+1:]
	}
	questionMarkIndex := strings.Index(input, "?")
	if questionMarkIndex != -1 {
		input = input[:questionMarkIndex]
	}
	dotIndex := strings.LastIndex(input, ".")
	if dotIndex != -1 {
		input = input[:dotIndex]
	}
	return input
}

func FilterAcceptableChars(input string, limit *int) string {

	defaultLimit := DEFAULT_FILELENGTH_LIMIT
	if limit != nil {
		defaultLimit = *limit
	}

	if len(input) > defaultLimit {
		input = input[:defaultLimit]
	}

	var result strings.Builder
	for _, char := range input {
		for _, acceptableChar := range acceptableChars {
			if char == acceptableChar {
				result.WriteRune(char)
				break
			}
		}
	}
	return result.String()
}

func Transliterate(input string, limit *int) string {
	// Транслитерация и ограничение длины строки до limit символов
	defaultLimit := DEFAULT_FILELENGTH_LIMIT
	if limit != nil {
		defaultLimit = *limit
	}

	var result strings.Builder
	for _, char := range input {
		if val, ok := translitMap[char]; ok {
			result.WriteString(val)
		} else {
			result.WriteRune(char)
		}
	}

	if result.Len() > defaultLimit {
		return result.String()[:defaultLimit]
	}

	return result.String()
}
