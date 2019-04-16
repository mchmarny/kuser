
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
		--tag gcr.io/${GCP_PROJECT}/kuser

sample-image:
	go mod tidy
	go mod vendor
	gcloud builds submit \
		--project knative-samples \
		--tag gcr.io/knative-samples/kuser

# DEPLOYMENT

service:
	kubectl apply -f deployment/config.yaml -n demo
	kubectl apply -f deployment/service.yaml  -n demo

rm:
	kubectl delete -f deployment/service.yaml  -n demo


