package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	kubectl := "/usr/local/bin/kubectl"
	k3s := "/usr/local/bin/k3s"

	info, err := os.Lstat(kubectl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if info.Mode()&os.ModeSymlink != 0 {
		target, err := os.Readlink(kubectl)
		if err != nil {
			fmt.Println("Error reading symlink:", err)
			return
		}

		// Resolve both to absolute paths
		resolvedkubectl, err := filepath.EvalSymlinks(kubectl)
		if err != nil {
			fmt.Println("Error resolving kubectl symlink:", err)
			return
		}
		resolvedk3s, err := filepath.EvalSymlinks(k3s)
		if err != nil {
			fmt.Println("Error resolving k3s symlink:", err)
			return
		}

		fmt.Println("kubectl points to:", target)
		if resolvedkubectl == resolvedk3s {
			fmt.Println("/usr/local/bin/kubectl is a symlink to /usr/local/bin/k3s (same resolved path)")
		} else {
			fmt.Println("/usr/local/bin/kubectl is a symlink, but not to /usr/local/bin/k3s")
		}
	} else {
		fmt.Println("/usr/local/bin/kubectl is not a symlink")
	}
}
