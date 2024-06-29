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
var AppStack3 string

var addAppStack3Cmd = &cobra.Command{
	Use:   "appstack3",
	Short: "Install ArgoCD, Crossplane, OPA, Kube-Prometheus",
	Long:  `Install The Following App Stack: ArgoCD, Crossplane, OPA, Kube-Prometheus`,
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
		// Kube-Prometheus
		//

		var (
			chartNameProm   = "prometheus-community/kube-prometheus-stack"
			releaseNameProm = "kube-prometheus"
			namespaceProm   = "monitoring"
		)

		// Call upon the CLI package
		settingsProm := cli.New()

		// Create a new instance of the configuration
		configProm := new(action.Configuration)

		// Collect local Helm information
		if err := configProm.Init(settingsProm.RESTClientGetter(), namespaceProm, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
			log.Printf("%+v", err)
		}

		// Create a new instance of the `Install` action, which is similar to running `helm install`
		clientProm := action.NewInstall(configProm)

		// Set metadata
		clientProm.CreateNamespace = true
		clientProm.Namespace = namespaceProm
		clientProm.ReleaseName = releaseNameProm

		// Find the Helm Chart
		chProm, err := clientProm.LocateChart(chartNameProm, settingsProm)
		if err != nil {
			fmt.Println(err)
		}

		// Load the chart to install
		chartProm, err := loader.Load(chProm)
		if err != nil {
			log.Println(err)
		}

		// Install Chart
		rel, err := clientProm.Run(chartProm, nil)
		if err != nil {
			panic(err)
		}

		log.Printf("Installed Chart from path: %s in namespace: %s\n", rel.Name, rel.Namespace)
		// this will confirm the values set during installation
		log.Println(rel.Config)

	},
}

func init() {
	cmdd.RootCmd.AddCommand(addAppStack3Cmd)
}
