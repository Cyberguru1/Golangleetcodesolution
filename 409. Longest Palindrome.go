// Time C. O(n), Space C. O(n)
// Runtime 2ms
func longestPalindrome(s string) int {

	hashmap := map[rune]int{}

	for _, k := range s {
		hashmap[k]++
	}

	fmt.Println(hashmap)

	count, bOdd  := 0, 0

	for _, v := range hashmap {
		if v % 2 == 0 {
			count += v
		}else {
			count += (v - 1)
			bOdd = 1
		}
	}

	return count+bOdd
}

// tried using rust

// impl Solution {
// 	pub fn longest_palindrome(s: String) -> i32 {
// 		let mut array: [i32; 58] = [0;58];
// 		for i in s.bytes() {
// 			array[<u8 as Into<usize>>::into(i - 65)] += 1;
// 		}

// 		let (mut count, mut odd) = (0, 0);

// 		for val in array.iter(){
// 			if val % 2 == 0 {
// 				count += val;
// 			}else {
// 				count += (val - 1);
// 				odd = 1;
// 			}
// 		}
// 		count+odd   
// 	}
// }
