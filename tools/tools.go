package tools

func ReplaceAll[T comparable](a []T, oldValue T, newValue T) {
	for i, v := range a {
		if v == oldValue {
			a[i] = newValue
		}
	}
}
