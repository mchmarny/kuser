
# Assumes following env vars set
#  GCP_PROJECT - ID of your project
#  CLUSTER_ZONE - GCP Zone, ideally same as your Knative k8s cluster


.PHONY: test image rm

# DEV
test:
	go test ./... -v

# BUILD

image:
	gcloud builds submit \
		--project ${GCP_PROJECT} \
		--tag gcr.io/${GCP_PROJECT}/kuser:latest

# DEPLOYMENT

push:
	kubectl apply -f deployments/service.yaml

rm:
	kubectl delete -f deployments/service.yaml


