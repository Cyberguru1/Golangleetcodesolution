// Time C. O(n), Space C. O(n)
// runtime 107ms

func countSubarrays(nums []int, kk int) (res int64) {

	mV := 0
	mVs := []int{}

	for i, k := range nums {
		if k > mV {
			mV, res, mVs = k, 0, []int{}
		}

		if k == mV {
			mVs = append(mVs, i)
		}

		if len(mVs) >= kk {
			res += int64(mVs[len(mVs)-kk]) + 1
		}
	}

	return

}
