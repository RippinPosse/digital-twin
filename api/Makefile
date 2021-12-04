include .env

format:
	gci -w ./

build:
	docker build -t lod2021rp/api:latest .

publish:
	docker build --no-cache -t lod2021rp/api:1.0.3 .
	docker push lod2021rp/api:1.0.3

deploy:
	helm install \
		--generate-name \
		./deployments/api \
		--kubeconfig $(kubepath) \
		--set port=${PORT} \
		--set hdfs_host=${HDFS_HOST} \
		--set hdfs_port=${HDFS_PORT} \
		--set tvs_address=${TVS_ADDRESS} \
		--set tvs_client_id=${TVS_CLIENT_ID} \
		--set tvs_client_secret=${TVS_CLIENT_SECRET}

PHONY: format publish deploy
