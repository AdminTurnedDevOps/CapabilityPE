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
var Flux string

var addFluxCmd = &cobra.Command{
	Use:   "flux",
	Short: "Add Flux",
	Long:  `Add Flux as your GitOps Controller`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			chartName   = "fluxcd-community/flux2"
			releaseName = "flux"
			namespace   = "flux"
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

		// Find the Flux Helm Chart
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
	rootCmd.AddCommand(addFluxCmd)

	// addFluxCmd.PersistentFlags().StringVarP(&Flux, "flux", "flux", "", "Add Flux to your cluster")
}
