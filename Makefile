PROJECT:=k2

.PHONY: build
build:
	go build -tags sqlite -o k2 main.go
docs-admin:
	swag init -o app/admin/docs --exclude app/kx,app/hsh
docs-kx:
	swag init -o app/kx/docs --exclude app/admin,app/hsh
docs-hsh:
	swag init -o app/hsh/docs --exclude app/admin,app/kx,common/middleware

#.PHONY: test
#test:
#	go test -v ./... -cover

#.PHONY: docker
#docker:
#	docker build . -t go-admin:latest
