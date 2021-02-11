{{ template "shared/_alert.tpl" . }}

<form class="form form--rounded-input form-login" action="/sessions" method="post">
  <div>
    <div class="form__input-group">
      <label for="email" class="sr-only">Email</label>
      <input id="email" name="email" type="email" autocomplete="email" required class="form__input" placeholder="Email">
    </div>
    <div class="form__input-group">
      <label for="password" class="sr-only">Password</label>
      <input id="password" name="password" type="password" autocomplete="password" required class="form__input" placeholder="Password">
    </div>
  </div>
  <input type="submit" value="Submit" class="form__submit-button"/>
</form>
