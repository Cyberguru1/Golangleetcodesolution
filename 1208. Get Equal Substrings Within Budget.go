// Time C. O(nlogn), Space C. O(n)

func equalSubstring(s string, t string, maxCost int) (ress int) {

	arr := []int{}
	for j := 0; j < len(s); j++ {
		diff := rune(s[j]) - rune(t[j])
		if diff < 0 {
			diff = diff * -1
		}
		arr = append(arr, int(diff))
	}

	st, end, summ := 0, 0, 0

	for end < len(s){
		summ += arr[end]
		if (summ > maxCost){
			for summ > maxCost {
				summ -= arr[st]
				st++
			}
		}
		ress = max(ress, end-st+1)
		end++
	}
	return
}