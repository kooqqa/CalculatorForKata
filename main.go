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
	"I":        1,
	"II":       2,
	"III":      3,
	"IV":       4,
	"V":        5,
	"VI":       6,
	"VII":      7,
	"VIII":     8,
	"IX":       9,
	"X":        10,
	"XI":       11,
	"XII":      12,
	"XIII":     13,
	"XIV":      14,
	"XV":       15,
	"XVI":      16,
	"XVII":     17,
	"XVIII":    18,
	"XIX":      19,
	"XX":       20,
	"XXI":      21,
	"XXII":     22,
	"XXIII":    23,
	"XXIV":     24,
	"XXV":      25,
	"XXVI":     26,
	"XXVII":    27,
	"XXVIII":   28,
	"XXIX":     29,
	"XXX":      30,
	"XXXI":     31,
	"XXXII":    32,
	"XXXIII":   33,
	"XXXIV":    34,
	"XXXV":     35,
	"XXXVI":    36,
	"XXXVII":   37,
	"XXXVIII":  38,
	"XXXIX":    39,
	"XL":       40,
	"XLI":      41,
	"XLII":     42,
	"XLIII":    43,
	"XLIV":     44,
	"XLV":      45,
	"XLVI":     46,
	"XLVII":    47,
	"XLVIII":   48,
	"XLIX":     49,
	"L":        50,
	"LI":       51,
	"LII":      52,
	"LIII":     53,
	"LIV":      54,
	"LV":       55,
	"LVI":      56,
	"LVII":     57,
	"LVIII":    58,
	"LIX":      59,
	"LX":       60,
	"LXI":      61,
	"LXII":     62,
	"LXIII":    63,
	"LXIV":     64,
	"LXV":      65,
	"LXVI":     66,
	"LXVII":    67,
	"LXVIII":   68,
	"LXIX":     69,
	"LXX":      70,
	"LXXI":     71,
	"LXXII":    72,
	"LXXIII":   73,
	"LXXIV":    74,
	"LXXV":     75,
	"LXXVI":    76,
	"LXXVII":   77,
	"LXXVIII":  78,
	"LXXIX":    79,
	"LXXX":     80,
	"LXXXI":    81,
	"LXXXII":   82,
	"LXXXIII":  83,
	"LXXXIV":   84,
	"LXXXV":    85,
	"LXXXVI":   86,
	"LXXXVII":  87,
	"LXXXVIII": 88,
	"LXXXIX":   89,
	"XC":       90,
	"XCI":      91,
	"XCII":     92,
	"XCIII":    93,
	"XCIV":     94,
	"XCV":      95,
	"XCVI":     96,
	"XCVII":    97,
	"XCVIII":   98,
	"XCIX":     99,
	"C":        100,
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
	"XI",
	"XII",
	"XIII",
	"XIV",
	"XV",
	"XVI",
	"XVII",
	"XVIII",
	"XIX",
	"XX",
	"XXI",
	"XXII",
	"XXIII",
	"XXIV",
	"XXV",
	"XXVI",
	"XXVII",
	"XXVIII",
	"XXIX",
	"XXX",
	"XXXI",
	"XXXII",
	"XXXIII",
	"XXXIV",
	"XXXV",
	"XXXVI",
	"XXXVII",
	"XXXVIII",
	"XXXIX",
	"XL",
	"XLI",
	"XLII",
	"XLIII",
	"XLIV",
	"XLV",
	"XLVI",
	"XLVII",
	"XLVIII",
	"XLIX",
	"L",
	"LI",
	"LII",
	"LIII",
	"LIV",
	"LV",
	"LVI",
	"LVII",
	"LVIII",
	"LIX",
	"LX",
	"LXI",
	"LXII",
	"LXIII",
	"LXIV",
	"LXV",
	"LXVI",
	"LXVII",
	"LXVIII",
	"LXIX",
	"LXX",
	"LXXI",
	"LXXII",
	"LXXIII",
	"LXXIV",
	"LXXV",
	"LXXVI",
	"LXXVII",
	"LXXVIII",
	"LXXIX",
	"LXXX",
	"LXXXI",
	"LXXXII",
	"LXXXIII",
	"LXXXIV",
	"LXXXV",
	"LXXXVI",
	"LXXXVII",
	"LXXXVIII",
	"LXXXIX",
	"XC",
	"XCI",
	"XCII",
	"XCIII",
	"XCIV",
	"XCV",
	"XCVI",
	"XCVII",
	"XCVIII",
	"XCIX",
	"C",
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
	vals := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""

	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			num -= vals[i]
			roman += symbols[i]
		}
	}
	return roman
}
