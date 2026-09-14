package fsutils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// -------------------------------------------------------------
// Internal private helpers (do NOT expose bool-return versions)
// -------------------------------------------------------------

func canRead(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	_ = f.Close()
	return true
}

func canWrite(path string) bool {
	// Directory
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		testFile := filepath.Join(path, ".sr_write_test")
		err := os.WriteFile(testFile, []byte("ok"), 0644)
		if err != nil {
			return false
		}
		_ = os.Remove(testFile)
		return true
	}

	// File (or file-to-be-created)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return false
	}
	_ = f.Close()
	return true
}

// -------------------------------------------------------------
// Existence & Type Checks
// -------------------------------------------------------------

func Exists(path string) bool {
	err := CheckExists(path)
	if err != nil {
		return false
	}

	return true
}

func CheckExists(path string) error {
	abs, err := NormalizePath(path)
	if err != nil {
		return err
	}
	if _, err := os.Stat(abs); err != nil {
		return ErrPathNotFound
	}
	return nil
}

func CheckDir(path string) error {
	if err := CheckExists(path); err != nil {
		return err
	}
	abs, _ := NormalizePath(path)
	info, _ := os.Stat(abs)
	if !info.IsDir() {
		return ErrPathNotDir
	}
	return nil
}

func CheckFile(path string) error {
	if err := CheckExists(path); err != nil {
		return err
	}
	abs, _ := NormalizePath(path)
	info, _ := os.Stat(abs)
	if info.IsDir() {
		return ErrPathNotFile
	}
	return nil
}

// -------------------------------------------------------------
// Permission Checks
// -------------------------------------------------------------

func CheckReadable(path string) error {
	abs, err := NormalizePath(path)
	if err != nil {
		return err
	}
	if !canRead(abs) {
		return ErrPathNotReadable
	}
	return nil
}

func CheckWritable(path string) error {
	abs, err := NormalizePath(path)
	if err != nil {
		return err
	}
	if !canWrite(abs) {
		return ErrPathNotWritable
	}
	return nil
}

func CheckReadableDir(path string) error {
	if err := CheckDir(path); err != nil {
		return err
	}
	return CheckReadable(path)
}

func CheckWritableDir(path string) error {
	if err := CheckDir(path); err != nil {
		return err
	}
	return CheckWritable(path)
}

func CheckAccessibleDir(path string) error {
	if err := CheckDir(path); err != nil {
		return err
	}
	if err := CheckReadable(path); err != nil {
		return err
	}
	if err := CheckWritable(path); err != nil {
		return err
	}
	return nil
}

// -------------------------------------------------------------
// Directory creation capability
// -------------------------------------------------------------

func CheckCanCreateDir(parent string) error {
	abs, err := NormalizePath(parent)
	if err != nil {
		return err
	}

	if err := CheckAccessibleDir(abs); err != nil {
		return err
	}

	test := filepath.Join(abs, ".sr_dir_test")
	if err := os.Mkdir(test, 0755); err != nil {
		return ErrDirNotCreatable
	}
	_ = os.Remove(test)

	return nil
}

// -------------------------------------------------------------
// Network / NAS detection
// -------------------------------------------------------------

func IsNetworkPath(path string) bool {
	if runtime.GOOS == "windows" {
		return strings.HasPrefix(path, `\\`)
	}

	// Common NAS mount prefixes on macOS/Linux
	if strings.HasPrefix(path, "/mnt/") ||
		strings.HasPrefix(path, "/media/") ||
		strings.HasPrefix(path, "/Volumes/") {
		return true
	}

	return false
}

func PingUNCPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CheckNetworkReachable(path string) error {
	abs, err := NormalizePath(path)
	if err != nil {
		return err
	}

	if !IsNetworkPath(abs) {
		return nil // not a network path, so it's "reachable"
	}

	if !PingUNCPath(abs) {
		return ErrNetworkUnreachable
	}

	return nil
}

// -------------------------------------------------------------
// Validate *any* user-provided location (file or directory)
// -------------------------------------------------------------

func ValidateLocation(path string) error {
	if path == "" {
		return ErrInvalidPath
	}

	abs, err := NormalizePath(path)
	if err != nil {
		return err
	}

	// Network path reachability
	if err := CheckNetworkReachable(abs); err != nil {
		return err
	}

	// Existing path: validate read/write
	if _, err := os.Stat(abs); err == nil {
		// Existing directory
		info, _ := os.Stat(abs)
		if info.IsDir() {
			if err := CheckReadableDir(abs); err != nil {
				return err
			}
			if err := CheckWritableDir(abs); err != nil {
				return err
			}
			return nil
		}

		// Existing file
		if err := CheckReadable(abs); err != nil {
			return err
		}
		if err := CheckWritable(abs); err != nil {
			return err
		}
		return nil
	}

	// Non-existing path: check parent
	parent := filepath.Dir(abs)

	if err := CheckDir(parent); err != nil {
		return fmt.Errorf("parent directory does not exist or is invalid: %s", parent)
	}

	if err := CheckWritableDir(parent); err != nil {
		return ErrPathNotWritable
	}

	return nil
}
