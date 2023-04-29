package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"todo"
)

var todoFile = os.Getenv("TODO_FILENAME")

func main() {
	add := flag.Bool("add", false, "Task to include")
	list := flag.Bool("list", false, "gonna list?")
	complete := flag.Int("complete", 0, "item to complete")

	if todoFile == "" {
		todoFile = ".todo.json"
	}
	flag.Parse()
	l := &todo.List{}
	if err := l.Get(todoFile); err != nil {
		panic(err)
	}

	switch {
	case *list:
		fmt.Print(l)
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFile); err != nil {
			panic(err)
		}
	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			panic(err)
		}
		l.Add(t)
		if err := l.Save(todoFile); err != nil {
			panic(err)
		}
	default:
		fmt.Fprintln(os.Stderr, "invalid option")
	}
}

func getTask(reader io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	s := bufio.NewScanner(reader)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}
	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task can't be nil")
	}

	return s.Text(), nil
}
