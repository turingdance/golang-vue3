package site

import (
	"fmt"
	"io/fs" // ✅ 必须这个，不是 fscompiler
	"path"
	"testing"
)

func printEmbedFiles(fsys fs.FS, prefix string) error {
	entries, err := fs.ReadDir(fsys, prefix)
	if err != nil {
		return err
	}
	for _, e := range entries {
		full := path.Join(prefix, e.Name())
		fmt.Println(full)
		if e.IsDir() {
			_ = printEmbedFiles(fsys, full)
		}
	}
	return nil
}

func TestAsset(t *testing.T) {
	printEmbedFiles(Assets, ".")
}
