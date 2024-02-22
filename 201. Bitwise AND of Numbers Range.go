// Time C. O(32) -> O(1), Space c. O(1)
// runtime 15ms

func rangeBitwiseAnd(left int, right int) int {

	leftstr := strconv.FormatInt(int64(left), 2)
	rightstr := strconv.FormatInt(int64(right), 2)

	leftstr = fmt.Sprintf("%032s", leftstr)
	rightstr = fmt.Sprintf("%032s", rightstr)

	d := 0

	for i := 0; i < len(leftstr); i++ {
		if leftstr[i] != rightstr[i] {
			break
		}
		d = i
	}

	res := leftstr[:d+1] + strings.Repeat("0", 32-d-1)
	out, _ := strconv.ParseInt(res, 2, 32)
	return int(out)
}

// more concise solution
// Time C. O(1), Space C. O(1)
// Runtime 3ms

func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right = right & (right - 1)
	}
	return right
}