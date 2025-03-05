## helm useful commands

repo
helm repo add <repo-name>
helm repo update
helm repo remove <repo-name>
helm repo ls

install
helm install <release-name> <chart-path> -n <namespace> --create-namespace
helm uninstall <release-name> <chart-path> -n <namespace>
helm upgrade <release-name> <chart-path>
helm ls -A (for all namespace)

helm status <release-name>
helm rollback <release-name> <revision-name>
