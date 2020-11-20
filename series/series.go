package series

// All returns array with all substrings of s with length of n
func All(n int, s string) (r []string) {
	if len(s) < n {
		return r
	}

	for i := 0; i <= len(s)-n; i++ {
		r = append(r, s[i:i+n])
	}

	return r
}

// UnsafeFirst returns the first substring of s of the length n
func UnsafeFirst(n int, s string) (r string) {
	if len(s) < n {
		return ""
	}

	return s[0:n]
}

// First does exactly the same as UnsafeFirst, but also returns indicator of success
func First(n int, s string) (r string, ok bool) {
	r = UnsafeFirst(n, s)

	return r, r != ""
}
