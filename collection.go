package helper

func InSlice[T comparable](list []T, str T) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func MergeSlices[T any](vals ...[]T) (merged []T) {
	for _, val := range vals {
		if val == nil {
			continue
		}
		merged = append(merged, val...)
	}

	return merged
}
