$(document).ready(() => {
    verifyLogin();
});

function verifyLogin(){
    var avoToken = getParameterByName('avotoken');
    var hasToken = localStorage.getItem('avotoken') != null;

    if(!hasToken && avoToken) {
        localStorage.setItem('avotoken', avoToken);
    } else {
        let loginURL = 'http://secure.localhost/v1/login';
        window.location.replace(loginURL);
    }
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