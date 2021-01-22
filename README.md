# go-google-scraper-challenge
A project for Nimble Go Internal Certification on Web

[Staging](https://google-scraper-staging.herokuapp.com)
[Production](https://google-scraper-web.herokuapp.com)

## Development

### Create an ENV file

  Copy the `.env.example` file and rename it to `.env`, then set the `APP_RUN_MODE` to `dev`

### Build development dependencies

  ```sh
  make build-dependencies
  ```

### Compile assets files

  ```sh
  make assets
  ```

### Run Database service on Docker

  ```sh
  make db/setup
  ```

### Run migrations and the Go application for development

  ```sh
  make dev
  ```

  The application would be running locally at `http://localhost:8080`

### Run test

  ```sh
  make test
  ```

### SVG Icons

  The [SVG Sprite](https://github.com/jkphl/svg-sprite) is used on this project

  #### Add a new SVG file
  - put the new SVG file on `assets/images/icons` directory
  - install [SVG Sprite](https://github.com/jkphl/svg-sprite) globally
  ```sh
  npm install svg-sprite -g
  ```
  - generate the SVG sprite
  ```sh
  make assets/icon-sprite
  ```

  #### Use SVG inline
  ```html
  <svg class="icon" viewBox="0 0 20 20">
    <use xlink:href="svg/sprite.symbol.svg#{{ $iconName }}" />
  </svg>
  ```
