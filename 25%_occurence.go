package main

func findSpecialInteger(arr []int) int {
	threshold := int(len(arr) / 4)

	if len(arr) == 2 {
		return arr[1]
	}

	for i := range arr {
		if arr[i] == arr[i + threshold]{
			return arr[i]
		}
	}
	return -1 
}

func findSpecialInteger(arr []int) int {
    out := make(map[int]int)

    for _, v := range arr {
        if _, ok := out[v]; ok {
            out[v] += 1
        } else {
            out[v] = 1
        }

    }
    var Maxx int = 1
    var Val int = 1

    for i, v := range out {

        if v > Maxx {
            Maxx = v
            Val = i
        }
    }

    return Val
}