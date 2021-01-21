<!DOCTYPE html>

<html>
<head>
  <title>Google Scraper</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  {{ assets_css "static/css/application.css" }}
  {{ assets_css "static/css/tailwind.css" }}
</head>

<body>
  <div class="app-container {{.ControllerName}} {{.ActionName}}">
    <header class="app-header">
      {{ .Title }}
    </header>
    <div class="app-content">
      {{ .LayoutContent }}
    </div>
  </div>
</body>
</html>
