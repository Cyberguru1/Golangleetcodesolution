// second brute force solution
// Time C. O(n), Space C. O(1)


func minOperations(s string) int {

	var f, ff rune = '0','1'
	fcnt, scnt := 0, 0

	for i, v := range s {
		if (((i % 2 == 0) && (v != f)) || ((i % 2 != 0) && ( v == f))){
			fcnt++
		}
		if (((i % 2 == 0) && (v != ff)) || ((i % 2 != 0) && ( v == ff))){
			scnt++
		}
	}
	return min(fcnt, scnt)   
}