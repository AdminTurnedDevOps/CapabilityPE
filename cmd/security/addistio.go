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

// add an Istio Service Mesh
var Istio string

var addIstioCmd = &cobra.Command{
	Use:   "istio",
	Short: "Add Istio",
	Long:  `Add Istio as your Service Mesh`,
	Run: func(cmd *cobra.Command, arg []string) {
		var (
			istiodChartName   = "istio/istiod"
			istiodReleaseName = "istiod"
			namespace         = "istio-system"
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
		client.ReleaseName = istiodReleaseName

		// Find the Helm Chart for Istio
		ch2, err := client.LocateChart(istiodChartName, settings)
		if err != nil {
			fmt.Println(err)
		}

		// Load the chart to install
		chartConfig, err := loader.Load(ch2)
		if err != nil {
			log.Println(err)
		}

		result, err := client.Run(chartConfig, nil)
		if err != nil {
			fmt.Println((err))
		}

		fmt.Println(result)

	},
}

func init() {
	cmdd.RootCmd.AddCommand(addIstioCmd)
}
