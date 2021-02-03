{{ if .flash }}
  {{ .flash }}
{{ end }}
<form class="" action="/users" method="post">
  {{ .Form | renderform }}
  <br>
  <input type="submit" value="Submit" />
</form>
