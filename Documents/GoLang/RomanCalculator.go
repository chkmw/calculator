package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	romanToInt()
	intToRoman()

	var str string

	fmt.Println("Введите выражение")
	fmt.Scan(&str)

	parts := strings.Split(str, " ")
	a, err := strconv.Atoi(parts[0])
	b, err1 := strconv.Atoi(parts[1])
	if err != nil && err1 != nil {
		fmt.Println("Ошибка преобразования:", err)
	} else {

		fmt.Println(a + b)
	}
}

func romanToInt() {

}

func intToRoman() {

}

func numsCheck() {

}
