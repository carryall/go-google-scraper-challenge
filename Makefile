# Variables
# APP=application
AUTOPREFIXER_BROWSERS="> 2%"

# JS=assets/javascripts
ASSETS_DIR=assets/stylesheets
# SASS_FILE=assets/stylesheets/index.scss
# CSS_FILE=static/css/application.css
BIN=node_modules/.bin
DIST_DIR=static/css

# POSTCSS = ./node_modules/.bin/postcss
POSTCSS_FLAGS = --use autoprefixer --autoprefixer.overrideBrowserslist "> 2%"
# POSTCSS_FLAGS = --use autoprefixer autoprefixer.browsers "> 2%"
# # wildcard expansion is performed only in targets and in prerequisites so here we need $(wildcard)
# SOURCES = $(wildcard static/css/*.css)
# # use $(SOURCES) to determine what we actually need using a bit of pattern substitution
# TARGETS = $(SOURCES)

# CSS_C=sass
# CSS_FLAGS=
# CSS_SRC=assets/stylesheets
# CSS_OUT=static/css
# CSS_TARGETS=$(patsubst $(CSS_SRC)/%.scss,%.css,$(wildcard $(CSS_SRC)/*.scss))

.PHONY: dev assets test

build-dependencies:
	go get github.com/beego/bee

dev:
	make assets
	bee run

assets:
	$(BIN)/node-sass $(ASSETS_DIR)/index.scss $(DIST_DIR)/application.css
	$(BIN)/postcss $(POSTCSS_FLAGS) -o $(DIST_DIR)/tailwind.css $(ASSETS_DIR)/vendors/tailwind.css

test:
	go test -v -p 1 ./...
