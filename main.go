package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Tufin struct {
		Cluster struct{} `cmd:"cluster" help:"Deploy a k3s k8s cluster."`
		Deploy  struct{} `cmd:"deploy" help:"Deploy 2 pods in the cluster."`
		Status  struct{} `cmd:"status" help:"Print out status table containing the status of pod names and status in default namespace."`
	} `cmd:"tufin" help:""`
}

var (
	clusterInstallCommand = "curl -sfL https://get.k3s.io | sh -"
	deployCommand         = "helm repo add bitnami https://charts.bitnami.com/bitnami && helm repo update && helm install wp-mysql bitnami/mysql --namespace default --set auth.rootPassword=root --set auth.database=wordpress --set auth.username=wordpress --set auth.password=wordpresspassword && helm install wp bitnami/wordpress --namespace default --set externalDatabase.host=wp-mysql --set externalDatabase.user=wordpress --set externalDatabase.password=wordpresspassword --set externalDatabase.database=wordpress --set mariadb.enabled=false --set service.type=ClusterIP --set networkPolicy.enabled=false"
	statusCommand         = "kubectl get pods --namespace default"
)

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "tufin cluster":
		cmdCheckK3s := exec.Command("/bin/sh", "-c", "which k3s")
		err := cmdCheckK3s.Run()
		if err == nil {
			fmt.Println("k3s is already installed. Exiting.")
			return
		}

		cmd := exec.Command("/bin/sh", "-c", clusterInstallCommand)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("failed with error: %v, output: %s", err, string(out))
			return
		}
		fmt.Println(string(out))
	case "tufin deploy":
		cmd := exec.Command("/bin/sh", "-c", deployCommand)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("failed with error: %v, output: %s", err, string(out))
			return
		}
		fmt.Println(string(out))
	case "tufin status":
		cmd := exec.Command("/bin/sh", "-c", statusCommand)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("failed with error: %v, output: %s", err, string(out))
			return
		}
		fmt.Println(string(out))
	default:
		panic(ctx.Command())
	}

}
