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
	result = append(result, arr[0])
	for i := range arr {
		for j := range result {
			if i == j {
				fmt.Println(i)
				continue
			} else {
				result = append(result, j)
				fmt.Println(j)

			}

		}

	}
	fmt.Println(result)
	ListElementDelete := []int{1, 3, 5} // Список індексів які треба видалити

	for i := 0; i < len(ListElementDelete); i++ {
		j := ListElementDelete[i] - i // Додаю змішення, бо при видаленню елемента в наступній ітерації циклу індекси не відповідають старому слайсу

		copy(arr[j:], arr[j+1:]) //  Лінк де подивився як видаляти елемент в слайсі бо немає вбудованої функції https://yourbasic.org/golang/delete-element-slice/
		arr[len(arr)-1] = 0
		arr = arr[:len(arr)-1]

	}
	result = arr
	fmt.Println(result) // Вивести result

}
