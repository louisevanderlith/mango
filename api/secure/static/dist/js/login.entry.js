this.login = this.login || {};
this.login.entry = this.login.entry || {};
(function (exports) {
'use strict';

var classCallCheck = function (instance, Constructor) {
  if (!(instance instanceof Constructor)) {
    throw new TypeError("Cannot call a class as a function");
  }
};

var createClass = function () {
  function defineProperties(target, props) {
    for (var i = 0; i < props.length; i++) {
      var descriptor = props[i];
      descriptor.enumerable = descriptor.enumerable || false;
      descriptor.configurable = true;
      if ("value" in descriptor) descriptor.writable = true;
      Object.defineProperty(target, descriptor.key, descriptor);
    }
  }

  return function (Constructor, protoProps, staticProps) {
    if (protoProps) defineProperties(Constructor.prototype, protoProps);
    if (staticProps) defineProperties(Constructor, staticProps);
    return Constructor;
  };
}();

var _stateControl = {};
var _submitButton = {};

var _onValid = function _onValid(targetName) {
    _stateControl[targetName] = {
        valid: true,
        errors: []
    };
};

var _onInvalid = function _onInvalid(targetName, errors) {
    _stateControl[targetName] = { valid: false, errors: errors };
    _submitDisabled(true);
};

var _submitDisabled = function _submitDisabled(disable) {
    _submitButton.prop('disabled', disable);
};

var _formValid = function _formValid() {
    var result = true;
    var stateControlKeys = Object.keys(_stateControl);
    var keysLen = stateControlKeys.length;

    for (var i = 0; i < keysLen; i++) {
        var currKey = stateControlKeys[i];
        var ctrl = _stateControl[currKey];

        if (!ctrl.valid) {
            result = false;
            break;
        }
    }

    return result;
};

var FormState = function () {
    function FormState(submitButton) {
        classCallCheck(this, FormState);

        _submitButton = submitButton;
    }

    createClass(FormState, [{
        key: "onValidate",
        value: function onValidate(event) {
            var isValid = event.type === "valid";
            var targetName = event.relatedTarget.id;

            if (isValid) {
                _onValid(targetName);
            } else {
                var errors = event.detail;
                _onInvalid(targetName, errors);
            }

            if (_formValid()) {
                _submitDisabled(false);
            }
        }
    }, {
        key: "isFormValid",
        value: function isFormValid() {
            return _formValid();
        }
    }, {
        key: "submitDisabled",
        value: function submitDisabled(disable) {
            _submitDisabled(disable);
        }
    }]);
    return FormState;
}();

var LoginForm = function LoginForm() {
    classCallCheck(this, LoginForm);

    console.log('LOGIN FORM INSTANCE');

    $(document).ready(function () {
        console.log('DOCUMENT IS READY!!');
        fs = new FormState(form.loginButton);
        fs.submitDisabled(true);

        var avoToken = localStorage.getItem('avotoken');
        returnURL = document.referrer; //getParameterByName('returnURL');

        if (!avoToken) {
            registerEvents();
            getLocation();
            getIP();
        } else {
            afterLogin(avoToken);
        }
    });
};

var lgnForm = new LoginForm();

var form = {
    id: $('#frmLogin'),
    identity: $('#txtIdentity'),
    password: $('#txtPassword'),
    loginButton: $('#btnLogin'),
    registerButton: $('#btnRegister')
};

var fs = {};
var returnURL = '';
var location = '';
var ip = '';

function registerEvents() {
    form.loginButton.on('click', tryLogin);
    form.registerButton.on('click', gotoRegister);

    var validForm = form.id.validator();
    validForm.on('invalid.bs.validator', fs.onValidate);
    validForm.on('valid.bs.validator', fs.onValidate);
}

function tryLogin(e) {
    form.id.validator('validate');

    if (fs.isFormValid()) {
        submitLogin();
    }
}

function gotoRegister() {
    window.location.replace('/v1/register');
}

function submitLogin() {
    fs.submitDisabled(true);

    $.ajax({
        url: "/v1/login",
        type: "POST",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify({
            Identifier: form.identity.val(),
            Password: form.password.val(),
            IP: ip,
            Location: location,
            ReturnURL: returnURL
        }),
        cache: false,
        success: function success(result) {
            localStorage.setItem('avotoken', result);
            afterLogin(result);

            //clear all fields
            form.id.trigger("reset");
        },
        error: function error(err) {
            console.error(err);
            // Fail message
            $('#success').html("<div class='alert alert-danger'>");
            $('#success > .alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;").append("</button>");
            $('#success > .alert-danger').append($("<strong>").text("Sorry, it seems something went wrong. Please try again."));
            $('#success > .alert-danger').append('</div>');
            //clear all fields
            form.id.trigger("reset");
        },
        complete: function complete() {
            setTimeout(function () {
                fs.submitDisabled(false);
            }, 1000);
        }
    });
}

function getLocation() {
    if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(setPosition);
    }
}

function setPosition(position) {
    location = position.coords.latitude + ", " + position.coords.longitude;
}

function getIP() {
    $.getJSON('//jsonip.com/?callback=?', function (data) {
        ip = data.ip;
    });
}

function afterLogin(token) {
    var finalURL = returnURL || 'http://www.localhost/';

    finalURL += '?avotoken=' + token;
    window.location.replace(finalURL);
}

exports.lgnForm = lgnForm;

}((this.login.entry.js = this.login.entry.js || {})));
