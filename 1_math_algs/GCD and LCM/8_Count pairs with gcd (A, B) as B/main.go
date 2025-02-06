package main

// Дано число n. Нам нужно найти количество упорядоченных пар a и b, таких, что gcd(a, b) равен самому b.

/*
Наивный подход: gcd(a, b) = b означает, что b является множителем a.
Таким образом, общее количество пар будет равно сумме делителей для каждого a = 1 до n.

Эффективный подход: gcd(a, b) = b означает, что a является кратным b.
Вместо перебора всех i достаточно учитывать, что floor(n / i) принимает
не более 2 * sqrt(n) различных значений.
Для i ≤ sqrt(n) значения floor(n / i) идут по порядку,
а для больших i — повторяются, поэтому можно сразу посчитать их сумму группами,
избегая лишних операций. Это сокращает сложность с O(n) до O(sqrt(n)).
*/

// Time complexity: O(n). This is because the while loop takes O(n) time to complete since it is looping over all elements of the array.
// Auxiliary space: O(1), as no extra space is used.
func countPairs(n int) int {
	k := n
	imin := 1 // Минимальное значение i в текущем диапазоне
	ans := 0

	for imin <= n {
		// Максимальное i, для которого floor(n / i) = k
		imax := n / k

		// adding k*(number of i with
		// floor(n/i) = k to ans
		ans += k * (imax - imin + 1)

		imin = imax + 1 // Переход к следующему диапазону значений
		k = n / imin    // Новое значение k = floor(n / imin)
	}

	return ans
}
