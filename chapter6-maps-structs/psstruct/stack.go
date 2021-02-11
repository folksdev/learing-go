/*
Write IntStack type which has push and pop methods as following

`push(val int)` -> push into stack
`pop() (int, bool)` -> pop first element from stack, false if empty

Pseudocode:

Initial stack
{}

push(1)
{1}

push(2)
{2, 1}

printl(pop()) -> 2
printl(pop()) -> 1
printl(stack) -> {}
*/
package main

import "fmt"

type IntStack struct {
	list []int
}

func NewIntStack() *IntStack {
	return &IntStack{make([]int, 0, 0)}
}

func (is IntStack) isEmpty() bool {
	return is.list == nil || len(is.list) == 0
}

func (is *IntStack) push(val int) {
	head := []int{val}
	is.list = append(head, is.list...)
}

func (is *IntStack) pop() (int, bool) {
	if is.isEmpty() {
		return 0, false
	} else {
		list := is.list
		is.list = list[1:]
		return list[0], true
	}
}

func (is *IntStack) peek() []int {
	return is.list
}

func main() {

	stack := NewIntStack()

	fmt.Println("IsEmpty: ", stack.isEmpty())

	stack.push(1)
	fmt.Println("Peek: ", stack.peek())

	stack.push(2)
	fmt.Println("Peek: ", stack.peek())

	fmt.Println("IsEmpty: ", stack.isEmpty())

	popped, _ := stack.pop()
	fmt.Println("Popped:", popped, " | Peek: ", stack.peek())

	popped, _ = stack.pop()
	fmt.Println("Popped:", popped, " | Peek: ", stack.peek())

	fmt.Println("IsEmpty: ", stack.isEmpty())

	list := stack.peek()
	list = append(list, 5)

	fmt.Println("List: ", list)
	fmt.Println("Stack List: ", stack.peek())
}
