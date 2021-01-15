# Include variables from ENV file
include .env
export

# Variables
BIN=node_modules/.bin
ASSETS_DIR=assets
SCSS_DIR=$(ASSETS_DIR)/stylesheets
JS_DIR=$(ASSETS_DIR)/javascripts
DIST_DIR=static
CSS_DIST=$(DIST_DIR)/css
JS_DIST=$(DIST_DIR)/js

.PHONY: build-dependencies assets dev db/setup db/migrate db/rollback test

build-dependencies:
	go get github.com/beego/bee/v2
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.35.2
	npm install

assets:
	$(BIN)/node-sass $(SCSS_DIR)/index.scss $(CSS_DIST)/application.css
	npx tailwindcss build $(SCSS_DIR)/vendors/tailwind.css -o $(CSS_DIST)/tailwind.css
	$(BIN)/minify $(JS_DIR) --out-dir $(JS_DIST)

dev:
	make db/migrate
	bee run

db/setup:
	docker-compose -f docker-compose.dev.yml up -d

db/migrate:
	bee migrate -driver=postgres -conn="$(DATABASE_URL)"

db/rollback:
	bee migrate rollback -driver=postgres -conn="$(DATABASE_URL)"

lint:
	golangci-lint run

test:
	go test -v -p 1 ./...
