<!DOCTYPE html>

<html class="layout-default">
<head>
  <title>Google Scraper</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  {{ assets_css "/css/application.css" }}
</head>
<body class="{{.ControllerName}} {{.ActionName}}">
  <div class="icon-sprite" hidden="true">
    {{ render_file "static/symbol/svg/sprite.symbol.svg" }}
  </div>

  <div class="app-content">
    {{ .LayoutContent }}
  </div>
</body>
</html>
