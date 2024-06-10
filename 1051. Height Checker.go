// Time C. O(NlogN), Space C. O(N)
// Runtime 2ms

func heightChecker(heights []int) (count int) {

	arr2 := append([]int{}, heights...)

	sort.Slice(arr2, func(i, j int)bool {
		return arr2[i] < arr2[j]
	})

	for i := 0; i < len(heights); i++ {
		if heights[i] != arr2[i] {
			count++
		}
	}
	return
}

// Rust impl

// impl Solution {
// 	pub fn height_checker(heights: Vec<i32>) -> i32 {
// 		let mut b = heights.clone();
// 		b.sort();
// 		let mut lent:i8 = heights.len().try_into().unwrap();
// 		lent = lent - 1;
// 		let mut count = 0;

// 		while (lent >= 0) {
// 			if heights[lent as usize] != b[lent as usize] {
// 				count += 1;
// 			}
// 			lent -= 1;
// 		}
// 		count
// 	}
// }

