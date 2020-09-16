package main

/*
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type CQueue struct {
	stackA []int
	stackB []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.stackA = append(this.stackA, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stackB) == 0 {
		if len(this.stackA) == 0 {
			return -1
		}
		for len(this.stackA) > 0 {
			index := len(this.stackA) - 1
			value := this.stackA[index]
			this.stackB = append(this.stackB, value)
			this.stackA = this.stackA[:index]
		}
	}
	index := len(this.stackB) - 1
	value := this.stackB[index]
	this.stackB = this.stackB[:index]
	return value
}

func main() {}
