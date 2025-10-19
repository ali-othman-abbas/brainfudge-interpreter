package main

import (
	"fmt"
	"os"
)

type Stack[Type any] struct {
	arr []Type
}

func (this *Stack[Type]) Push(ele Type) {
	this.arr = append(this.arr, ele)
	
}

func (this *Stack[Type]) Pop() Type {
	if this.IsEmpty() {
		panic("stack is empty")
	}
	res := this.Peek()
	this.arr = this.arr[:len(this.arr) - 1]
	return res
}

func (this *Stack[Type]) Peek() Type {
	res := this.arr[len(this.arr) - 1]
	return res
}

func (this *Stack[Type]) Size() int {
	return len(this.arr)
}

func (this *Stack[Type]) IsEmpty() bool {
	return len(this.arr) == 0
}

func CreateStack[Type any]() Stack[Type] {
	return Stack[Type]{
		arr: make([]Type, 0, 15),
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("go run . [brainfudge string]")
		os.Exit(1)
	}
	
	runInterpreter(args[0])
}

func runInterpreter(input string) {
	jumpMap := initalizeJumpMap(input)
	tape := make([]byte, 2048)
	ptr := 0
	i := 0
	for i < len(input) {
		token := input[i]
		jump := false
		switch token {
			case '.':
				fmt.Printf("%c", tape[ptr])
			case '>':
				ptr++
				if ptr >= len(tape) {
					tape = append(tape, make([]byte, 2048)...)
				}
			case '<':
				ptr--
				if ptr < 0 {
					panic("pointer moved below tape limit")
				}
			case '+':
				tape[ptr]++
			case '-':
				tape[ptr]--
			case '[':
				if tape[ptr] == 0 {
					jump = true
				}
			case ']':
				if tape[ptr] != 0 {
					jump = true
				}
		}
		if jump {
			i = jumpMap[i]
		} else {
			i++
		}
	}
}

func initalizeJumpMap(input string) map[int]int {
	res := make(map[int]int)
	s := CreateStack[int]()
	for i, token := range input {
		switch token {
			case '[':
				s.Push(i)
			case ']':
				if s.IsEmpty() || input[s.Peek()] == ']' {
					fmt.Println("Invalid Syntax")
					os.Exit(2)
				}
				LeftBracIdx := s.Pop()
				rightBracIdx := i
				res[LeftBracIdx] = rightBracIdx + 1
				res[rightBracIdx] = LeftBracIdx + 1
		}
	}
	
	if !s.IsEmpty() {
		fmt.Println("Invalid Syntax")
		os.Exit(2)
	}
	
	return res
}