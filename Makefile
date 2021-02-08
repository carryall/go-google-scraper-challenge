# Include variables from ENV file
ENV =
include .env
ifdef ENV
include .env.$(ENV)
endif
export

# Variables
ASSETS_DIR=assets
SCSS_DIR=$(ASSETS_DIR)/stylesheets
JS_DIR=$(ASSETS_DIR)/javascripts
DIST_DIR=static
CSS_DIST=$(DIST_DIR)/css
JS_DIST=$(DIST_DIR)/js

.PHONY: build-dependencies test-dependency assets assets/css assets/js assets/icon-sprite dev db/setup db/migrate db/rollback lint test test/run

build-dependencies:
	go get github.com/beego/bee/v2
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.35.2
	npm install

assets:
	make assets/css
	make assets/js
	make assets/icon-sprite

assets/css:
	npm run build-scss
	npx tailwindcss build $(CSS_DIST)/application.css -o $(CSS_DIST)/application.css

assets/js:
	npm run minify-js

assets/icon-sprite:
	npm run generate-svg-sprite

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
	make test/run ENV=test

test/run:
	docker-compose -f docker-compose.test.yml up -d
	APP_RUN_MODE=test go test -v -p 1 ./...
	docker-compose -f docker-compose.test.yml down
