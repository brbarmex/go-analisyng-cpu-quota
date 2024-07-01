k3d-up:
	k3d cluster create k3d-cluster

k3d-down:
	k3d cluster delete k3d-cluster

prom-up:
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo add grafana https://grafana.github.io/helm-charts
	helm repo update
	helm install prometheus prometheus-community/prometheus
	helm install grafana grafana/grafana	
#kubectl port-forward svc/prometheus-server 9090

run:
	docker network create monitoring