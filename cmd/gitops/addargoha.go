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
var ArgoHA string

var addArgoHACmd = &cobra.Command{
	Use:   "argocdha",
	Short: "Add ArgoCD In High-Availability Mode",
	Long:  `Add ArgoCD as your GitOps Controller In HA Mode`,
	Run: func(cmd *cobra.Command, arg []string) {

		var (
			chartName   = "argo/argo-cd"
			releaseName = "argocdha"
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

		vals := map[string]interface{}{
			"server": map[string]interface{}{
				"service": map[string]interface{}{
					"type": "LoadBalancer",
				},
				"autoscaling": map[string]interface{}{
					"enabled":     true,
					"minReplicas": 2,
				},
			},

			"redis-ha": map[string]interface{}{
				"enabled": true,
			},

			"controller": map[string]interface{}{
				"replicas": 1,
			},

			"repoServer": map[string]interface{}{
				"autoscaling": map[string]interface{}{
					"enabled":     true,
					"minReplicas": 2,
				},
			},

			"applicationSet": map[string]interface{}{
				"replicaCount": 2,
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
	cmdd.RootCmd.AddCommand(addArgoHACmd)
}
