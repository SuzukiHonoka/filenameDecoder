package rename

import (
	"filenameDecoder/pkg/config"
	"filenameDecoder/pkg/escape"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func BatchRename(cfg *config.Config) error {
	err := filepath.Walk(cfg.InputPath,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			return Rename(path, cfg, info)
		})
	return err
}

func Rename(path string, cfg *config.Config, info fs.FileInfo) error {
	if info.IsDir() {
		return nil
	}
	fileName := info.Name()
	if !escape.IsEscape(fileName) {
		return nil
	}
	unescapedName, err := escape.Decode(fileName)
	if err != nil {
		return err
	}
	unescapedName = DoFilter(unescapedName)
	dst := filepath.Join(filepath.Dir(path), unescapedName)
	fmt.Printf("%s -> %s\n", path, dst)
	if !cfg.DryRun {
		err = os.Rename(path, dst)
	}
	return err
}
