package main

import (
	"capipe/cmd"
	_ "capipe/cmd/gitops"
	_ "capipe/cmd/iac"
	_ "capipe/cmd/networking"
	_ "capipe/cmd/security"
)

func main() {
	cmd.Execute()
}
