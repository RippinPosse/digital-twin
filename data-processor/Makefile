rebuild:
	docker build --no-cache -t lod2021rp/proc:latest .

build:
	docker build -t lod2021rp/proc:latest .

publish:
	docker build -t lod2021rp/proc:latest .
	docker push lod2021rp/proc:latest

deploy:
	helm install \
		--generate-name \
		./deployments/proc \
		--kubeconfig $(kubepath)

PHONY: publish deploy
