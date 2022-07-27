package xutil

func InArray[T comparable](item T, Target []T) bool {
	for i := range Target {
		if item == Target[i] {
			return true
		}
	}

	return false
}

func InArrayAny[T comparable](items []T, Target []T) bool {
	for i := range items {
		if InArray(items[i], Target) {
			return true
		}
	}

	return false
}

func InArrayAll[T comparable](items []T, Target []T) bool {
	for i := range items {
		if !InArray(items[i], Target) {
			return false
		}
	}

	return true
}

func ArraySplit[T any](src []T, loopSize int) [][]T {
	res := make([][]T, 0, len(src)/loopSize+1)
	for start := 0; start < len(src); start += loopSize {
		end := start + loopSize
		if end > len(src) {
			end = len(src)
		}
		res = append(res, src[start:end])
	}

	return res
}

func ArrayHasDuplicate[T comparable](src []T) bool {
	m := make(map[T]struct{})
	for _, v := range src {
		if _, ok := m[v]; ok {
			return true
		}

		m[v] = struct{}{}
	}
	return false
}

func ArrayRemoveDuplicate[T comparable](src []T) []T {
	res := make([]T, 0)

	m := make(map[T]struct{})
	for _, v := range src {
		if _, ok := m[v]; ok {
			continue
		}

		res = append(res, v)
		m[v] = struct{}{}
	}

	return res
}
