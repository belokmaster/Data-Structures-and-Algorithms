package main

import "fmt"

// Сложите две дроби a/b и c/d и выведите ответ в простейшей форме.

// 1. Найти общий знаменатель, найдя LCM (наименьшее общее кратное) из двух знаменателей.
// 2. Измените дроби, чтобы они имели одинаковый знаменатель, и добавьте оба члена.
// 3. Сведите полученную конечную дробь к ее более простой форме, разделив числитель
// и знаменатель на их наибольший общий множитель

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lowest(den, num *int) {
	commonFactor := gcd(*den, *num) // Вызываем gcd с разыменованными значениями
	*den = *den / commonFactor      // Изменяем значения по указателям
	*num = *num / commonFactor
}

// Функция для сложения двух дробей
func addFraction(num1, den1, num2, den2 int, num3, den3 *int) {
	// Вычисляем НОД знаменателей
	g := gcd(den1, den2)

	// Вычисляем общий знаменатель (НОК)
	*den3 = (den1 * den2) / g

	// Приводим дроби к общему знаменателю и складываем
	*num3 = (num1 * (*den3 / den1)) + (num2 * (*den3 / den2))

	// Сокращаем результат
	lowest(den3, num3)
}

func main() {
	var num1, den1 int
	var num2, den2 int
	fmt.Scan(&num1, &den1)
	fmt.Scan(&num2, &den2)
	var num3, den3 int

	addFraction(num1, den1, num2, den2, &num3, &den3)

	fmt.Printf("%d/%d + %d/%d = %d/%d\n", num1, den1, num2, den2, num3, den3)
}
