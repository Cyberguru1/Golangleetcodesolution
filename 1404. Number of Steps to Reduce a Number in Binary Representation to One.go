//Time C. O(n), Space C. O(1)
// Runtime 3ms, Memory 3.09mbs

func numSteps(s string) (count int ){

	val, _:= new(big.Int).SetString(s, 2)

	for int32(val.Int64()) != 1 {
		res := new(big.Int).Mod(val, big.NewInt(2))
		if (int32(res.Int64()) == 0) {
			val = val.Div(val, big.NewInt(2))
		}else{
			val = val.Add(val, big.NewInt(1))
		}
		count++
	}
	return
}
