package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var conv = morse.DefaultConverter

func Convert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return "", fmt.Errorf("входная строка пуста")
	}
	isMorse, err := isMorseCode(input)
	if err != nil {
		return "", err
	}
	if isMorse {
		result := conv.ToText(input)
		if result == "" {
			return "", fmt.Errorf("не удалось преобразовать морзе в текст")
		}
		fmt.Println("преобразование из морзе успешно")
		return result, nil
	} else {
		result := conv.ToMorse(input)
		if result == "" {
			return "", fmt.Errorf("не удалось преобразовать текст в морзе")
		}
		fmt.Println("преобразование в морзе успешно")
		return result, nil
	}
}

func isMorseCode(s string) (bool, error) {
	if len(s) == 0 {
		return false, fmt.Errorf("пустая строка")
	}
	for _, ch := range s {
		if !(ch == '.' || ch == '-' || ch == ' ' || ch == '/') {
			return false, nil
		}
	}
	return true, nil
}
