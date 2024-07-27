package common

func PtrFrom[T any](val *T) T {
	if val == nil {
		var newVal T
		return newVal
	}

	return *val
}
