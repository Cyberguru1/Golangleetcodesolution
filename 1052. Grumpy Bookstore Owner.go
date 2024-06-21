// Time C. O(nm), Space C. O(1)

func maxSatisfied(customers []int, grumpy []int, minutes int)(res int) {     

	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			res += customers[i]
		}
	}

	old_res := res
	for i := 0; i < len(customers); i++ {
		curr := old_res
		for j := i; j < len(customers) && j < i + minutes; j++ {
			if grumpy[j] == 1 {
				curr += customers[j]
			}
		} 
		res = max(curr, res)
	}

	return
}