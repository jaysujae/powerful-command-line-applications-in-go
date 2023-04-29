package todo_test

import (
	"os"
	"testing"
	"todo"
)

func TestList_Add(t *testing.T) {
	l := todo.List{}
	taskName := "New York"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q intead", taskName, l[0].Task)
	}
}

func TestList_Complete(t *testing.T) {
	l := todo.List{}
	taskName := "New York"
	l.Add(taskName)
	if l[0].Done {
		t.Errorf("New task shouldn't be completed")
	}

	_ = l.Complete(1)

	if !l[0].Done {
		t.Errorf("Task should be completed.")
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}
	l1.Add("new challenge")

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("%s", err)
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("%s", err)
	}
	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("%s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Two should be same")
	}
}
