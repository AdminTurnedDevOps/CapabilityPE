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
var Cilium string

var addCiliumCmd = &cobra.Command{
	Use:   "cilium",
	Short: "Add Cilium",
	Long:  `Add Cilium as your Container Network Interface(CNI)`,
	Run: func(cmd *cobra.Command, arg []string) {
		var (
			chartName   = "cilium/cilium"
			releaseName = "cilium"
			namespace   = "kube-system"
			args        = map[string]interface{}{
				// comma seperated values to set
				"set": "kubeProxyReplacement=strict,encryption.enabled=true,encryption.type=wireguard,l7Proxy=false",
			}
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
		client.Namespace = namespace
		client.ReleaseName = releaseName

		// Find the Cilium Helm Chart
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
		results, err := client.Run(chart, args)
		if err != nil {
			log.Printf("%+v", err)
		}

		fmt.Println(results)

	},
}

func init() {
	cmdd.RootCmd.AddCommand(addCiliumCmd)
}
