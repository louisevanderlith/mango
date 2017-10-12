var form = {
    identity: $('#txtIdentity'),
    password: $('#txtPassword')
};

var returnURL = '';
var location = '';
var ip = '';

$(document).ready(() => {
    var avoToken = localStorage.getItem('avotoken');
    returnURL = document.referrer;//getParameterByName('returnURL');
    console.log(returnURL);

    if (!avoToken) {
        getLocation();
        getIP();
    } else {
        afterLogin(avoToken);
    }
});

function submitLogin() {
    $('#btnLogin').prop('disabled', true);

    var obj = {
        Identifier: form.identity.val(),
        Password: form.password.val(),
        IP: 'localhost',
        Location: location,
        ReturnURL: returnURL
    };
}

function postMessage(obj) {
    $.ajax({
        url: "/login",
        type: "POST",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify({
            "Body": message,
            "Email": email,
            "Name": name,
            "Phone": phone
        }),
        cache: false,
        success: function (result) {
            localStorage.setItem('avotoken', result);
            afterLogin(result);
        },
        error: function (err) {
            console.log(err);
        },
        complete: function () {
            setTimeout(function () {
                $('#btnLogin').prop('disabled', false);
            }, 1000);
        }
    });
}

function getParameterByName(name, url) {
    if (!url)
        url = window.location.href;

    name = name.replace(/[\[\]]/g, "\\$&");

    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)");
    var results = regex.exec(url);

    if (!results)
        return null;

    if (!results[2])
        return '';

    return decodeURIComponent(results[2].replace(/\+/g, " "));
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
    var finalURL = 'http://www.localhost/';

    // redirect user
    if (returnURL) {
        finalURL = returnURL;
    }

    finalURL += ('?avotoken=' + token);
    window.location.replace(finalURL);
}