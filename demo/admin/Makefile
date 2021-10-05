PROJECT:=admin

.PHONY: build
build:
	go build -tags sqlite -o admin main.go
docs:
	swag init -o docs --parseDependency --parseDepth=6

#.PHONY: test
#test:
#	go test -v ./... -cover

#.PHONY: docker
#docker:
#	docker build . -t admin:latest
