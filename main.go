package main

import (
	"capipe/cmd"
	_ "capipe/cmd/appstack"
	_ "capipe/cmd/gitops"
	_ "capipe/cmd/iac"
	_ "capipe/cmd/monitoring"

	// _ "capipe/cmd/netstack"
	_ "capipe/cmd/networking"
	_ "capipe/cmd/security"
	// _ "capipe/cmd/virtstack"
)

func main() {
	cmd.Execute()
}
