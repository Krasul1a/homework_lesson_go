package main

// Одне яблуко коштує 5.99 грн. Ціна однієї груші - 7 грн.
// Ми маємо 23 грн.
// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
// 2. Скільки груш ми можемо купити?
// 3. Скільки яблук ми можемо купити?
// 4. Чи ми можемо купити 2 груші та 2 яблука?
//
//
// Задача:
// Опишіть вирішення всіх пунктів задачі використовуючи необхідні змінні чи/та константи.
// Під опишіть, я маю на увазі наступне:
// Я маю збілдити ваш код і запустити програму. В результаті цього, я маю побачити роздруковані // відповіді на поставлені вище запитання. Перед відповідю треба роздрукувати саме питання.
// Правильно обирайте типи даних та назви змінних чи констант.
// Публікація:
// Створити папку в своєму репозиторії в github та завантажити туди main.go файл, в котрому буде зроблено дане завдання.

import (
	"fmt"
)

func main() {
	const Apple = 5.99
	const Pear = 7.0
	var Wallet float32 = 23
	var ShoppingCart float32 = 0

	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	ShoppingCart = Apple*9 + Pear*8
	fmt.Println(ShoppingCart, "грн")

	fmt.Println("Скільки груш ми можемо купити?")
	ShoppingCart = float32(int(Wallet / Pear))
	fmt.Println(ShoppingCart, "штук")

	fmt.Println("Скільки яблук ми можемо купити?")
	ShoppingCart = float32(int(Wallet / Apple))
	fmt.Println(ShoppingCart, "штук")

	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?")
	ShoppingCart = Apple*2 + Pear*2
	if Wallet >= ShoppingCart {
		fmt.Println("Так!")
	} else {
		fmt.Println("Ні!")
	}

}
