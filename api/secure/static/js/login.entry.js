import FormState from './formState.js';

const form = {
    id: $('#frmLogin'),
    identity: $('#txtIdentity'),
    password: $('#txtPassword'),
    loginButton: $('#btnLogin'),
    registerButton: $('#btnRegister')
};

var fs = {};

$(document).ready(() => {
    fs = new FormState(form.loginButton);
    fs.submitDisabled(true);

    localStorage.setItem('return', getParameterByName("return"));

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

    $(document).on('keyup', pressEnter);
}

function pressEnter(e) {
    if (e.key !== 'Enter')
        return;

    tryLogin();
    e.preventDefault();
}

function tryLogin() {
    form.id.validator('validate');

    if (fs.isFormValid()) {
        submitLogin();
    }
}

function gotoRegister() {
    window.location.replace('/v1/register');
}

function getApp() {
    let appUrl = localStorage.getItem('return');
    let ip = localStorage.getItem('ip');
    let location = localStorage.getItem('location');

    let result = {
        Name: appUrl,
        IP: ip,
        Location: location,
        InstanceID: instanceID
    };

    return result;
}

function submitLogin() {
    fs.submitDisabled(true);

    $.ajax({
        url: "/v1/login",
        type: "POST",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify({
            App: getApp(),
            Email: form.identity.val(),
            Password: form.password.val()
        }),
        cache: false,
        success: function (res) {
            console.info(res);
            afterLogin(res.Data);
        },
        error: function (res) {
            // Fail message
            $('#success').html("<div class='alert alert-danger'>");
            $('#success > .alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
                .append("</button>");
            $('#success > .alert-danger').append($("<strong>").text(res.Error));
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
        navigator.geolocation.getCurrentPosition(setLocation);
    }
}

function setLocation(position) {
    let location = position.coords.latitude + ", " + position.coords.longitude;
    localStorage.setItem('location', location);
}

function getIP() {
    $.getJSON('//jsonip.com/?callback=?', function (data) {
        localStorage.setItem('ip', data.ip);
    });
}

function afterLogin(sessionID) {
    let finalURL = localStorage.getItem('return') || 'https://avosa.co.za';
    finalURL += "?token=" + sessionID

    window.location.replace(finalURL);
}

function getParameterByName(name, url) {
    if (!url)
        url = window.location.href;

    name = name.replace(/[\[\]]/g, "\\$&");

    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)");
    const results = regex.exec(url);

    if (!results)
        return null;

    if (!results[2])
        return '';

    return decodeURIComponent(results[2].replace(/\+/g, " "));
}