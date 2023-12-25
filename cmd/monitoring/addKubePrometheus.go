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

var addKubePrometheusCmd = &cobra.Command{
	Use:   "kube-prometheus",
	Short: "Add kube-prometheus to your Kuberentes cluster",
	Long:  `Add kube-prometheus, an open-source monitoring and metrics stack, to your environment`,
	Run: func(cmd *cobra.Command, arg []string) {

		var (
			chartName   = "prometheus-community/kube-prometheus-stack"
			releaseName = "kube-prometheus"
			namespace   = "monitoring"
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

		// Find the Helm Chart
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
		rel, err := client.Run(chart, nil)
		if err != nil {
			panic(err)
		}

		log.Printf("Installed Chart from path: %s in namespace: %s\n", rel.Name, rel.Namespace)
		// this will confirm the values set during installation
		log.Println(rel.Config)

	},
}

func init() {
	cmdd.RootCmd.AddCommand(addKubePrometheusCmd)
}
