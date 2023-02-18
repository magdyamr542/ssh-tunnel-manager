package configmanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Entry is an SSH configuration entry
type Entry struct {
	Name        string
	Description string
	Server      string
	User        string
	KeyFile     string
	LocalPort   int
	RemoteHost  string
	RemotePort  int
}

// ConfigManager manages SSH entries
type ConfigManager interface {
	GetConfiguration(entryName string) (Entry, error)
	GetConfigurations() ([]Entry, error)
	AddConfiguration(entry Entry) error
	RemoveConfiguration(entryName string) error
}

type manager struct {
	// dir is a path to the directory containing the configurations
	dir string
}

func NewManager(dir string) *manager {

	return &manager{dir: dir}
}

func (m *manager) GetConfiguration(entryName string) (Entry, error) { return Entry{}, nil }
func (m *manager) GetConfigurations() ([]Entry, error)              { return []Entry{}, nil }
func (m *manager) RemoveConfiguration(entryName string) error       { return nil }

func (m *manager) AddConfiguration(entry Entry) error {
	if _, err := os.Stat(m.dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(m.dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("couldn't create directory %s: %v", m.dir, err)
		}
	}

	file, err := json.MarshalIndent(entry, "", " ")
	if err != nil {
		return fmt.Errorf("couldn't marshal entry %v to JSON: %v", entry, err)
	}
	filename := filepath.Join(m.dir, entry.Name)
	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		return fmt.Errorf("couldn't write file %s: %v", filename, err)
	}
	return nil
}
