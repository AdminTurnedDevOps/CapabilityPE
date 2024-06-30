// COMMAND
// ./capipe aks --name "myakscluster" --rg "devrelasaservice" --region "eastus" --nodecount "3" --k8sversion "1.29.4"
//

package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"

	cmdd "capipe/cmd"
)

var name string
var resource_group_name string
var location string
var node_count string
var k8s_version string

var addTerraformAKSCmd = &cobra.Command{
	Use:   "aks",
	Short: "Create an AKS cluster",
	Long:  `Create an Azure Kubernetes Service (AKS) cluster`,
	Run: func(cmd *cobra.Command, args []string) {

		init := exec.Command("terraform", "-chdir=cmd/cluster-creation/aks", "init")

		init.Run()
		init.Output()

		plan := exec.Command("terraform", "-chdir=cmd/cluster-creation/aks", "plan", "-var=name="+name, "-var=resource_group_name="+resource_group_name, "-var=location="+location, "-var=node_count="+node_count, "-var=k8s_version="+k8s_version)
		plan.Run()
		plan.Output()

		apply := exec.Command("terraform", "-chdir=cmd/cluster-creation/aks", "apply", "-var=name="+name, "-var=resource_group_name="+resource_group_name, "-var=location="+location, "-var=node_count="+node_count, "-var=k8s_version="+k8s_version, "-auto-approve")
		apply.Run()
		apply.Output()

		//
		//
		//
		//

		// THE BELOW WORKS IF YOU DONT WANT TO PASS IN VARIABLES AT RUNTIME

		// plan := exec.Command("terraform", "plan", "-chdir=cmd/cluster-creation/aks")

		// plan.Run()
		// plan.Output()

		// apply := exec.Command("terraform", "-chdir=cmd/cluster-creation/aks", "apply", "-auto-approve")
		// apply.Run()
		// apply.Output()
	},
}

func init() {
	cmdd.RootCmd.AddCommand(addTerraformAKSCmd)

	addTerraformAKSCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name of the cluster")
	addTerraformAKSCmd.PersistentFlags().StringVarP(&resource_group_name, "rg", "r", "", "Name of resource group that the cluster resides in")
	addTerraformAKSCmd.PersistentFlags().StringVarP(&location, "region", "l", "", "Region that the cluster resides in")
	addTerraformAKSCmd.PersistentFlags().StringVarP(&node_count, "nodecount", "c", "", "How many Worker Nodes in the cluster (default is 3)")
	addTerraformAKSCmd.PersistentFlags().StringVarP(&k8s_version, "k8sversion", "v", "", "Kubernetes version (default is 1.29.4)")
}
