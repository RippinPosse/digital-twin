include .env

format:
	gci -w ./

build:
	docker build -t lod2021rp/processor:latest .

publish:
	docker build -t lod2021rp/processor:latest .
	docker push lod2021rp/processor:latest

deploy:
	helm install \
		--generate-name \
		./deployments/api \
		--kubeconfig $(kubepath) \
		--set port=${PORT} \
		--set hdfs_host=${HDFS_HOST} \
		--set hdfs_port=${HDFS_PORT} \

PHONY: format publish deploy
