package circlequeue

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	MaxSize int
	Array   [5]int //数组
	Head    int    // 指向队列队首0
	Tail    int    // 指向队尾 0
}

//判断队列是否已满
func (this *CircleQueue) IsFull() bool {
	return (this.Tail+1)%this.MaxSize == this.Head
}

//判断队里是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.Tail == this.Head
}

//进队列
func (this *CircleQueue) AddQueue(val int) (err error) {
	if this.IsFull() {
		return errors.New("队列已满")
	}
	//把值给尾部
	this.Array[this.Tail] = val
	this.Tail = (this.Tail + 1) % this.MaxSize
	return
}

//出队列
func (this *CircleQueue) GetQueue() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("队列已空")
	}
	val = this.Array[this.Head]
	this.Head = (this.Head + 1) % this.MaxSize
	return
}

//显示队列元素
func (this *CircleQueue) ListQueue() {
	fmt.Println("队列情况如下：")
	//计算出队列多少元素
	//比较关键的一步
	size := (this.Tail + this.MaxSize - this.Head) % this.MaxSize
	if size == 0 {
		fmt.Println("队列已空")
	}
	//定义一个辅助变量 指向Head
	tempHead := this.Head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.Array[tempHead])
		tempHead = (tempHead + 1) % this.MaxSize
	}
	fmt.Println()
}

