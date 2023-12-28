// Time C. O(n) Space C. O(1)
// runtime 5ms


func compress(chars []byte) int {
	count := 1
	total := 0
	s    := 0
	if len(chars) == 1 {
		return 1
	}
	for i:= 1; i < len(chars); i++ {
		if chars[i] == chars[i-1] {
			count++
		}
		if ((chars[i] != chars[i-1]) || (i == (len(chars) - 1))){

			if (count == 1){
				total += 1
				chars[s] = chars[i-1]
				s++

				if (i == (len(chars) - 1)){
					total += 1
					chars[s] = chars[i]
					s++
			}
			} else if (count != 1){
				total += 1
				chars[s] = chars[i-1]
				s++
				lc := chars[i-1]
				total += len(fmt.Sprint(count))
				for i, v := range fmt.Sprint(count){
					chars[s + i] = byte(v)
				}
				s += len(fmt.Sprint(count))
				count = 1

				if ((i == (len(chars) - 1)) && ( chars[i] != lc) && s < len(chars)){
					total += 1
					chars[s] = chars[i]
					s++
			}
			}
			count = 1
		}
	}
	return total
}
