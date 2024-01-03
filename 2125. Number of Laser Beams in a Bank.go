// first brute force solution
// Time C. O(n*m) where m is average length of string
// Space C. O(1)
// best runtime 24ms

func numberOfBeams(bank []string) int {
	result := 0

	var count func(s string) int

	count = func(s string) int {
		total := 0
		for _, i := range s {
			if i == '1' {
				total++
			}
		}
		return total
	}

	prev := count(bank[0])

	for i := 1; i < len(bank); i++ {
		d := count(bank[i])
		if d > 0 {
			result += (d * prev)
			prev = d
		}
	}

	return result

}