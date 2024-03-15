// runtime 21ms
// Time C. O(n), Space C. O(n)
// don't mind my soln, i no it aint the best of the best.............would do better next time though..

func productExceptSelf(nums []int) (res []int) {
	lent := len(nums) - 1

	F, R := make([]int, lent+1), make([]int, lent+1)

	for i := 0; i <= lent; i++ {
		F[i] = 1
		R[i] = 1
	}

	F[lent], R[0] = 1, 1

	j := lent - 1

	for i := 1; i < lent+1 && j >= 0; i++ {
		F[i] = nums[i-1] * F[i-1]
		R[j] = nums[j+1] * R[j+1]
		j--
	}

	for i := 0; i <= lent; i++ {
		res = append(res, F[i]*R[i])
	}

	return
}