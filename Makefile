.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o chatgpt-web ./main.go

.PHONY: docker
docker:
	docker build . -t chatgpt-web:latest
