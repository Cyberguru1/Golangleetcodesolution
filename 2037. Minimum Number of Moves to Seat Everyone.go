// Time C. O(nlogn), Space C. O(1)
// Runtime 6ms

func minMovesToSeat(seats []int, students []int) (count int ){
	sort.Ints(seats)
	sort.Ints(students)
	for i := 0; i < len(seats); i++ {
		count += abs(seats[i] - students[i])
	}
	return
}

func abs(v int) int {
	if v < 0 {
		return (v * -1)
	}
	return v
}


// Rust implementation
// runtime 2ms

// impl Solution {
// 	pub fn min_moves_to_seat(seats: Vec<i32>, students: Vec<i32>) -> i32 {

// 		let mut sort1 = seats.clone();
// 		sort1.sort();
// 		let mut sort2 = students.clone();
// 		sort2.sort();

// 		let mut count = 0;

// 		for i in 0..sort1.len() {
// 			count += (sort2[i as usize] - sort1[i as usize]).abs();
// 		}
// 		count

// 	}
// }