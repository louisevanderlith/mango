$(document).ready(() => {
    verifyLogin();
});

function verifyLogin(){
    var hasToken = localStorage.getItem('avotoken') != null;

    if(!hasToken){
        var avoToken = getParameterByName('avotoken');

        if(avoToken){
            localStorage.setItem('avotoken', avoToken);
        }
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