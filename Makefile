check-module:
	@xcaddy list-modules

run-module:
	xcaddy run -c example/caddy.json

build-module:
	xcaddy build v2.7.4 --with github.com/devetek/caddy-dataprovider@develop=./

run-build-module:
	./caddy run -c example/caddy.json

run-test:
	@go test -v -cover ./...


include .makefiles/*.mk