// first solution, runtime 15ms
// Time C. O(n), Space C. O(1)

func maxProductDifference(nums []int) int {

	var m1, m2 int
	const MaxInt = int(^uint(0) >> 1)
	l1,l2 :=  MaxInt, MaxInt
	for _, v := range nums {
		if v > m1 {
			m2 = m1
			m1 = v
		} else if ( v > m2) {
			m2 = v
		}
		if v < l1 {
			l2 = l1
			l1 = v
		} else if ( v < l2){
			l2 = v
		}
	}
	return (m1 * m2) - (l1 * l2)
}