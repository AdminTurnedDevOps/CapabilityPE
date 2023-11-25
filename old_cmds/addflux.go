package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// add a GitOps Controller
var Flux string

var addFluxCmd = &cobra.Command{
	Use:   "flux",
	Short: "Add Flux",
	Long:  `Add Flux as your GitOps Controller`,
	Run: func(cmd *cobra.Command, args []string) {

		command := exec.Command("helm", "install", "flux", "-n", "flux", "fluxcd-community/flux2", "--create-namespace")

		output, err := command.Output()

		if err != nil {
			fmt.Println(err)
		}

		stringOut := string(output[:])
		fmt.Println(stringOut)

		fmt.Println("Success!")
	},
}

func init() {
	rootCmd.AddCommand(addFluxCmd)

	// addArgoCDCmd.PersistentFlags().StringVarP(&Flux, "flux", "flux", "", "Add Flux to your cluster")
}
