package main

import (
	"fmt"
)

type Myqueue struct {
	queue []int
}

func constructor() *Myqueue {
	return &Myqueue{}
}

func (this *Myqueue) push(x int) {

	this.queue = append(this.queue, x)

}

func (this *Myqueue) pop() int {

	top := this.queue[0]

	this.queue = this.queue[1:]
	return top
}

func (this *Myqueue) top() int {

	return this.queue[0]
}

func (this *Myqueue) empty() bool {
	return len(this.queue) == 0
}

func (this *Myqueue) printQueue() {

	current := this.queue

	for _,val := range current {

		fmt.Println(val)
	}
}

func main() {

	l := constructor()

	l.push(10)
	l.push(20)
	l.push(30)
	l.push(40)
	l.push(50)
	fmt.Println(l.empty())
	l.printQueue()
}
