// Time C. O(n) Space C. O(n)
/// runtime 3ms

import "strings"

type cont struct {
	char  rune
	prior int
}

func frequencySort(s string) (res string) {

	smapp := map[rune]int{}

	arr := []cont{}

	for _, k := range s {
		smapp[k]++
	}

	for c, v := range smapp {
		arr = append(arr, cont{
			char:  c,
			prior: v,
		})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].prior < arr[j].prior
	})

	fmt.Println(arr)

	for i := len(arr) - 1; i > -1; i-- {
		res = fmt.Sprintf("%s%s", res, strings.Repeat(string(arr[i].char), arr[i].prior))
	}

	return

}