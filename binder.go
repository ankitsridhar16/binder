package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Type def for user input
type InputBuffer struct {
	buffer       []byte
	bufferLength int
	inputLength  int
}

func NewInputBuffer() *InputBuffer {
	return &InputBuffer{} // assign nil/zero values to `InputBuffer` fields
}

// Take user input from stdin and initialize input buffer
func (ib *InputBuffer) ReadInput() error {
	// take input from os stdin and prepare reader
	reader := bufio.NewReader(os.Stdin)
	line, lineErr := reader.ReadBytes('\n')
	if lineErr != nil {
		return lineErr
	}

	// initialize values from line input -> input buffer
	ib.buffer = line
	ib.inputLength = len(line) - 1 // remove newline
	ib.bufferLength = cap(line)

	return nil
}

// Print DB prompt in terminal after each line
func printDBprompt() {
	fmt.Print("binderDB > ")
}
func main() {
	inputBuffer := NewInputBuffer()

	for {
		printDBprompt()
		inputErr := inputBuffer.ReadInput()
		if inputErr != nil {
			fmt.Println("Error reading input: ", inputErr)
			os.Exit(1)
		}

		input := strings.TrimSpace(string(inputBuffer.buffer[:inputBuffer.inputLength]))
		if input == ".exit" {
			os.Exit(0)
		} else {
			fmt.Printf("Unrecognised command '%s' \n", input)
		}
	}
}
