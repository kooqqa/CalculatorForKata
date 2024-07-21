package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romarab = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var arabrom = []string{
	"",
	"I",
	"II",
	"III",
	"IV",
	"V",
	"VI",
	"VII",
	"VIII",
	"IX",
	"X",
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение:")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

	result, err := calculate(expression)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func calculate(expression string) (string, error) {
	parts := strings.Fields(expression)
	if len(parts) != 3 {
		return "", errors.New("неверный формат выражения")
	}

	a, b, operator := parts[0], parts[2], parts[1]
	aNum, aIsRoman := parseNumber(a)
	bNum, bIsRoman := parseNumber(b)

	if aIsRoman != bIsRoman {
		return "", errors.New("используйте либо только арабские, либо только римские цифры")
	}

	if aNum < 1 || aNum > 10 || bNum < 1 || bNum > 10 {
		return "", errors.New("числа должны быть от 1 до 10 включительно")
	}

	result := 0
	switch operator {
	case "+":
		result = aNum + bNum
	case "-":
		result = aNum - bNum
	case "*":
		result = aNum * bNum
	case "/":
		if bNum == 0 {
			return "", errors.New("деление на ноль")
		}
		result = aNum / bNum
	default:
		return "", errors.New("неизвестный оператор")
	}

	if aIsRoman {
		if result < 1 {
			return "", errors.New("римские числа должны быть положительными")
		}
		return intToRoman(result), nil
	}

	return strconv.Itoa(result), nil
}

func parseNumber(s string) (int, bool) {
	if val, ok := romarab[s]; ok {
		return val, true
	}
	val, err := strconv.Atoi(s)
	if err != nil {
		panic("неверный формат числа")
	}
	return val, false
}

func intToRoman(num int) string {
	vals := []int{10, 9, 5, 4, 1}
	symbols := []string{"X", "IX", "V", "IV", "I"}
	roman := ""

	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			num -= vals[i]
			roman += symbols[i]
		}
	}
	return roman
}
