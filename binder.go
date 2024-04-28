package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// InputBuffer represents a buffer for user input
type InputBuffer struct {
	buffer     []byte
	inputLen   int
	bufferSize int
}

// MetaCommandResult represents the result of executing a meta command
type MetaCommandResult int

const (
	MetaCommandSuccess MetaCommandResult = iota
	MetaCommandUnrecognizedCommand
)

// PrepareResult represents the result of preparing a SQL statement
type PrepareResult int

const (
	PrepareSuccess PrepareResult = iota
	PrepareUnrecognizedStatement
)

// StatementType represents the type of a SQL statement
type StatementType int

const (
	StatementInsert StatementType = iota
	StatementSelect
)

// Statement represents a SQL statement
type Statement struct {
	Type StatementType
}

// NewInputBuffer creates a new InputBuffer instance
func NewInputBuffer() *InputBuffer {
	return &InputBuffer{}
}

// ReadInput reads a line of input from stdin and stores it in the InputBuffer
func (ib *InputBuffer) ReadInput() error {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadBytes('\n')
	if err != nil {
		return err
	}

	ib.buffer = line
	ib.inputLen = len(line) - 1 // Remove the newline character
	ib.bufferSize = cap(line)

	return nil
}

func printPrompt() {
	fmt.Print("binderDB > ")
}

// doMetaCommand handles meta commands like .exit
func doMetaCommand(ib *InputBuffer) MetaCommandResult {
	if strings.TrimSpace(string(ib.buffer[:ib.inputLen])) == ".exit" {
		os.Exit(0)
	}
	return MetaCommandUnrecognizedCommand
}

// prepareStatement parses the input and determines the statement type
func prepareStatement(ib *InputBuffer, stmt *Statement) PrepareResult {
	input := strings.TrimSpace(string(ib.buffer[:ib.inputLen]))
	if strings.HasPrefix(input, "insert") {
		stmt.Type = StatementInsert
		return PrepareSuccess
	}
	if input == "select" {
		stmt.Type = StatementSelect
		return PrepareSuccess
	}
	return PrepareUnrecognizedStatement
}

// executeStatement executes the given statement (placeholder for now)
func executeStatement(stmt *Statement) {
	switch stmt.Type {
	case StatementInsert:
		fmt.Println("This is where we would do an insert.")
	case StatementSelect:
		fmt.Println("This is where we would do a select.")
	}
}

func main() {
	inputBuffer := NewInputBuffer()

	for {
		printPrompt()
		err := inputBuffer.ReadInput()
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		if inputBuffer.buffer[0] == '.' {
			switch doMetaCommand(inputBuffer) {
			case MetaCommandSuccess:
				continue
			case MetaCommandUnrecognizedCommand:
				fmt.Printf("Unrecognized command '%s'\n", inputBuffer.buffer[:inputBuffer.inputLen])
				continue
			}
		}

		stmt := &Statement{}
		switch prepareStatement(inputBuffer, stmt) {
		case PrepareSuccess:
			executeStatement(stmt)
			fmt.Println("Executed.")
		case PrepareUnrecognizedStatement:
			fmt.Printf("Unrecognized keyword at start of '%s'.\n", inputBuffer.buffer[:inputBuffer.inputLen])
			continue
		}
	}
}
