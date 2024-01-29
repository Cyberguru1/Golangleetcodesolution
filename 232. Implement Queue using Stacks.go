// Time C. O(1), Space C. O(n)
// runtime 1ms232. Implement Queue using Stacks


type MyQueue struct {
	arr []int
}

func Constructor() MyQueue {

	return MyQueue{
		arr: []int{},
	}

}

func (this *MyQueue) Push(x int) {
	this.arr = append(this.arr, x)

}

func (this *MyQueue) Pop() int {
	if len(this.arr) >= 1 {
		con := this.arr[0]
		this.arr = this.arr[1:]
		return con
	}
	return 0
}

func (this *MyQueue) Peek() int {
	if len(this.arr) >= 1 {
		return this.arr[0]
	}
	return 0

}

func (this *MyQueue) Empty() bool {
	if len(this.arr) >= 1 {
		return false
	}
	return true

}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */