// Time C. O(NlogN), Space C. O(n)
// Runtime 2ms

func relativeSortArray(arr1 []int, arr2 []int) (new_arr []int) {

	hashmap := make(map[int]int)

	for _, val := range arr1 {
		hashmap[val]++
	}

	for _, val := range arr2 {
		count := hashmap[val]
		for i := 0; i < count; i++{
			new_arr = append(new_arr, val)
			hashmap[val]--
		}
	}

	last_part := make([]int, 0)

	for val, c := range hashmap {
		if c > 0 {
			for i := 0; i < c; i++{
				last_part = append(last_part, val)
			}
		}
	}
	
	sort.Slice(last_part, func(i, j int)bool {
		return last_part[i] < last_part[j]
	})

	new_arr = append(new_arr, last_part...)

	return
}

// Rust implementation
// use std::collections::HashMap;
// impl Solution {
// 	pub fn relative_sort_array(arr1: Vec<i32>, arr2: Vec<i32>) -> Vec<i32> {
// 		let mut map: HashMap<i32, i32> = HashMap::new();
// 		let mut new_arr = vec![];
// 		let mut new_arr2 = vec![];

// 		for i in 0..arr1.len() {
// 			let mut val = map.get(&arr1[i]);
// 			let mut dval = 1;
// 			if val != None {
// 				dval = *val.unwrap() + 1;
// 			}
// 			map.insert(arr1[i], dval);
// 		}

// 		for k in 0..arr2.len() {
// 			let mut count = *map.get(&arr2[k]).unwrap();
// 			while (count > 0) {
// 				new_arr.push(arr2[k]);
// 				count -= 1;
// 			}
// 			map.insert(arr2[k], count);
// 		}
// 		for (val, count) in map {
// 			if count > 0 {
// 				for j in 0..count {
// 					new_arr2.push(val);
// 				}
// 			}
// 		}
// 		new_arr2.sort();
// 		[new_arr, new_arr2].concat()
// 	}
// }