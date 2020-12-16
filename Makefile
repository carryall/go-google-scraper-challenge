# Variables
BIN=node_modules/.bin
ASSETS_DIR=assets
SCSS_DIR=$(ASSETS_DIR)/stylesheets
JS_DIR=$(ASSETS_DIR)/javascripts
DIST_DIR=static
CSS_DIST=$(DIST_DIR)/css
JS_DIST=$(DIST_DIR)/js

.PHONY: dev assets test

build-dependencies:
	go get github.com/beego/bee
	go mod tidy
	npm install

dev:
	bee run

assets:
	$(BIN)/node-sass $(SCSS_DIR)/index.scss $(CSS_DIST)/application.css
	npx tailwindcss build $(SCSS_DIR)/vendors/tailwind.css -o $(CSS_DIST)/tailwind.css
	$(BIN)/minify $(JS_DIR) --out-dir $(JS_DIST)

test:
	go test -v -p 1 ./...
