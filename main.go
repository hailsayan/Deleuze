package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type command string

const (
	INSERT    command = "INSERT"
	LEFT      command = "LEFT"
	RIGHT     command = "RIGHT"
	BACKSPACE command = "BACKSPACE"
)

type element struct {
	data byte
	prev *element
	next *element
}

func main() {
	var count int
	fmt.Scan(&count)

	reader := bufio.NewReader(os.Stdin)

	for idx := 0; idx < count; idx++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		fields := strings.Fields(line)

		switch command(fields[0]) {
		case INSERT:
			fmt.Println("INSERT")

		case LEFT:
			fmt.Println("LEFT")

		case RIGHT:
			fmt.Println("RIGHT")

		case BACKSPACE:
			fmt.Println("BACKSPACE")

		default:
			panic("invalid command given")

		}

	}
}
