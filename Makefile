# Variables
BIN=node_modules/.bin
ASSETS_DIR=assets
SCSS_DIR=$(ASSETS_DIR)/stylesheets
JS_DIR=$(ASSETS_DIR)/javascripts
DIST_DIR=static
CSS_DIST=$(DIST_DIR)/css
JS_DIST=$(DIST_DIR)/js
DATABASE_URL=postgres://postgres@127.0.0.1:5432/google_scraper_development?sslmode=disable

.PHONY: build-dependencies assets dev test

build-dependencies:
	go get github.com/beego/bee/v2
	npm install

assets:
	$(BIN)/node-sass $(SCSS_DIR)/index.scss $(CSS_DIST)/application.css
	npx tailwindcss build $(SCSS_DIR)/vendors/tailwind.css -o $(CSS_DIST)/tailwind.css
	$(BIN)/minify $(JS_DIR) --out-dir $(JS_DIST)

dev:
	docker-compose -f docker-compose.dev.yml up -d
	make db-migrate
	bee run

db-migrate:
	bee migrate -driver=postgres -conn="$(DATABASE_URL)"

test:
	go test -v -p 1 ./...
