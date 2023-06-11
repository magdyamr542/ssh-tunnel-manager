package configmanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// DefaultConfigDir is the default directory where configurations are saved
const DefaultConfigDir = "~/.ssh-tunnel-manager"

// Entry is an SSH configuration entry
type Entry struct {
	Name        string
	Description string
	Server      string
	User        string
	KeyFile     string
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
	m := &manager{dir: dir}
	if err := m.ensurePersistenceDirExists(); err != nil {
		panic(err)
	}
	return m
}

func (m *manager) ensurePersistenceDirExists() error {
	if _, err := os.Stat(m.dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(m.dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("couldn't create directory %s: %v", m.dir, err)
		}
	}
	return nil
}

func (m *manager) AddConfiguration(entry Entry) error {

	file, err := json.MarshalIndent(entry, "", " ")
	if err != nil {
		return fmt.Errorf("couldn't marshal entry %v to JSON: %v", entry, err)
	}
	filename := filepath.Join(m.dir, entry.Name+".json")
	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		return fmt.Errorf("couldn't write file %s: %v", filename, err)
	}
	return nil
}

func (m *manager) GetConfigurations() ([]Entry, error) {

	entries := make([]Entry, 0)
	files, err := ioutil.ReadDir(m.dir)
	if err != nil {
		return entries, fmt.Errorf("couldn't read directory %s: %v", m.dir, err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		byteValue, err := ioutil.ReadFile(filepath.Join(m.dir, file.Name()))
		if err != nil {
			return entries, fmt.Errorf("couldn't read file %s: %v", file.Name(), err)
		}

		var entry Entry
		if err := json.Unmarshal(byteValue, &entry); err != nil {
			return []Entry{}, fmt.Errorf("couldn't parse JSON file %s: %v", file.Name(), err)
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

func (m *manager) RemoveConfiguration(entryName string) error {
	return os.Remove(filepath.Join(m.dir, entryName+".json"))
}

func (m *manager) GetConfiguration(entryName string) (Entry, error) {
	filename := filepath.Join(m.dir, entryName+".json")
	byteValue, err := ioutil.ReadFile(filename)
	if err != nil {
		return Entry{}, fmt.Errorf("couldn't read file %q: %v", filename, err)
	}

	var entry Entry
	if err := json.Unmarshal(byteValue, &entry); err != nil {
		return Entry{}, fmt.Errorf("couldn't parse JSON file %q: %v", filename, err)
	}
	return entry, nil
}
