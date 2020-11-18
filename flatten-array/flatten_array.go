package flatten

// Flatten takes nested array of integers and nils and flattens it skipping all nils
func Flatten(input interface{}) []interface{} {

	switch input.(type) {
	case []interface{}:

		r := []interface{}{}

		for _, v := range input.([]interface{}) {
			r = append(r, Flatten(v)...)
		}

		return r
	case int:
		return []interface{}{input.(int)}
	}

	return []interface{}{}
}
