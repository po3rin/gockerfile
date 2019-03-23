GATEWAY_IMAGE := po3rin/gocker
IMAGE := po3rin/gockersample

.PHONY: debug
debug:
	go run cmd/gocker/main.go -graph | buildctl debug dump-llb --dot | dot -T png -o out.png

.PHONY: build-gocker
build-gocker:
	go run cmd/gocker/main.go -graph | buildctl build --exporter=docker --exporter-opt name=buildkit0 | docker load

.PHONY: build-buildctl
build-buildctl:
	buildctl build \
		--frontend=gateway.v0 \
		--frontend-opt=source=$(GATEWAY_IMAGE) \
		--local dockerfile=. \
		--exporter=docker \
		--exporter-opt name=$(IMAGE) | docker load

.PHONY: image
image:
	docker build . -t $(GATEWAY_IMAGE) && docker push $(GATEWAY_IMAGE)

.PHONY: run
run:
	docker run -it -p 8080:8080 po3rin/gockersample:latest /bin/server
