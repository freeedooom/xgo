package xutil

func IsZeroValue[T comparable](src T) bool {
	return src == *new(T)
}

func If[T any](condition bool, valueIfTrue T, valueIfFalse T) T {
	if condition {
		return valueIfTrue
	}
	
	return valueIfFalse
}

func IfNullThen[T comparable](data, defaultData T) T {
	if IsZeroValue(data) {
		return defaultData
	}

	return data
}

func TypeV[T any](src *T) T {
	if src == nil {
		src = new(T)
	}
	return *src
}

func TypeVWithError[T any](src *T, err error) (T, error) {
	return TypeV(src), err
}

func TypeP[T any](src T) *T {
	return &src
}

func TypePWithError[T any](src T, err error) (*T, error) {
	return TypeP(src), err
}

func Must[T any](data T, err error) T {
	if err != nil {
		panic(err)
	}

	return data
}

func IgnoreError[T any](data T, err error) T {
	return data
}
