import FormState from './formState.js';

const form = {
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

$(document).ready(() => {
    fs = new FormState(form.loginButton);
    fs.submitDisabled(true);

    returnURL = document.referrer;

    registerEvents();
    getLocation();
    getIP();
});

function registerEvents() {
    form.loginButton.on('click', tryLogin);
    form.registerButton.on('click', gotoRegister);

    let validForm = form.id.validator();
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
        success: function () {
            //clear all fields
            form.id.trigger("reset");

            afterLogin();
        },
        error: function (err) {
            console.error(err);
            // Fail message
            $('#success').html("<div class='alert alert-danger'>");
            $('#success > .alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
                .append("</button>");
            $('#success > .alert-danger').append($("<strong>").text("Sorry, it seems something went wrong. Please try again."));
            $('#success > .alert-danger').append('</div>');
            //clear all fields
            form.id.trigger("reset");
        },
        complete: function () {
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

function afterLogin() {
    let finalURL = returnURL || 'http://www.localhost/';
    window.location.replace(finalURL);
}