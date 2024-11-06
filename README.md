## Prerequisites

- [Helm 3.16+](https://helm.sh/docs/intro/install/)

## Steps to Run the Exercise

1. Run `go run main.go tufin cluster`.
   **_After running `tufin cluster`, make sure you set up the kubeconfig file._**

   ```bash
   mkdir ~/.kube
   sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
   export KUBECONFIG=~/.kube/config
   ```

   _If required, `chown` the file for your user._

2. Wait for the cluster to be completely initialized, and then run `go run main.go tufin deploy`.
3. Run `go run main.go tufin status` to check the status of the pods.

## Notes

- To access the Wordpress UI running in the deployed pod, you can utilize `kubectl port-forward` (the container port is `8080`).
- To access the Wordpress Admin UI, you can reach `/wp-admin`.
- To get the password for the admin user, use `kubectl get secret --namespace default wp-wordpress -o jsonpath="{.data.wordpress-password}" | base64 --decode` (username is `user`).
