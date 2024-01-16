// Time C. O(1), Space C. O(n)

import "math/rand"

type RandomizedSet struct {
	container map[int]int
	keys      []int
}

func Constructor() RandomizedSet {

	return RandomizedSet{
		container: map[int]int{},
		keys:      []int{},
	}

}

func (this *RandomizedSet) Insert(val int) bool {

	if _, ok := this.container[val]; ok {
		return false
	}

	this.container[val] = len(this.keys)
	this.keys = append(this.keys, val)

	return true

}

func (this *RandomizedSet) Remove(val int) bool {

	if ind, ok := this.container[val]; ok {

		this.keys = append(this.keys[:ind], this.keys[ind+1:]...)
		delete(this.container, val)

		for v, k := range this.container {
			if k > ind {
				this.container[v] = k - 1
			}
		}
		return true
	}
	return false

}

func (this *RandomizedSet) GetRandom() int {

	return this.keys[rand.Intn(len(this.keys))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */



// second solution

import "math/rand"

type RandomizedSet struct {
	container map[int]bool          
}



func Constructor() RandomizedSet {

	return RandomizedSet{
		container: map[int]bool{},
	}

}


func (this *RandomizedSet) Insert(val int) bool {

	if _, ok := this.container[val]; ok{
		return false
	}

	this.container[val] = true    
	return true

}


func (this *RandomizedSet) Remove(val int) bool {


	if _, ok := this.container[val]; ok{
		delete(this.container, val)
		return true
	}

	return false

}


func (this *RandomizedSet) GetRandom() int {

	lenn := len(this.container)
	randVal := rand.Intn(lenn)

	count := 0

	for k, _ := range this.container {
		if count == randVal {
			return k
		}
		count++
	}

	return 0
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */