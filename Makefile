
# Assumes following env vars set
#  GCP_PROJECT - ID of your project
#  CLUSTER_ZONE - GCP Zone, ideally same as your Knative k8s cluster

.PHONY: test image rm

# DEV
test:
	go test ./... -v

# BUILD

image:
	go mod tidy
	go mod vendor
	gcloud builds submit \
		--project ${GCP_PROJECT} \
		--tag gcr.io/${GCP_PROJECT}/kuser:latest

# DEPLOYMENT

push:
	kubectl apply -f deployment/config.yaml
	kubectl apply -f deployment/service.yaml

rm:
	kubectl delete -f deployment/service.yaml


