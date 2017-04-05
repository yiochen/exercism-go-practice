package accumulate

const testVersion = 1

func Accumulate(given []string, converter func(string) string) []string {
	var results = make([]string, len(given))
	for index, value := range given {
		results[index] = converter(value)
	}
	return results
}
