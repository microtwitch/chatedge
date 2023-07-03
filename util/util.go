package util

func Contains(l []string, e string) bool {
	for _, v := range l {
		if v == e {
			return true
		}
	}

	return false
}
