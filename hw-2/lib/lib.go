package lib

func Ternary[T any](cond bool, x, y T) T {
	if cond {
		return x
	}

	return y
}

func Filter[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(s))
	for i, v := range s {
		result[i] = f(v)
	}

	return result
}

func Reduce[T1, T2 any](s []T1, init T2, f func(T1, T2) T2) T2 {
	result := init
	for _, v := range s {
		result = f(v, result)
	}

	return result
}
