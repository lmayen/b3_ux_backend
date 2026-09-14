package fsutils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func NormalizePath(path string) (string, error) {
	if path == "" {
		return "", ErrInvalidPath
	}

	if runtime.GOOS == "windows" && strings.HasPrefix(path, `\\`) {
		return normalizeUNC(path)
	}

	clean := filepath.Clean(path)
	abs, err := filepath.Abs(clean)
	if err != nil {
		return "", fmt.Errorf("failed to resolve absolute path: %w", err)
	}

	return abs, nil
}

func normalizeUNC(path string) (string, error) {
	// Keep first two \\ as-is
	if !strings.HasPrefix(path, `\\`) {
		return "", ErrInvalidPath
	}

	// Split: ["", "", "server", "share", "folder", "file"]
	parts := strings.Split(path, `\`)

	if len(parts) < 4 {
		return "", fmt.Errorf("invalid UNC path: %s", path)
	}

	// Reconstruct UNC preserving double leading slash
	// and cleaning the rest
	server := parts[2]
	share := parts[3]
	rest := parts[4:]

	// Clean internal part using filepath.Clean (safe here)
	cleaned := filepath.Clean(filepath.Join(rest...))

	// Full UNC path: \\server\share\cleaned
	unc := fmt.Sprintf(`\\%s\%s`, server, share)
	if cleaned != "." {
		unc = filepath.Join(unc, cleaned)
	}

	return unc, nil
}

func EnsureDir(path string) error {
	if path == "" {
		return ErrInvalidPath
	}
	return os.MkdirAll(path, 0755)
}

func HomeDir() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to find home directory: %w", err)
	}
	return NormalizePath(dir)
}

func AppDir(appName string) (string, error) {
	home, err := HomeDir()
	if err != nil {
		return "", err
	}

	var docPath string

	switch runtime.GOOS {
	case "windows":
		docPath = filepath.Join(home, "Documents")
	case "darwin":
		docPath = filepath.Join(home, "Documents")
	default: // Linux & others
		docPath = filepath.Join(home, "Documents")
	}

	return NormalizePath(filepath.Join(docPath, appName))
}
