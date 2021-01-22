{{ range $flashType, $flashMessage := .flash }}
  {{ if $flashMessage }}
    <div class="alert alert--{{$flashType}}">
      <div class="alert__icon">
        <svg class="icon" viewBox="0 0 20 20">
          <use xlink:href="svg/sprite.symbol.svg#{{$flashType}}" />
        </svg>
      </div>
      <div class="alert__body">
        <h3 class="alert__title">
          {{ $flashType | titlecase }}
        </h3>
        <p class="alert__message">
          {{ $flashMessage }}
        </p>
      </div>
    </div>
  {{ end }}
{{ end }}

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
