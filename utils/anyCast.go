package utils

func CastToAnySlice[T any](inputSlice []T) (outputSlice []any) {
	outputSlice = make([]any, len(inputSlice))
	for i, item := range inputSlice {
		outputSlice[i] = any(item)
	}
	return outputSlice
}

func CastToIntSlice(inputSlice []any) (outputSlice []int) {
	outputSlice = make([]int, len(inputSlice))
	for i, item := range inputSlice {
		outputSlice[i] = item.(int)
	}
	return outputSlice
}

func CastToFloat64Slice(inputSlice []any) (outputSlice []float64) {
	outputSlice = make([]float64, len(inputSlice))
	for i, item := range inputSlice {
		outputSlice[i] = item.(float64)
	}
	return outputSlice
}

func CastToFloat32Slice(inputSlice []any) (outputSlice []float32) {
	outputSlice = make([]float32, len(inputSlice))
	for i, item := range inputSlice {
		outputSlice[i] = item.(float32)
	}
	return outputSlice
}

func CastToInt32Slice(inputSlice []any) (outputSlice []int32) {
	outputSlice = make([]int32, len(inputSlice))
	for i, item := range inputSlice {
		outputSlice[i] = item.(int32)
	}
	return outputSlice
}

func CastToInt64Slice(inputSlice []any) (outputSlice []int64) {
	outputSlice = make([]int64, len(inputSlice))
	for i, item := range inputSlice {
		outputSlice[i] = item.(int64)
	}
	return outputSlice
}

func CastToStringSlice(inputSlice []any) (outputSlice []string) {
	outputSlice = make([]string, len(inputSlice))
	for i, item := range inputSlice {
		outputSlice[i] = item.(string)
	}
	return outputSlice
}
