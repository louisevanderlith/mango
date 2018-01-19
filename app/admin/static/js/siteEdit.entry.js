
import FormState from './formState.js';
import * as lookup from './pathLookup';

const form = {
    panel: $('#pnlEdit'),
    id: $('#frmSite'),
    title: $('#txtTitle'),
    description: $('#txtDescription'),
    email: $('#txtEmail'),
    phone: $('#txtPhone'),
    url: $('#txtURL'),
    image: $('#uplImage'),
    styleSheet: $('#txtStylesheet'),
    socialLinks: $('#lstSocial'),
    portfolio: $('#lstPortfolio'),
    aboutSections: $('#lstAbout'),
    saveButton: $('#btnSave'),
    editButton: $('#btnEdit'),
    addSocialButton: $('#btnAddSocial'),
    addPortfolioButton: $('#btnAddPortfolio'),
    addParagraphButton: $('#btnAddParagraph'),
    imageHolder: $('#imageHolder')
};

var fs = {};
var currentID = 0;
var imageURL = '';
var uploadStore = {};

$(document).ready(() => {
    fs = new FormState(form.saveButton);
    fs.submitDisabled(true);

    getImageURL();

    registerEvents();
});

function registerEvents() {
    form.saveButton.on('click', trySave);

    let validForm = form.id.validator();
    validForm.on('invalid.bs.validator', fs.onValidate);
    validForm.on('valid.bs.validator', fs.onValidate);

    form.addParagraphButton.on('click', addParagraphRow);
    form.addSocialButton.on('click', addSocialRow);
    form.addPortfolioButton.on('click', addPortfolioRow);

    const body = $('body');
    body.on('click', '.removeRow', removeRow);
    body.on('change', 'input[type="file"]', uploadFile);
}

function getImageURL() {
    lookup.buildPath('Artifact.API', "upload", ["file"]).then((url) => {
        imageURL = url;
    }).catch((err) => {
        console.error('Image URL Error: ', err.Error);
    });
}

function uploadFile(e) {
    const fileElem = e.target;
    const files = fileElem.files;

    const fileData = $(fileElem);
    const forAttr = fileData.data('for');
    const nameAttr = fileData.data('name');
    const idAttr = fileData.data('id');
    const infoObj = {
        For: forAttr,
        ItemName: nameAttr,
        ItemID: idAttr
    };

    if (files.length > 0) {
        let formData = new FormData();
        formData.append('file', files[0]);
        formData.append('info', JSON.stringify(infoObj));

        doUpload(formData, infoObj);
    }
}

function doUpload(formData, infoObj) {
    lookup.buildPath('Artifact.API', 'upload').then((buildPath) => {
        $.ajax({
            url: buildPath,
            type: 'POST',
            contentType: false,
            processData: false,
            data: formData,
            success: (obj) => {
                finishUpload(obj, infoObj)
            },
            error: function (obj) {
                $('#success').html("<div class='alert alert-danger'>");
                $('#success > .alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
                    .append("</button>");
                $('#success > .alert-danger').append($("<strong>").text(obj.Error));
                $('#success > .alert-danger').append('</div>');
            }
        });
    });
}

function finishUpload(obj, infoObj) {
    const key = `${infoObj.ItemName}_${infoObj.ItemID}`;
    uploadStore[key] = obj.Data;
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
            ImageID: form.image.val(),
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

function setList(elem, data, htmlFunc) {
    const dataLen = data.length;
    let items = [];

    for (let i = 0; i < dataLen; i++) {
        let row = htmlFunc(i, data[i]);
        items.push(row);
    }

    elem.append(items.join(''));
}

function removeRow(e) {
    let confirmed = confirm("Are you sure you want to remove this item?");

    if (confirmed)
        $(e.target.parentNode).remove();
}

function addSocialRow(obj) {
    let id = form.socialLinks.children.length + 1;
    let html = socialRowHTML(id, obj);

    form.socialLinks.append(html);
}

function addPortfolioRow(obj) {
    let id = form.portfolio.children.length + 1;
    let html = portfolioRowHTML(id, obj);

    form.portfolio.append(html);
}

function addParagraphRow(obj) {
    let id = form.aboutSections.children.length + 1;
    let html = paragraphRowHTML(id, obj);

    form.aboutSections.append(html);
}