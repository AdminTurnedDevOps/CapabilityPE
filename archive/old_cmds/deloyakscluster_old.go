// ./capipe aks --name "aksenvironment01" --resourcegroup "devrelasaservice --location "eastus" --nodecount "3" --version "1.28.0

package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"

	cmdd "capipe/cmd"
)

var name string
var resourceGroupName string
var location string
var nodeCount string
var k8sVersion string

var addTerraformAKSCmd = &cobra.Command{
	Use:   "aks",
	Short: "Create an AKS cluster",
	Long:  `Create an Azure Kubernetes Service (AKS) cluster`,
	Run: func(cmd *cobra.Command, args []string) {

		installer := &releases.ExactVersion{
			Product: product.Terraform,
			Version: version.Must(version.NewVersion("1.0.7")),
		}

		execPath, err := installer.Install(context.Background())
		if err != nil {
			log.Fatalf("error installing Terraform: %s", err)
		}

		workingDir := "cmd/cluster-creation/aks"
		tf, err := tfexec.NewTerraform(workingDir, execPath)
		if err != nil {
			log.Fatalf("error running NewTerraform: %s", err)
		}

		err = tf.Init(context.Background(), tfexec.Upgrade(true))
		if err != nil {
			log.Fatalf("error running Init: %s", err)
		}

		_, err = tf.Plan(context.Background())
		if err != nil {
			log.Fatalf("error running Init: %s", err)
		}

		// variables := map[string]interface{}{
		// 	"name":                name,
		// 	"resource_group_name": resourceGroupName,
		// 	"location":            location,
		// 	"node_count":          nodeCount,
		// 	"k8s_version":         k8sVersion,
		// }

		// variables := tfexec.Var()

		tf.SetEnv(map[string]string{
			"name":                name,
			"resource_group_name": resourceGroupName,
			"location":            location,
			"node_count":          nodeCount,
			"k8s_version":         k8sVersion,
		})

		err = tf.Apply(context.Background())

		if err != nil {
			log.Fatalf("error running Apply: %s", err)
		}

		fmt.Println("done")

	},
}

func init() {
	cmdd.RootCmd.AddCommand(addTerraformAKSCmd)

	addTerraformAKSCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Enter the name of the Kubernetes cluster")
	addTerraformAKSCmd.PersistentFlags().StringVarP(&resourceGroupName, "resourcegroup", "g", "", "Enter Resource Group name")
	addTerraformAKSCmd.PersistentFlags().StringVarP(&location, "location", "l", "", "Enter region")
	addTerraformAKSCmd.PersistentFlags().StringVarP(&nodeCount, "nodecount", "", "", "Enter node count")
	addTerraformAKSCmd.PersistentFlags().StringVarP(&k8sVersion, "version", "v", "", "Enter Kubernetes version")
}
