// Second bruteforce soultion
// Time C. O(N), Space C. O(N)
// Best runtime 0ms, beats 100%, Space 2.22mb


// Pair represents a pair of values.
type Pair struct {
	First  int
	Second int
}

type void struct {}

var val void

func isPathCrossing(path string) bool {
	d := map[byte]int{'N':1, 'S':-1, 'E':1, 'W':-1}
	p := Pair{
		First: 0,
		Second: 0,
	}
	mapCheck := map[Pair]void{p:val}

	for _, s := range path {
		ss := byte(s)
		if (ss == 'N') || (ss == 'S') {
			p.First += d[ss]

			if _, ok := mapCheck[p]; ok {
				return true
			} else {
				mapCheck[p] = val
			}
		} else {
			p.Second += d[ss]

			if _, ok := mapCheck[p]; ok {
				return true
			} else {
				mapCheck[p] = val
			}
		}
	}
	return false
}



