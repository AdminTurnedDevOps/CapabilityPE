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

// add Datadog to your Kubernetes cluster
var apiKey string
var clusterName string

var addDatadogCmd = &cobra.Command{
	Use:   "datadog",
	Short: "Add Datadog to your Kuberentes cluster",
	Long:  `Add Datadog, a monitoring and observability enterprise platform, to your environment`,
	Run: func(cmd *cobra.Command, arg []string) {

		var (
			chartName   = "datadog/datadog"
			releaseName = "datadog"
			namespace   = "datadog"
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

		vals := map[string]interface{}{
			"datadog": map[string]interface{}{
				"site": "datadoghq.com",

				"clusterName": clusterName,

				"kubeStateMetricsEnabled": true,

				"kubeStateMetricsCore": map[string]interface{}{
					"enabled": true,
				},

				"apiKey": apiKey,

				"logs": map[string]interface{}{
					"enabled":             true,
					"containerCollectAll": true,
				},

				"processAgent": map[string]interface{}{
					"enabled": true,
				},

				"clusterAgent": map[string]interface{}{
					"replicas":                  "LoadBalancer",
					"createPodDisruptionBudget": true,
				},

				"autoscaling": map[string]interface{}{
					"enabled":     true,
					"minReplicas": 2,
				},
			},
		}

		// Install Chart
		rel, err := client.Run(chart, vals)
		if err != nil {
			panic(err)
		}

		log.Printf("Installed Chart from path: %s in namespace: %s\n", rel.Name, rel.Namespace)
		// this will confirm the values set during installation
		log.Println(rel.Config)

	},
}

func init() {
	cmdd.RootCmd.AddCommand(addDatadogCmd)

	addDatadogCmd.PersistentFlags().StringVarP(&apiKey, "apikey", "a", "", "Enter your Datadog API key")
	addDatadogCmd.PersistentFlags().StringVarP(&clusterName, "clustername", "n", "", "Enter the name of the cluster to be associated with Datadog")

}
