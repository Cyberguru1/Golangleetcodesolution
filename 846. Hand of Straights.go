// Time C. O(n^2), Space C. O(n)
// Runtime 73ms

func isNStraightHand(hand []int, groupSize int) bool {

	lent := len(hand)

	if (lent / groupSize) * groupSize != lent {
		return false
	}

	sort.Slice(hand, func(i, j int)bool{
		return hand[i] < hand[j]
	})

	hashmap := map[int]int{}

	for _, v := range hand {
		hashmap[v]++
	}

	for len(hand) > 0 {
		v := hand[0]
		for ii := 0; ii < groupSize; ii++ {
			if vv, ok := hashmap[v+ii]; !ok || vv <= 0 {
				return false
			}else {
				hashmap[v+ii]--
				for l:= 0; l < len(hand); l++ {
					if hand[l] == v+ii {
						hand = append(hand[:l], hand[l+1:]...)
						break
					}
				}
			}
		}
	}

	return true

}
