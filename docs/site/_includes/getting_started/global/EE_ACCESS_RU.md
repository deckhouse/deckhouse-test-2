{% assign revision=include.revision %}

<div class="license-form__wrap">
<div class="license-form-enter">
<h3 class="text text_h3">
  Введите лицензионный ключ
</h3>

<div class="form form--inline">
  <div class="form__row" style="max-width: 383px;">
    <label class="label">
      Лицензионный ключ
    </label>
    <input id="license-token-input" class="textfield"
      type="text" license-token-{{ revision }} name="license-token-{{ revision }}"
      autocomplete="off" />
  </div>
  <a href="#" id="enter-license-key-{{ revision }}" class="button button_alt">Ввести</a>
  <span></span>
</div>
</div>

<script>
$(document).ready(function() {

    tokenInputElement{{ revision | replace: '-', '' }} = $('[license-token-{{ revision }}]');
    if ($.cookie("demotoken") || $.cookie("license-token")) {
        let token = $.cookie("license-token") ? $.cookie("license-token") : $.cookie("demotoken");
        tokenInputElement{{ revision | replace: '-', '' }}.val(token);
    }
})
</script>

<div class="license-form-request">
<h3 class="text text_h3">
  Нет ключа?
</h3>
<div class="button-group">
  <a href="" data-open-modal="request_access" class="button button_alt">Запросить бесплатный триал</a>
</div>
</div>
</div>
