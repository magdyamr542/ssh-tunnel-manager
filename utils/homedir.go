package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ResolveDir resolves the given dir if it uses the telda
func ResolveDir(dir string) (string, error) {
	if strings.HasPrefix(dir, "~/") {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("couldn't resolve user's HOME directory from %s: %v", dir, err)
		}
		return filepath.Join(homedir, dir[2:]), nil
	}
	return dir, nil
}
