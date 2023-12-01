package main

import (
	"capipe/cmd"
	_ "capipe/cmd/gitops"
	_ "capipe/cmd/iac"
)

func main() {
	cmd.Execute()
}
