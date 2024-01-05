package assert

// containsShallow does a shallow == comparison test to check if the slice contains the type.
func containsShallow[T comparable](s []T, v T) bool {
	for _, item := range s {
		if item == v {
			return true
		}
	}
	return false
}
