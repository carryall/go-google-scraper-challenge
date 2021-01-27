{{ template "shared/_alert.tpl" . }}

<form class="form form-registration" action="/users" method="post">
  <div>
    <div class="form__input-group">
      <label for="email" class="sr-only">Email</label>
      <input id="email" name="email" type="email" autocomplete="email" required class="form__input" placeholder="Email">
    </div>
    <div class="form__input-group">
      <label for="password" class="sr-only">Password</label>
      <input id="password" name="password" type="password" autocomplete="current-password" required class="form__input" placeholder="Password">
    </div>
    <div class="form__input-group">
      <label for="password_confirmation" class="sr-only">Password</label>
      <input id="password_confirmation" name="password_confirmation" type="password" required class="form__input" placeholder="Password Confirmation">
    </div>
  </div>
  <input type="submit" value="Submit" class="form__submit-button"/>
</form>
