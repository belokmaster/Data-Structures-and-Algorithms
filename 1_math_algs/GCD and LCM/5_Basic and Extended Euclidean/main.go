package main

// Простой способ найти НОД — разложить оба числа на множители и умножить их на общие простые множители.
// Расширенный алгоритм Евклида также находит целочисленные коэффициенты x и y, такие что: ax + by = gcd(a, b)

//Time Complexity: O(log N)
//Auxiliary Space: O(log N)
func extendedGCD(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}

	gcd, x1, y1 := extendedGCD(b%a, a)
	x := y1 - (b/a)*x1
	y := x1

	return gcd, x, y
}

/*
если D(a, b) == 1 => a и b - взаимно простые. => x - обратное по модулю для a
и y обратное по модулю для b. Нужно для RSA-шифрования
*/
