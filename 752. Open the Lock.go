// Time C. O(1), Space C. O(1) .. since all initialization of n is fixed as 100000
// runtime 9ms

type combination [4]uint16

func newCombination(n int) combination {
	return combination{
		uint16(n/1000) % 10,
		uint16(n/100) % 10,
		uint16(n/10) % 10,
		uint16(n) % 10,
	}
}

func (n *combination) int() uint16 {
	return n[0]*1000 + n[1]*100 + n[2]*10 + n[3]
}

func openLock(deadends []string, target string) int {
	var seen [10000]bool

	for _, dead := range deadends {
		n, _ := strconv.Atoi(dead)
		seen[n] = true
	}
	if seen[0] {
		return -1
	}

	// toVisit contains combinations reachable in "turn" turns
	// nextVisit contains the next round of combinations to check
	curCombs := []combination{{0, 0, 0, 0}}
	nextCombs := []combination{}
	seen[curCombs[0].int()] = true

	// Parse target combination
	targetInt, _ := strconv.Atoi(target)
	targetComb := newCombination(targetInt)

	// Visit all combinations in toVisit each round, adding new combs to nextVisit
	var turns int
	for len(curCombs) > 0 {
		nextCombs = nextCombs[:0]
		for _, cur := range curCombs {
			if cur == targetComb {
				return turns
			}
			// Note: +9 % 10 is the same as decrementing by 1
			for _, d := range [][4]uint16{
				{9, 0, 0, 0}, {1, 0, 0, 0}, {0, 9, 0, 0}, {0, 1, 0, 0},
				{0, 0, 9, 0}, {0, 0, 1, 0}, {0, 0, 0, 9}, {0, 0, 0, 1}} {
				next := cur // copy current combination
				for idx, v := range d {
					next[idx] += v
					next[idx] %= 10
				}
				if seen[next.int()] {
					continue
				}
				nextCombs = append(nextCombs, next)
				seen[next.int()] = true
			}
		}
		turns++
		curCombs, nextCombs = nextCombs, curCombs
	}

	return -1
}