import * as services from './services'

$(document).ready(() => {
    registerEvents();
});

function registerEvents() {
    $('#btnAddSite').on('click', addSite);
}

function addSite() {
    let data = {
        Title: "New Site",
        Description: "Newly created website",
        ContactEmail: "fake@email.com",
        ContactPhone: "000",
        URL: "https://avosa.co.za",
        StyleSheet: "none.css",
    };

    let success = function (obj) {
        location.href = "/site/" + obj.Data;
    }

    services.createSite(data, success);
}