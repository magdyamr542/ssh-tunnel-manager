package editor

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	DefaultEditor = "vim"
)

// An Editor enables the user to "interactively" edit the originalContent in the user's editor.
// It returns the new content after editing and any errors encountered.
type Editor interface {
	Edit(originalContent []byte) ([]byte, error)
}

type editor struct{}

func New() Editor {
	return editor{}
}

func (e editor) Edit(content []byte) ([]byte, error) {
	ed := os.Getenv("EDITOR")
	if ed == "" {
		ed = DefaultEditor
	}

	f, err := os.CreateTemp("", "ssh-tunnel-manager-edit-*.json")
	if err != nil {
		return nil, fmt.Errorf("can't create a temp file for editing: %v", err)
	}

	defer func() {
		f.Close()
		os.Remove(f.Name())
	}()

	if _, err := f.Write(content); err != nil {
		return nil, err
	}

	cmd := exec.Command(ed, f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(f.Name())
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
