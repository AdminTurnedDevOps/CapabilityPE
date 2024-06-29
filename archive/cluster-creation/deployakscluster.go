package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"

	cmdd "capipe/cmd"
)

var directory string

var addTerraformAKSCmd = &cobra.Command{
	Use:   "aks",
	Short: "Create an AKS cluster",
	Long:  `Create an Azure Kubernetes Service (AKS) cluster`,
	Run: func(cmd *cobra.Command, args []string) {

		init := exec.Command("terraform", "-chdir=%q", "init")
		ex, _ := init.Output()
		fmt.Println(ex, directory)

		// plan := exec.Command("terraform", "-chdir=%q", "plan", directory)
		// plan.Output()

		// apply := exec.Command("terraform", "-chdir=cmd/cluster-creation/aks")
		// out, err := apply.Output()
		// if err != nil {
		// 	fmt.Println(out)
		// }
	},
}

func init() {
	cmdd.RootCmd.AddCommand(addTerraformAKSCmd)

	addTerraformAKSCmd.PersistentFlags().StringVarP(&directory, "dir", "d", "", "Enter the directory where your TF code exists")
}
