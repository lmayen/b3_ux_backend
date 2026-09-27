package assets

import (
	"b3_ux_backend/internal/config"
	"b3_ux_backend/internal/fsutils"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/olekukonko/errors"
)

func exportAssets(assets embed.FS, destination string) error {
	return fs.WalkDir(assets, ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path == "." {
			return nil
		}

		targetPath := filepath.Join(destination, filepath.FromSlash(path))
		if fsutils.Exists(targetPath) {
			return nil
		}

		if entry.IsDir() {
			return os.MkdirAll(targetPath, 0o755)
		}

		data, err := assets.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read embedded file %q: %w", path, err)
		}

		if err := os.MkdirAll(filepath.Dir(targetPath), 0o755); err != nil {
			return fmt.Errorf("create directory for %q: %w", targetPath, err)
		}

		if err := os.WriteFile(targetPath, data, 0o644); err != nil {
			return fmt.Errorf("write file %q: %w", targetPath, err)
		}

		return nil
	})
}

func InitAssets() error {
	err := exportAssets(DataFS, config.CONFIG.App.RootDir)
	if err != nil {
		return errors.Wrapf(err, "export assets %q", config.CONFIG.App.DataDir)
	}

	err = exportAssets(ImageFS, config.CONFIG.App.RootDir)
	if err != nil {
		return errors.Wrapf(err, "export assets %q", config.CONFIG.App.ImageDir)
	}

	return nil
}
