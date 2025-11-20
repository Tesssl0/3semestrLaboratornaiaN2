package main

// MatchPattern проверяет соответствие строки шаблону с поддержкой * и ?
func MatchPattern(str, pattern string) bool {
	n, m := len(str), len(pattern)
	i, j, star, match := 0, 0, -1, 0

	for i < n {
		if j < m && (pattern[j] == '?' || pattern[j] == str[i]) {
			i++
			j++
		} else if j < m && pattern[j] == '*' {
			star = j
			j++
			match = i
		} else if star != -1 {
			j = star + 1
			match++
			i = match
		} else {
			return false
		}
	}

	for j < m && pattern[j] == '*' {
		j++
	}
	return j == m
}
