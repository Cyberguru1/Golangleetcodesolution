// Time C. O(n), Space C. O(1)

func minOperations(logs []string) (dist int ){
	sF, bF := "./", "../"

	for _, v := range logs {
		if v == sF {
			continue
		} else if v == bF && dist != 0 {
			dist--
		}else if v != bF{
			dist++
		}
	}

	return
}