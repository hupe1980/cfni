package main

import (
	"github.com/hupe1980/cfni/cmd"
)

var (
	version = "dev"
)

func main() {
	cmd.Execute(version)
}
