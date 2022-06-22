// Прибрати всі дублікати з слайсу int.
// Приклад даних на вхід: [3, 4, 4, 3, 6, 3]
// видаляємо 3 по індексу 0
// видаляємо 4 по індексу 1
// видаляємо 3 по індексу 3
// Правильний результат: [3, 4, 6]
// Якщо вам потрібні змінні чи константи - вони мають бути локальними, в межах функції main.
// func main(){
//     arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
//     var result []int
//     //...
//     // тут має бути ваш код
//     // змінна result в кінці функції має тримати слайс з вже видаленими дублікатами відповідно до правил

package main

import "fmt"

func main() {
	arr := []int{3, 4, 4, 3, 6, 3} //{3(0), 4(1), 4(2), 3(3), 6(4), 3(5)}
	var result []int
	tempArr := []int{arr[0]}

	var number bool = true
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
		for j := 0; j < len(tempArr); j++ {
			if arr[i] == tempArr[j] {
				number = false
				fmt.Println(false)
				break
			}

		}
		if number {
			tempArr = append(tempArr, arr[i])
		}
		number = true
	}
	result = tempArr
	_ = result
	_ = number
	fmt.Println(result) // Вивести result

}
