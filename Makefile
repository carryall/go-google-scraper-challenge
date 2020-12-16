# Variables
AUTOPREFIXER_BROWSERS="> 2%"
ASSETS_DIR=assets/stylesheets
BIN=node_modules/.bin
DIST_DIR=static/css
POSTCSS_FLAGS = --use autoprefixer --autoprefixer.overrideBrowserslist "> 2%"

.PHONY: dev assets test

build-dependencies:
	go get github.com/beego/bee
	go mod tidy

dev:
	make assets
	bee run

assets:
	$(BIN)/node-sass $(ASSETS_DIR)/index.scss $(DIST_DIR)/application.css
	$(BIN)/postcss $(POSTCSS_FLAGS) -o $(DIST_DIR)/tailwind.css $(ASSETS_DIR)/vendors/tailwind.css

test:
	go test -v -p 1 ./...
