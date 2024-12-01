package util

func Valid(name string) bool {
	for _, v := range name {
		if v > 0 && v <= 32 {
			return false
		}
	}
	return true
}
