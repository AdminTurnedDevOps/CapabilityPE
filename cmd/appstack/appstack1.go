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

// add a GitOps Controller
var AppStack1 string
var apiKey string
var clusterName string

var addAppStack1Cmd = &cobra.Command{
	Use:   "appstack1",
	Short: "ArgoCD, Crossplane, OPA, Datadog",
	Long:  `Install The Following App Stack: ArgoCD, Crossplane, OPA, Datadog`,
	Run: func(cmd *cobra.Command, args []string) {

		//
		// ArgoCD
		//

		var (
			chartName   = "argo/argo-cd"
			releaseName = "argocd"
			namespace   = "argocd"
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

		//
		// Crossplane
		//

		var (
			chartNameCrossplane   = "crossplane-stable/crossplane"
			releaseNameCrossplane = "crossplane"
			namespaceCrossplane   = "crossplane"
		)

		// Call upon the CLI package
		settingsCrossplane := cli.New()

		// Create a new instance of the configuration
		configCrossplane := new(action.Configuration)

		// Collect local Helm information
		if err := configCrossplane.Init(settingsCrossplane.RESTClientGetter(), namespaceCrossplane, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
			log.Printf("%+v", err)
		}

		// Create a new instance of the `Install` action, which is similar to running `helm install`
		clientCrossplane := action.NewInstall(configCrossplane)

		// Set metadata
		clientCrossplane.CreateNamespace = true
		clientCrossplane.Namespace = namespaceCrossplane
		clientCrossplane.ReleaseName = releaseNameCrossplane

		// Find the crossplane Helm Chart
		chCrossplane, err := clientCrossplane.LocateChart(chartNameCrossplane, settingsCrossplane)
		if err != nil {
			fmt.Println(err)
		}

		// Load the chart to install
		chartCrossplane, err := loader.Load(chCrossplane)
		if err != nil {
			log.Println(err)
		}

		// Install Chart
		resultsCrossplane, err := clientCrossplane.Run(chartCrossplane, nil)
		if err != nil {
			log.Printf("%+v", err)
		}

		fmt.Println(resultsCrossplane)

		//
		// OPA
		//

		var (
			chartNameOPA   = "gatekeeper/gatekeeper"
			releaseNameOPA = "opa"
			namespaceOPA   = "gatekeeper-system"
			argsOPA        = map[string]interface{}{
				// comma seperated values to set
				"set": "name-template=gatekeeper",
			}
		)

		// Call upon the CLI package
		settingsOPA := cli.New()

		// Create a new instance of the configuration
		configOPA := new(action.Configuration)

		// Collect local Helm information
		if err := configOPA.Init(settingsOPA.RESTClientGetter(), namespaceOPA, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
			log.Printf("%+v", err)
		}

		// Create a new instance of the `Install` action, which is similar to running `helm install`
		clientOPA := action.NewInstall(configOPA)

		// Set metadata
		clientOPA.CreateNamespace = true
		clientOPA.Namespace = namespaceOPA
		clientOPA.ReleaseName = releaseNameOPA

		// Find the OPA Helm Chart
		chOPA, err := clientOPA.LocateChart(chartNameOPA, settingsOPA)
		if err != nil {
			fmt.Println(err)
		}

		// Load the chart to install
		chartOPA, err := loader.Load(chOPA)
		if err != nil {
			log.Println(err)
		}

		// Install Chart
		resultsOPA, err := clientOPA.Run(chartOPA, argsOPA)
		if err != nil {
			log.Printf("%+v", err)
		}

		fmt.Println(resultsOPA)

		//
		// Datadog
		//

		var (
			chartNameDatadog   = "datadog/datadog"
			releaseNameDatadog = "datadog"
			namespaceDatadog   = "datadog"
		)

		// Call upon the CLI package
		settingsDatadog := cli.New()

		// Create a new instance of the configuration
		configDatadog := new(action.Configuration)

		// Collect local Helm information
		if err := configDatadog.Init(settingsDatadog.RESTClientGetter(), namespaceDatadog, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
			log.Printf("%+v", err)
		}

		// Create a new instance of the `Install` action, which is similar to running `helm install`
		clientDatadog := action.NewInstall(configDatadog)

		// Set metadata
		clientDatadog.CreateNamespace = true
		clientDatadog.Namespace = namespaceDatadog
		clientDatadog.ReleaseName = releaseNameDatadog

		// Find the Helm Chart
		chDatadog, err := clientDatadog.LocateChart(chartNameDatadog, settingsDatadog)
		if err != nil {
			fmt.Println(err)
		}

		// Load the chart to install
		chartDatadog, err := loader.Load(chDatadog)
		if err != nil {
			log.Println(err)
		}

		valsDatadog := map[string]interface{}{
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
		// Install Chart
		resultsDatadog, err := clientDatadog.Run(chartDatadog, valsDatadog)
		if err != nil {
			log.Printf("%+v", err)
		}

		fmt.Println(resultsDatadog)

	},
}

func init() {
	cmdd.RootCmd.AddCommand(addAppStack1Cmd)

	addAppStack1Cmd.PersistentFlags().StringVarP(&apiKey, "apikey", "a", "", "Enter your Datadog API key")
	addAppStack1Cmd.PersistentFlags().StringVarP(&clusterName, "clustername", "n", "", "Enter the name of the cluster to be associated with Datadog")
}
