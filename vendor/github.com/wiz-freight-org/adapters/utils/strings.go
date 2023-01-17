package utils

func ContainsInStringArray(a []string, v string) bool {
	contains := false
	for _, str := range a {
		if str == v {
			contains = true
			break
		}
	}

	return contains
}

func AppendWithoutDuplicates(a []string, s string) []string {
	m := make(map[string]bool)
	for _, v := range a {
		m[v] = true
	}

	if _, ok := m[s]; !ok {
		a = append(a, s)
	}

	return a
}

func AppendWithoutDuplicatesArray(a []string, strs []string) []string {
	m := make(map[string]bool)
	for _, v := range a {
		m[v] = true
	}

	for _, s := range strs {
		if _, ok := m[s]; !ok {
			a = append(a, s)
		}
	}

	return a
}
