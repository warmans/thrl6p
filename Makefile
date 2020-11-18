ifndef MAKE_DEBUG
    MAKEFLAGS += -s
endif

DEV_DB_USER ?= thrl6p
DEV_DB_PASS ?= insecure
DEV_DB_HOST ?= localhost
DEV_DB_DSN ?= "postgres://$(DEV_DB_USER):$(DEV_DB_PASS)@$(DEV_DB_HOST):5432/$(DEV_DB_USER)?sslmode=disable"

.PHONY: server.vendor
server.vendor:
	cd server && \
	go mod download && \
	go mod vendor && \
	mkdir -p ./vendor/github.com/grpc-ecosystem/grpc-gateway && \
	go list -f {{.Dir}} -m github.com/grpc-ecosystem/grpc-gateway | grep "github.com" && \
	rsync -aq --no-perms --no-owner --no-group --chmod=ugo=rwX `go list -f {{.Dir}} -m github.com/grpc-ecosystem/grpc-gateway`/* ./vendor/github.com/grpc-ecosystem/grpc-gateway/.

.PHONY: server.tools
server.tools:
	cd server && \
	for TOOL in `cat tools.go | grep _ | awk -F'"' '{print $$2}'`; do \
		go install $${TOOL}; \
	done;

.PHONY: generate
generate: generate.proto generate.code generate.api-client

.PHONY: generate.proto
generate.proto:
	cd server && rm -rf gen/api && prototool generate proto

.PHONY: generate.code
generate.code:
	cd server && go generate `go list ./...`

.PHONY: generate.api-client
generate.api-client:
	cd client && npm run generate-api-client

.PHONY: server.build
server.build:
	for CMD in $(shell cd server/cmd; ls -d *); do \
		$(MAKE) server.build.$$CMD; \
	done

.PHONY: server.build.%
server.build.%:
	echo "Building $*..."
	cd server/cmd/$*; GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o ../../bin/$*


.PHONY: server.run
server.run:
	cd server/bin && DB_DSN=$(DEV_DB_DSN) ./server


.PHONY: env.run
env.run:
	cd dev; POSTGRES_USER=$(DEV_DB_USER) POSTGRES_PASSWORD=$(DEV_DB_PASS) docker-compose up