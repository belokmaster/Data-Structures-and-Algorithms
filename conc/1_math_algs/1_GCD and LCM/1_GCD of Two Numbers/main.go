package main

import (
	"fmt"
	"time"
)

func benchmarkAll(a, b int) {
	measureTime("Naive approach", gcdNativeApproach, a, b)
	measureTime("Euclidean Algorithm", gcdEuclideanAlgorithm, a, b)
	measureTime("Optimized by checking divisibility", gcdEuclideanAlgorithmOptimizationCheckingDivisibility, a, b)
	measureTime("Optimized using division", gcdEuclideanAlgOptimizationUsingDivision, a, b)
	measureTime("Iterative Euclidean Algorithm", gcdEuclideanAlg, a, b)
	fmt.Println("")
}

func measureTime(name string, fn func(int, int) int, a, b int) {
	start := time.Now()
	result := fn(a, b)
	duration := time.Since(start)

	fmt.Printf("%s: gcd(%d, %d) = %d, took %v\n", name, a, b, result, duration)
}

func main() {
	for {
		var input string
		fmt.Println("Enter the command: ")
		fmt.Println("q - quit from the program, else - calculate GCD of two numbers")
		fmt.Scan(&input)

		if input == "q" {
			fmt.Println("Exit from the program.")
			break
		}

		var a, b int
		fmt.Println("Enter two numbers: a b")
		fmt.Scan(&a, &b)

		fmt.Println("Specify which algorithm you want to implement?")
		fmt.Println("1 - naive approach of calc gcd")
		fmt.Println("2 - usual Euclidean Algorithm")
		fmt.Println("3 - usual Euclidean Algorithm optimazed by checking divisibility")
		fmt.Println("4 - usual Euclidean Algorithm optimazed by using division")
		fmt.Println("5 - usual iterative Euclidean Algorithm")
		fmt.Println("t - test all algs on big digits")
		fmt.Println("q - quit from the program")
		fmt.Scan(&input)

		var result int
		switch input {
		case "1":
			result = gcdNativeApproach(a, b)
		case "2":
			result = gcdEuclideanAlgorithm(a, b)
		case "3":
			result = gcdEuclideanAlgorithmOptimizationCheckingDivisibility(a, b)
		case "4":
			result = gcdEuclideanAlgOptimizationUsingDivision(a, b)
		case "5":
			result = gcdEuclideanAlg(a, b)
		case "t":
			benchmarkAll(a, b)
			continue
		case "q":
			return
		default:
			fmt.Println("Invalid choice, using Euclidean Algorithm by default.")
			result = gcdEuclideanAlgorithm(a, b)
		}

		fmt.Printf("GCD of %d and %d is %d.\n", a, b, result)
	}
}
