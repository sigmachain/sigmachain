package main

import (
	_ "embed"

	"github.com/sigmachain/sigmachain/command/root"
	"github.com/sigmachain/sigmachain/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
