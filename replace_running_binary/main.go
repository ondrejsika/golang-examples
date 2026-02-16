package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	if err := CopyOrReplaceBinary("hello-world-server", "hws"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("done: hello-world-server -> hws")
}

func CopyOrReplaceBinary(sourcePath, targetPath string) error {
	src, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("open source: %w", err)
	}
	defer src.Close()

	tmp := filepath.Join(filepath.Dir(targetPath), filepath.Base(targetPath)+".new")

	dst, err := os.Create(tmp)
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}

	if _, err := io.Copy(dst, src); err != nil {
		dst.Close()
		os.Remove(tmp)
		return fmt.Errorf("copy: %w", err)
	}

	if err := dst.Close(); err != nil {
		os.Remove(tmp)
		return fmt.Errorf("close temp file: %w", err)
	}

	if err := os.Chmod(tmp, 0755); err != nil {
		os.Remove(tmp)
		return fmt.Errorf("chmod: %w", err)
	}

	if err := os.Rename(tmp, targetPath); err != nil {
		os.Remove(tmp)
		return fmt.Errorf("rename: %w", err)
	}

	return nil
}
