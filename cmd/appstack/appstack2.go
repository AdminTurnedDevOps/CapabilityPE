package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"

	cmdd "capipe/cmd"
)

var addAppStack2Cmd = &cobra.Command{
	Use:   "appstack2",
	Short: "Install FluxCD, Crossplane, Kyverno, Grafana/Prometheus/Loki/Tempo",
	Long:  `Install The Following App Stack: FluxCD, Crossplane, Kyverno, Grafana/Prometheus/Loki/Tempo`,
	Run: func(cmd *cobra.Command, args []string) {

		//
		// FluxCD
		//

		flux := exec.Command(
			"helm",
			"install",
			"-n", "flux-system",
			"flux",
			"oci://ghcr.io/fluxcd-community/charts/flux2",
			"--create-namespace",
		)
		flux.Run()
		flux.Output()

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
		// Kyverno
		//

		var (
			chartNameKyverno   = "kyverno/kyverno"
			releaseNameKyverno = "kyverno"
			namespaceKyverno   = "kyverno"
		)

		// Call upon the CLI package
		settingsKyverno := cli.New()

		// Create a new instance of the configuration
		configKyverno := new(action.Configuration)

		// Collect local Helm information
		if err := configKyverno.Init(settingsKyverno.RESTClientGetter(), namespaceKyverno, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
			log.Printf("%+v", err)
		}

		// Create a new instance of the `Install` action, which is similar to running `helm install`
		clientKyverno := action.NewInstall(configKyverno)

		// Set metadata
		clientKyverno.CreateNamespace = true
		clientKyverno.Namespace = namespaceKyverno
		clientKyverno.ReleaseName = releaseNameKyverno

		// Find the Helm Chart
		chKyverno, err := clientKyverno.LocateChart(chartNameKyverno, settingsKyverno)
		if err != nil {
			fmt.Println(err)
		}

		// Load the chart to install
		chartKyverno, err := loader.Load(chKyverno)
		if err != nil {
			log.Println(err)
		}

		// Install Chart
		resultsKyverno, err := clientKyverno.Run(chartKyverno, nil)
		if err != nil {
			log.Printf("%+v", err)
		}

		fmt.Println(resultsKyverno)

		////////////////////////////////
		// Grafana/Prometheus/Loki/Tempo
		////////////////////////////////

		//
		// Kube-Prometheus
		//

		var (
			chartNameGraf   = "prometheus-community/kube-prometheus-stack"
			releaseNameGraf = "kubeprometheus"
			namespaceGraf   = "monitoring"
		)

		// Call upon the CLI package
		settingsGraf := cli.New()

		// Create a new instance of the configuration
		configGraf := new(action.Configuration)

		// Collect local Helm information
		if err := configGraf.Init(settingsGraf.RESTClientGetter(), namespaceGraf, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
			log.Printf("%+v", err)
		}

		// Create a new instance of the `Install` action, which is similar to running `helm install`
		clientGraf := action.NewInstall(configGraf)

		// Set metadata
		clientGraf.CreateNamespace = true
		clientGraf.Namespace = namespaceGraf
		clientGraf.ReleaseName = releaseNameGraf

		// Find the Helm Chart
		chGraf, err := clientGraf.LocateChart(chartNameGraf, settingsGraf)
		if err != nil {
			fmt.Println(err)
		}

		// Load the chart to install
		chartGraf, err := loader.Load(chGraf)
		if err != nil {
			log.Println(err)
		}

		// Install Chart
		resultsGraf, err := clientGraf.Run(chartGraf, nil)
		if err != nil {
			log.Printf("%+v", err)
		}

		fmt.Println(resultsGraf)

		//
		// Tempo
		//

		var (
			chartNameTempo   = "grafana/tempo"
			releaseNameTempo = "tempo"
			namespaceTempo   = "monitoring"
		)

		// Call upon the CLI package
		settingsTempo := cli.New()

		// Create a new instance of the configuration
		configTempo := new(action.Configuration)

		// Collect local Helm information
		if err := configTempo.Init(settingsTempo.RESTClientGetter(), namespaceTempo, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
			log.Printf("%+v", err)
		}

		// Create a new instance of the `Install` action, which is similar to running `helm install`
		clientTempo := action.NewInstall(configTempo)

		// Set metadata
		clientTempo.Namespace = namespaceTempo
		clientTempo.ReleaseName = releaseNameTempo

		// Find the Helm Chart
		chTempo, err := clientTempo.LocateChart(chartNameTempo, settingsTempo)
		if err != nil {
			fmt.Println(err)
		}

		// Load the chart to install
		chartTempo, err := loader.Load(chTempo)
		if err != nil {
			log.Println(err)
		}

		// Install Chart
		resultsTempo, err := clientTempo.Run(chartTempo, nil)
		if err != nil {
			log.Printf("%+v", err)
		}

		fmt.Println(resultsTempo)

		//
		// Loki
		//

		var (
			chartNameLoki   = "grafana/loki-stack"
			releaseNameLoki = "loki"
			namespaceLoki   = "monitoring"
		)

		// Call upon the CLI package
		settingsLoki := cli.New()

		// Create a new instance of the configuration
		configLoki := new(action.Configuration)

		// Collect local Helm information
		if err := configLoki.Init(settingsLoki.RESTClientGetter(), namespaceLoki, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
			log.Printf("%+v", err)
		}

		// Create a new instance of the `Install` action, which is similar to running `helm install`
		clientLoki := action.NewInstall(configLoki)

		// Set metadata
		clientLoki.Namespace = namespaceLoki
		clientLoki.ReleaseName = releaseNameLoki

		// Find the Helm Chart
		chLoki, err := clientLoki.LocateChart(chartNameLoki, settingsLoki)
		if err != nil {
			fmt.Println(err)
		}

		// Load the chart to install
		chartLoki, err := loader.Load(chLoki)
		if err != nil {
			log.Println(err)
		}

		// Install Chart
		resultsLoki, err := clientLoki.Run(chartLoki, nil)
		if err != nil {
			log.Printf("%+v", err)
		}

		fmt.Println(resultsLoki)

	},
}

func init() {
	cmdd.RootCmd.AddCommand(addAppStack2Cmd)

	// addAppStack1Cmd.PersistentFlags().StringVarP(&apiKey, "apikey", "a", "", "Enter your Datadog API key")
	// addAppStack1Cmd.PersistentFlags().StringVarP(&clusterName, "clustername", "n", "", "Enter the name of the cluster to be associated with Datadog")
}
