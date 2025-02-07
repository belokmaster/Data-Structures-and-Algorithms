package main

/*
Stein’s algorithm or binary GCD algorithm is an algorithm that computes the greatest common divisor
of two non-negative integers.
Stein’s algorithm replaces division with arithmetic shifts, comparisons, and subtraction.
*/

// он же бинарный алгоритм евклида, использующий свойства нода
// Главное отличие от классического алгоритма Евклида — отсутствие деления и использование битовых операций.

// Time Complexity: O(N*N)
// Auxiliary Space: O(1)
// не смотря на плохую ассимптотику, работает быстрее обычного евклида.
func binaryGCD(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	// Находим shift — наибольшую степень 2, на которую делятся a и b
	shift := 0
	for (a|b)&1 == 0 {
		a >>= 1
		b >>= 1
		shift++
	}

	// Делаем a нечетным
	for a&1 == 0 {
		a >>= 1
	}

	// Теперь a — нечетное
	for b != 0 {
		// Делаем b нечетным
		for b&1 == 0 {
			b >>= 1
		}

		// Теперь a и b нечетные, меняем местами, если нужно
		if a > b {
			a, b = b, a
		}

		b -= a // b теперь чётное
	}

	// Восстанавливаем общий множитель 2^shift
	return a << shift
}
