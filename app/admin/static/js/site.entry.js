import FormState from './formState.js';
import * as lookup from './pathLookup';
import { buildPath } from './pathLookup';

const form = {
    panel: $('#pnlEdit'),
    id: $('#frmSite'),
    title: $('#txtTitle'),
    description: $('#txtDescription'),
    email: $('txtEmail'),
    phone: $('#txtPhone'),
    url: $('#txtURL'),
    imageURL: $('#txtImage'),
    styleSheet: $('#txtStylesheet'),
    socialLinks: $('#lstSocial'),
    portfolio: $('#lstPortfolio'),
    aboutSections: $('#lstAbout'),
    saveButton: $('#btnSave'),
    editButton: $('#btnEdit')
};

var fs = {};
var currentID = 0;

$(document).ready(() => {
    fs = new FormState(form.saveButton);
    fs.submitDisabled(true);

    form.panel.hide();

    registerEvents();
});

function registerEvents() {
    form.saveButton.on('click', trySave);
    form.editButton.on('click', edit)

    let validForm = form.id.validator();
    validForm.on('invalid.bs.validator', fs.onValidate);
    validForm.on('valid.bs.validator', fs.onValidate);
}

function edit(e) {
    currentID = $(e.target).attr('data-rowid');

    if (currentID > 0) {
        lookup.buildPath('Folio.API', "site", currentID).then((buildPath) => {
            $.ajax({
                url: buildPath,
                type: "GET",
                contentType: "application/json; charset=utf-8",
                cache: false,
                success: fillForm,
                error: function (obj) {
                    // Fail message
                    $('#success').html("<div class='alert alert-danger'>");
                    $('#success > .alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
                        .append("</button>");
                    $('#success > .alert-danger').append($("<strong>").text(obj.Error));
                    $('#success > .alert-danger').append('</div>');

                    form.panel.hide();
                }
            });
        });
    }
}

function fillForm(obj) {
    const data = obj.Data;
    
    form.title.val(data.Title);
    form.description.val(data.Description);
    form.email.val(data.ContactEmail);
    form.phone.val(data.ContactPhone);
    form.url.val(data.URL);
    form.imageURL.val(data.ImageURL);
    form.styleSheet.val(data.StyleSheet);

    setList(form.socialLinks, data.SocialLinks);
    setList(form.portfolio, data.PortfolioItems);
    setList(form.aboutSections, data.AboutSections);

    form.panel.show();
}

function trySave() {
    form.id.validator('validate');

    if (fs.isFormValid()) {
        submitSite();
    }
}

function submitSite() {
    fs.submitDisabled(true);

    $.ajax({
        url: lookup.buildPath('Folio.API', "site"),
        type: "POST",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify({
            Id: currentID,
            Title: form.title.val(),
            Description: form.description.val(),
            ContactEmail: form.email.val(),
            ContactPhone: form.phone.val(),
            URL: form.url.val(),
            ImageURL: form.imageURL.val(),
            StyleSheet: form.styleSheet.val(),
            SocialLinks: getList(form.socialLinks),
            PortfolioItems: getList(form.portfolio),
            AboutSections: getList(form.aboutSections)
        }),
        cache: false,
        success: function (data) {
            $('#success').html("<div class='alert alert-danger'>");
            $('#success > .alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
                .append("</button>");
            $('#success > .alert-danger').append($("<strong>").text(data));
            $('#success > .alert-danger').append('</div>');
            //clear all fields
            form.id.trigger("reset");
        },
        error: function (err) {
            // Fail message
            $('#success').html("<div class='alert alert-danger'>");
            $('#success > .alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
                .append("</button>");
            $('#success > .alert-danger').append($("<strong>").text(err));
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

function getList(elem) {
    let result = [];
    const children = elem.children;
    const childLen = children.length;

    for (let i = 0; i < childLen; i++) {
        let child = children[i];
        result.push(child.text);
    }

    return result;
}

function setList(elem, data) {
    const dataLen = data.length;
    let items = [];

    for (let i = 0; i < dataLen; i++) {
        let row = `<span class="list-group-item">${data[i]}</span>`;
        items.push(row);
    }

    elem.html(items.join(''));
}