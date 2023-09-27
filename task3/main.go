package main

import (
	"bytes"
	"fmt"

	"github.com/zeroflucs-given/generics/collections/stack"
)

type slice struct {
	tok []int
}

/**
 * Implement scanner for wrapping input into integer array
 * Try: 1 3 2 5 3 4
 */
func (s *slice) Scan(state fmt.ScanState, verb rune) error {
	tok, err := state.Token(false, func(r rune) bool { return r != '\n' })
	if err != nil {
		return err
	}
	if _, _, err := state.ReadRune(); err != nil {
		if len(tok) == 0 {
			panic(err)
		}
	}
	b := bytes.NewReader(tok)
	for {
		var d int
		_, err := fmt.Fscan(b, &d)
		if err != nil {
			break
		}
		s.tok = append(s.tok, d)
	}
	return nil
}

func getNearestLarger(in []int) []int {
	var size = len(in)

	var result = make([]int, size)
	var stack = stack.NewStack[int](size)

	for i := size - 1; i > 0; i-- {
		stack.Push(i)
		currentIndex := i - 1
		_, headIndex := stack.Peek()

		head := in[headIndex]
		current := in[currentIndex]
		fmt.Println("Current and head values:", current, head)

		if current == head {
			stack.Pop()
		}

		for stack.Count() > 0 && current > head {
			stack.Pop()
			_, headIndex = stack.Peek()
			head = in[headIndex]
		}

		if stack.Count() == 0 {
			result[currentIndex] = 0
		} else {
			result[currentIndex] = head
		}
	}

	return result
}

func main() {
	var s slice
	fmt.Println("Enter array of integers separated by space")
	fmt.Scan(&s)
	fmt.Println("Nearest larger on the right for each value:", getNearestLarger(s.tok))
}
