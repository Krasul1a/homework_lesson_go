// На вхід подано стрінг з цілими числа, котри розділені пробілами.
// Треба повернути найбільше та найменше число.
// Наприклад:
// input := "1 2 3 4 5" // повертає "5 1"
// input := "1 9 3 4 -5" // повертає "9 -5"
// Уточнення:
// 1. Всі числа є не більше, ніж int32. Використовуйте цей тип даних.
// 2. В стрінгі завжди буде принаймні одне число.
// 3. Результатом має бути стрінг, в якому два числа розділені пробілом (або одне, якщо дано було лише одне
// число). Найбільше число має бути першим.
// 4. Якщо вам потрібні змінні чи константи - вони мають бути локальними, в межах функції main.
// func main(){
//     input := "1 9 3 4 -5"
//     var result string
//     //...
//     // тут має бути ваш код
//     // змінна result в кінці функції має тримати стрінг з правильними результатами, згідно до умови задачі
// }

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5" // дані які входят

	var sliceInt []int32
	var result string
	var tempInt int32 = 0

	inputSlice := strings.Split(input, " ") // Використання спліту, щоб зробити слайс з цифр, подивився тут https://www.educative.io/answers/how-to-split-a-string-in-golang

	for _, el := range inputSlice { // Цикл який переробляє стрінгові числа в int32 да додає в новий слайс int32
		temp, _ := strconv.Atoi(el)
		sliceInt = append(sliceInt, int32(temp))
	}

	for _, el := range sliceInt { // найбільший елемент
		if el > tempInt {
			tempInt = el
		}
	}

	result = strconv.Itoa(int(tempInt)) + " " // формування result (додав найбільший)

	for _, el := range sliceInt { // найменший елемент
		if el < tempInt {
			tempInt = el
		}

	}

	result += strconv.Itoa(int(tempInt)) // формування result (додав найменший)
	fmt.Println(result)
}
