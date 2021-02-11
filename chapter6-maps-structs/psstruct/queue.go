/*
Write IntQueue type which has push and pop methods as following

`enqueue(val int)` -> enqueue into queue
`dequeue() (int, bool)` -> dequeue first element from queue, false if empty

Pseudocode:

Initial queue
{}

enqueue(1)
{1}

enqueue(2)
{1, 2}

printl(dequeue()) -> 1
printl(dequeue()) -> 2
printl(queue) -> {}
*/
package main

import "fmt"

type IntQueue struct {
	list []int
}

func NewIntQueue() *IntQueue {
	return &IntQueue{make([]int, 0, 0)}
}

func (is IntQueue) isEmpty() bool {
	return is.list == nil || len(is.list) == 0
}

func (is *IntQueue) enqueue(val int) {
	is.list = append(is.list, val)
}

func (is *IntQueue) dequeue() (int, bool) {
	if is.isEmpty() {
		return 0, false
	} else {
		list := is.list
		is.list = list[1:]
		return list[0], true
	}
}

func (is IntQueue) peek() []int {
	return is.list
}

func main() {

	queue := NewIntQueue()

	fmt.Println("IsEmpty: ", queue.isEmpty())

	queue.enqueue(1)
	fmt.Println("Peek: ", queue.peek())

	queue.enqueue(2)
	fmt.Println("Peek: ", queue.peek())

	fmt.Println("IsEmpty: ", queue.isEmpty())

	popped, _ := queue.dequeue()
	fmt.Println("Dequeued:", popped, " | Peek: ", queue.peek())

	popped, _ = queue.dequeue()
	fmt.Println("Dequeued:", popped, " | Peek: ", queue.peek())

	fmt.Println("IsEmpty: ", queue.isEmpty())

	list := queue.peek()
	list = append(list, 5)
	fmt.Println("Copy List: ", list)
	queue.enqueue(20)

	fmt.Println("Peek: ", queue.peek())
	fmt.Println("Copy List: ", list)
}
