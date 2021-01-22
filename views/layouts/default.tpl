<!DOCTYPE html>

<html>
<head>
  <title>Google Scraper</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  {{ assets_css "css/application.css" }}
  {{ assets_css "css/tailwind.css" }}
</head>
<body class="{{.ControllerName}} {{.ActionName}}">
  <div class="icon-sprite">
    {{ render_file "static/symbol/svg/sprite.symbol.svg" }}
  </div>

  <div class="app-container">
    <div class="app-header">
      <h2 class="app-header__title">
        {{ .Title }}
      </h2>
    </div>
    <div class="app-content">
      {{ .LayoutContent }}
    </div>
  </div>
</body>
</html>
