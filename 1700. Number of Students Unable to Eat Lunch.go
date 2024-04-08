// TIme C. O(n), Space C. O(1)
// BestRuntime 0ms

func countStudents(students []int, sandwiches []int) int {
	chg := 0
	for true {
		for i := 0; i < 1; i++ {
			if students[i] == sandwiches[i]{
				students, sandwiches = students[i+1:], sandwiches[i+1:]
				chg--
				break // aren't really necessary (<_-)
			}else{
				students = append(students[1:], students[0])
				chg++
				break
			}
		}
	if chg > len(students){
		return chg-1
	}else if len(students) == 0 {break}
	}
	return 0
}
