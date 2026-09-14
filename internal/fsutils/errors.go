package fsutils

import "github.com/olekukonko/errors"

// -------------------------------------------------------------
// Path errors
// -------------------------------------------------------------

var (
	ErrInvalidPath     = errors.New("invalid path")
	ErrOutsideRoots    = errors.New("path is outside allowed root")
	ErrPathNotFound    = errors.New("path does not exist")
	ErrPathNotDir      = errors.New("path is not a directory")
	ErrPathNotFile     = errors.New("path is not a file")
	ErrPathNotReadable = errors.New("path is not readable")
	ErrPathNotWritable = errors.New("path is not writable")
)

// -------------------------------------------------------------
// Network / NAS errors
// -------------------------------------------------------------

var (
	ErrNetworkUnreachable = errors.New("network path is unreachable")
)

// -------------------------------------------------------------
// File I/O errors
// -------------------------------------------------------------

var (
	ErrReadFailure      = errors.New("failed to read file")
	ErrWriteFailure     = errors.New("failed to write file")
	ErrCreateFailure    = errors.New("failed to create file or directory")
	ErrDeleteFailure    = errors.New("failed to delete file")
	ErrRangeOutOfBounds = errors.New("requested read range is out of bounds")
)

// -------------------------------------------------------------
// Directory operations errors
// -------------------------------------------------------------

var (
	ErrDirNotAccessible = errors.New("directory is not accessible")
	ErrDirNotCreatable  = errors.New("directory cannot be created inside parent")
)
