package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("building tool..")

	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Can't build")
		os.Exit(1)
	}
	fmt.Println("running test...")
	result := m.Run()
	fmt.Println("cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task1 := "task test 1"
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)
	t.Run("Add new task from arguments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task1)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
	task2 := "task 2"
	t.Run("Add new task from command", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")

		cmdStdin, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		io.WriteString(cmdStdin, task2)
		cmdStdin.Close()
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("List task", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := fmt.Sprintf("  1: %s\n  2: %s\n", task1, task2)
		if expected != string(out) {
			t.Errorf("expected %q, got %q", expected, string(out))
		}
	})
}
