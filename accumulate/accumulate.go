package accumulate

// Accumulate performs op operation on each element of the input array
// and returns that as a result
func Accumulate(in []string, op func(string) string) []string {
	for i, v := range in {
		in[i] = op(v)
	}

	return in
}
