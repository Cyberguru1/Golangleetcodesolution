// Time C. O(n), Space C. O(1)
// Runtime 2ms

func singleNumber(nums []int) []int {

	xorAll := 0
	for _, num := range nums {
		xorAll ^= num
	}

	setBit := xorAll & -xorAll
	a, b := 0, 0

	for _, num := range nums {
		if num & setBit != 0 {
			a ^= num
		} else {
			b ^= num
		}

	}

	return []int{a, b}
}
