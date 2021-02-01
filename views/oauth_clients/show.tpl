{{ template "shared/_alert.tpl" . }}

<form class="form form-oauth-client" disabled>
  <div class="form__input-group">
    <label for="clientID" class="form__label">Client ID</label>
    <input id="clientID" name="clientID" type="text" class="form__input" value="{{ .ClientID }}" disabled>
  </div>
  <div class="form__input-group">
    <label for="clientSecret" class="form__label">Client Secret</label>
    <input id="clientSecret" name="clientSecret" type="text" class="form__input" value="{{ .ClientSecret }}" disabled>
  </div>
</form>
