
import FormState from './formState.js';
import * as lookup from './pathLookup';
import * as services from './services';

const form = {
    panel: $('#pnlEdit'),
    id: $('#frmSite'),
    title: $('#txtTitle'),
    description: $('#txtDescription'),
    email: $('#txtEmail'),
    phone: $('#txtPhone'),
    url: $('#txtURL'),
    image: $('#uplProfileImg'),
    styleSheet: $('#txtStylesheet'),
    socialLinks: $('#lstSocial'),
    portfolio: $('#lstPortfolio'),
    aboutSections: $('#lstAbout'),
    headers: $('#lstHeader'),
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

$(document).ready(() => {
    fs = new FormState(form.saveButton);
    fs.submitDisabled(true);

    let path = window.location.pathname;
    currentID = parseInt(path.substring(path.lastIndexOf('/') + 1));

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
    const idAttr = fileData.data('itemid');
    const ctrlID = fileElem.id;
    const infoObj = {
        For: forAttr,
        ItemName: nameAttr,
        ItemID: idAttr
    };

    if (files.length > 0) {
        let formData = new FormData();
        formData.append('file', files[0]);
        formData.append('info', JSON.stringify(infoObj));

        doUpload(formData, infoObj, ctrlID);
    }
}

function doUpload(formData, infoObj, ctrlID) {
    let success = function (obj) {
        finishUpload(obj, infoObj, ctrlID);
    }

    let error = function (obj) {
        $('#success').html("<div class='alert alert-danger'>");
        $('#success > .alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
            .append("</button>");
        $('#success > .alert-danger').append($("<strong>").text(obj.Error));
        $('#success > .alert-danger').append('</div>');
    }

    services.createUpload(formData, success, error);
}

function finishUpload(obj, infoObj, ctrlID) {
    let fullURL = imageURL + "/" + obj.Data;
    let imageHolder = $(`#${ctrlID.replace('Img', 'View')}`);
    let uploader = $('#' + ctrlID);

    imageHolder.removeAttr('hidden');
    imageHolder.attr("src", fullURL);

    uploader.data("id", obj.Data);
    uploader.removeAttr('required');
}

function trySave() {
    form.id.validator('validate');

    if (fs.isFormValid()) {
        submitSite();
    }
}

function submitSite() {
    fs.submitDisabled(true);

    let data = {
        Id: currentID,
        Title: form.title.val(),
        Description: form.description.val(),
        ContactEmail: form.email.val(),
        ContactPhone: form.phone.val(),
        URL: form.url.val(),
        ImageID: form.image.data('id'),
        StyleSheet: form.styleSheet.val(),
        SocialLinks: getList(form.socialLinks, "Social"),
        PortfolioItems: getList(form.portfolio, "Portfolio"),
        AboutSections: getList(form.aboutSections, "About"),
        Headers: getList(form.headers, "Header")
    };

    let success = function (data) {
        $('#success').html("<div class='alert alert-success'>");
        $('#success > .alert-success').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
            .append("</button>");
        $('#success > .alert-success').append($("<strong>").text(data.Data));
        $('#success > .alert-success').append('</div>');
        //clear all fields
        form.id.trigger("reset");
    };

    let fail = function (data) {
        // Fail message
        $('#success').html("<div class='alert alert-danger'>");
        $('#success > .alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
            .append("</button>");
        $('#success > .alert-danger').append($("<strong>").text(data.Error));
        $('#success > .alert-danger').append('</div>');
        //clear all fields
        form.id.trigger("reset");
    };

    let complete = function () {
        setTimeout(function () {
            fs.submitDisabled(false);
        }, 1000);
    };

    services.updateSite(data, success, fail, complete);
}

function getList(elem, container) {
    let result = [];
    const children = elem.children();
    const childLen = children.length;

    for (let i = 0; i < childLen; i++) {
        let child = children[i];
        let info = getContainerInfo(child, container);

        result.push(info);
    }

    return result;
}

function getContainerInfo(child, container) {
    let result = "";

    switch (container) {
        case "Social":
            result = getSocialInfo(child);
            break;
        case "Portfolio":
            result = getPortfolioInfo(child);
            break;
        case "About":
            result = getAboutInfo(child);
            break;
        case "Header":
            result = getHeaderInfo(child);
            break;
    }

    return result;
}

function getSocialInfo(child) {
    let id = child.id.replace('liSocial', '');
    let recordID = $(child).data('id');
    let icon = $(`#txtSocialIcon${id}`).val();
    let url = $(`#txtSocialURL${id}`).val();

    return { ID: recordID, Icon: icon, URL: url };
}

function getPortfolioInfo(child) {
    let id = child.id.replace('liPortfolio', '');
    let recordID = $(`#uplPortfolioImg${id}`).data('itemid');
    let imageID = $(`#uplPortfolioImg${id}`).data('id');
    let name = $(`#txtPortfolioName${id}`).val();
    let url = $(`#txtPortfolioURL${id}`).val();

    return { ID: recordID, ImageID: imageID, Name: name, URL: url };
}

function getAboutInfo(child) {
    let id = child.id.replace('liAbout', '');
    let recordID = $(child).data('id');
    let paragraph = $(`#txtAboutParagraph${id}`).val();

    return { ID: recordID, SectionText: paragraph };
}

function getHeaderInfo(child) {
    let id = child.id.replace('xxx', '');
    let recordID = $(child).data('id');
    let imageID = $(`#yyy${id}`).data('itemid');
    let text = $(`#xxx${id}`).val();

    return { ID: recordID, Text: text, ImageID: imageID };
}

function removeRow(e) {
    let confirmed = confirm("Are you sure you want to remove this item?");

    if (confirmed)
        $(e.target.parentNode).remove();
}

function addSocialRow(obj) {
    let data = {
        Icon: 'fa-ban',
        URL: 'none',
        Profile: {
            ID: currentID
        }
    };

    let success = function (data) {
        location.reload();
    };

    let fail = function (data) {
        console.error(data.Error);
    }

    services.createSocialLink(data, success, fail);
}

function addPortfolioRow(obj) {
    let data = {
        ImageID: 0,
        URL: 'https://avosa.co.za',
        Name: 'avosa',
        Profile: {
            ID: currentID
        }
    };

    let success = function (data) {
        location.reload();
    };

    let fail = function (data) {
        console.error(data.Error);
    }

    services.createPortfolioItem(data, success, fail);
}

function addParagraphRow(obj) {
    let data = {
        SectionText: 'new section',
        Profile: {
            ID: currentID
        }
    };

    let success = function (data) {
        location.reload();
    };

    services.createAboutSection(data, success);
}

function addHeaderItem(obj) {
    let data = {
        Text: 'blank header',
        Profile: {
            ID: currentID
        }
    };

    let success = function (data) {
        location.reload();
    };

    services.createHeaderItem(data, success);
}