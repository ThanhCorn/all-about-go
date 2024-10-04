package linq

import "golang.org/x/exp/constraints"

func Where[T any](arr []T, f func(T) bool) []T {
	var result []T
	for _, v := range arr {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func Map[T2, T1 any](arr []T1, f func(T1) T2) []T2 {
	var result []T2
	for _, v := range arr {
		selectedVal := f(v)
		result = append(result, selectedVal)
	}

	return result
}

func DistinctBy[T any, K comparable](arr []T, keyFunc func(T) K) []T {
	var result []T
	seen := make(map[K]bool)
	for _, v := range arr {
		key := keyFunc(v)
		if _, ok := seen[key]; !ok {
			seen[key] = true
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T any, U any](arr []T, f func(U, T) U, initial U) U {
	result := initial
	for _, v := range arr {
		result = f(result, v)
	}
	return result
}

func Any[T any](arr []T, f func(T) bool) bool {
	for _, v := range arr {
		if f(v) {
			return true
		}
	}
	return false
}

// SumBy sums a specific float field (float32 or float64)
func SumBy[T any, U constraints.Float](arr []T, amountFunc func(T) U) U {
	var total U
	for _, v := range arr {
		total += amountFunc(v)
	}
	return total
}
