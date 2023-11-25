package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
)

// add a GitOps Controller
var ArgoCD string

var addArgoCDCmd = &cobra.Command{
	Use:   "argocd",
	Short: "Add ArgoCD",
	Long:  `Add ArgoCD as your GitOps Controller`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			chartName   = "argo/argo-cd"
			releaseName = "argocd"
			namespace   = "argocd"
			// args        = map[string]string{
			// 	// comma seperated values to set
			// 	"set": "mysqlRootPassword=admin@123,persistence.enabled=false,imagePullPolicy=Always",
			// }
		)

		// Call upon the CLI package
		settings := cli.New()

		// Create a new instance of the configuration
		config := new(action.Configuration)

		// Collect local Helm information
		if err := config.Init(settings.RESTClientGetter(), namespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
			log.Printf("%+v", err)
		}

		// Create a new instance of the `Install` action, which is similar to running `helm install`
		client := action.NewInstall(config)

		// Set metadata
		client.CreateNamespace = true
		client.Namespace = namespace
		client.ReleaseName = releaseName

		// Find the ArgoCD Helm Chart
		ch, err := client.LocateChart(chartName, settings)
		if err != nil {
			fmt.Println(err)
		}

		// Load the chart to install
		chart, err := loader.Load(ch)
		if err != nil {
			log.Println(err)
		}

		// Install Chart
		results, err := client.Run(chart, nil)
		if err != nil {
			log.Printf("%+v", err)
		}

		fmt.Println(results)

	},
}

func init() {
	rootCmd.AddCommand(addArgoCDCmd)

	// addArgoCDCmd.PersistentFlags().StringVarP(&ArgoCD, "argocd", "argo", "", "Add ArgoCD to your cluster")
}
