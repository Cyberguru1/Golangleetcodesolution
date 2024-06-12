// Time C. O(n), Space C. O(1) -- because uses array of fixed size of 3
// Runtime 1ms

func sortColors(nums []int)  {
	hashmap := make([]int, 3)

	for _, k := range nums {
		hashmap[k]++
	}

	ini := 0
	for i := 0; i < 3; i++ {
		for j := ini; j < ini + hashmap[i]; j++ {
			nums[j] = i
		}
		ini += hashmap[i]
	}
}


// Rust implementation
// Runtime 0ms

// impl Solution {
// 	pub fn sort_colors(nums: &mut Vec<i32>) {
// 		let mut hashmap = vec![0; 3];

// 		for i in 0..nums.len() {
// 			hashmap[nums[i] as usize] += 1;
// 		}

// 		let mut ini = 0;
// 		let mut i = 0;

// 		loop {
// 			for j in ini..(hashmap[i as usize] + ini){
// 				nums[j] = i;
// 			}
// 			ini += hashmap[i as usize];
// 			i += 1;
// 			if ( i >= 3 ){
// 				break;
// 			}
// 		}
// 	}
// }