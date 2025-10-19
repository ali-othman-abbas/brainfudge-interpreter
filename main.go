package main

import (
	"fmt"
	"os"
)


func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("go run . [brainfudge source]")
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
	s := make([]int, 0, 15)
	for i, token := range input {
		switch token {
			case '[':
				s = append(s, i)
			case ']':
				if len(s) == 0 || input[s[len(s)-1]] == ']' {
					fmt.Println("Invalid Syntax")
					os.Exit(2)
				}
				LeftBracIdx := s[len(s)-1]
				s = s[:len(s) - 1]
				rightBracIdx := i
				res[LeftBracIdx] = rightBracIdx + 1
				res[rightBracIdx] = LeftBracIdx + 1
		}
	}
	
	if len(s) != 0{
		fmt.Println("Invalid Syntax")
		os.Exit(2)
	}
	
	return res
}