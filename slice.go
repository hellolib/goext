package goext

// Filter 切片过滤器
func Filter[T any](sourceSlice []T, fn func(index int, item T) bool) []T {
	resultSlice := make([]T, 0, len(sourceSlice))
	for index, item := range sourceSlice {
		if fn(index, item) {
			resultSlice = append(resultSlice, item)
		}
	}
	return resultSlice
}

// Map 切片元素加工(返回切片元素类型可以不同)
func Map[T any, R any](sourceSlice []T, fn func(index int, item T) R) []R {
	resultSlice := make([]R, len(sourceSlice))
	for index, item := range sourceSlice {
		resultSlice[index] = fn(index, item)
	}
	return resultSlice
}

// Reduce 对切片的所有元素聚合
func Reduce[T any, R any](sourceSlice []T, fn func(index int, item T, agg R) R, initial R) R {
	for index, item := range sourceSlice {
		initial = fn(index, item, initial)
	}

	return initial
}
