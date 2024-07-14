package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"

	cmdd "capipe/cmd"
)

var license string

var addKastenCmd = &cobra.Command{
	Use:   "kasten",
	Short: "Add Veeam Kasten For Kubernetes",
	Long:  `Add Kasten as your k8s backup solution`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			chartName   = "k10"
			releaseName = "k10"
			namespace   = "kasten-io"
		)

		vars := map[string]interface{}{
			"set": "license=" + license,
		}

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
		client.RepoURL = "https://charts.kasten.io/"

		// Find the Kasten Helm Chart
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
		results, err := client.Run(chart, vars)
		if err != nil {
			log.Printf("%+v", err)
		}

		fmt.Println(results)

	},
}

func init() {
	cmdd.RootCmd.AddCommand(addKastenCmd)
	addKastenCmd.PersistentFlags().StringVarP(&license, "license", "l", "", "Add Your Veeam Kasten License")

}
