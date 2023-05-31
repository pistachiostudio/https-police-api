.PHONY: dev
dev:
	wrangler dev

.PHONY: build
build:
	go run github.com/syumai/workers/cmd/workers-assets-gen@latest
	tinygo build -o ./build/app.wasm -target wasm ./...

.PHONY: deploy
deploy:
	wrangler deploy

.PHONY: generate
generate:
	go generate ./...

.PHONY: create-db
create-db:
	wrangler d1 create http-police-db

.PHONY: init-db-preview
init-db-local:
	wrangler d1 execute http-police-db --local --file=./schema.

.PHONY: init-db
init-db:
	wrangler d1 execute http-police-db --file=./schema.sql
