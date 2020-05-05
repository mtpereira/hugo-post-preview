.PHONY: docker-build

docker-build:
	docker build . --rm --tag hugo-post-preview:$(shell git rev-parse HEAD)

build: cmd/hugo-post-preview
	go build -o ./cmd/hugo-post-preview ./cmd/hugo-post-preview
