package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// add a GitOps Controller
var ArgoCD string

var addArgoCDCmd = &cobra.Command{
	Use:   "argocd",
	Short: "Add ArgoCD",
	Long:  `Add ArgoCD as your GitOps Controller`,
	Run: func(cmd *cobra.Command, args []string) {

		command := exec.Command("helm", "install", "argocd", "-n", "argocd", "argo/argo-cd", "--create-namespace")

		output, err := command.Output()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(addArgoCDCmd)

	// addArgoCDCmd.PersistentFlags().StringVarP(&ArgoCD, "argocd", "argo", "", "Add ArgoCD to your cluster")
}
