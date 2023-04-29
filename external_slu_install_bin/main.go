package main

import "github.com/sikalabs/slu/cmd/install_bin"

func main() {
	install_bin.InstallBinForExternalGoUse("kubectl", "v1.25.0", "darwin", "amd64", ".")
}
