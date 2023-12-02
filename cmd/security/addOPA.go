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
var OPA string

var addOPACmd = &cobra.Command{
	Use:   "openpolicyagent",
	Short: "Add OPA",
	Long:  `Add OPA as a Policy Enforcer`,
	Run: func(cmd *cobra.Command, arg []string) {
		var (
			chartName   = "gatekeeper/gatekeeper"
			releaseName = "opa"
			namespace   = "gatekeeper-system"
			args        = map[string]interface{}{
				// comma seperated values to set
				"set": "name-template=gatekeeper",
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
		client.CreateNamespace = true
		client.Namespace = namespace
		client.ReleaseName = releaseName

		// Find the OPA Helm Chart
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
	cmdd.RootCmd.AddCommand(addOPACmd)
}
