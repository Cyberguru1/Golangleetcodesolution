// first solution, Time C. O(n), Space C. O(1)
// best runtime 3ms 
func buyChoco(prices []int, money int) int {

	const MAX_INT = int(^uint(0) >> 1)

	first_min, second_min := MAX_INT, MAX_INT

	for _, v := range prices {
		if first_min > v {
			second_min = first_min
			first_min = v
		} else if (second_min > v){
			second_min = v
		}
	}

	if result := (money - first_min - second_min); result >= 0{
		return result
	} 

	return money

}