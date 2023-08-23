check-module:
	@xcaddy list-modules

run-module:
	xcaddy run -c example/caddy.json

build-module:
	xcaddy build

run-test:
	@go test -v -cover ./...


include .makefiles/*.mk