// third solution time: (O(n)), space (O(n)); runtime: 0ms
type void struct{}
var member void

func destCity(paths [][]string) string {
	lpaths := len(paths)
	arr := make(map[string]void, lpaths)

	for _, k := range paths {
		arr[k[0]] = member
	}

	for _, k := range paths {
		if _, ok := arr[k[1]]; !ok {
			return k[1]
		}
	}

	return "None"
}
