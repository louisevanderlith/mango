
import FormState from './formState.js';
import * as lookup from './pathLookup';
import * as services from './services';

const tabAdder = {
    addSocialButton: $('#btnAddSocial'),
    addPortfolioButton: $('#btnAddPortfolio'),
    addParagraphButton: $('#btnAddParagraph'),
    addHeaderButton: $('#btnAddHeader')
};

const forms = {
    frmAbout: {
        id: $('#frmAbout'),
        lstAbout: $('#lstAbout'),
        save: $('#btnSaveAbout'),
        saveEvent: trySaveAbout
    },
    frmBasicSite: {
        id: $('#frmBasicSite'),
        txtTitle: $('#txtTitle'),
        txtDescription: $('#txtDescription'),
        txtEmail: $('#txtEmail'),
        txtPhone: $('#txtPhone'),
        txtURL: $('#txtURL'),
        image: $('#uplProfileImg'),
        txtStylesheet: $('#txtStylesheet'),
        save: $('#btnSaveSite'),
        saveEvent: trySaveSite
    },
    frmHeader: {
        id: $('#frmHeader'),
        lstHeader: $('#lstHeader'),
        save: $('#btnSaveHeader'),
        saveEvent: trySaveHeader
    },
    frmPortfolio: {
        id: $('#frmPortfolio'),
        lstPortfolio: $('#lstPortfolio'),
        save: $('#btnSavePortfolio'),
        saveEvent: trySavePortfolio
    },
    frmSocialmedia: {
        id: $('#frmSocialmedia'),
        lstSocial: $('#lstSocial'),
        save: $('#btnSaveSocial'),
        saveEvent: trySaveSocial
    }
};

var formStates = {};
var currentID = 0;
var imageURL = '';

$(document).ready(() => {
    readyForms();

    let path = window.location.pathname;
    currentID = parseInt(path.substring(path.lastIndexOf('/') + 1));

    getImageURL();
    registerEvents();
});

function readyForms() {
    let keys = Object.keys(forms);
    let keyLen = keys.length;

    for (let i = 0; i < keyLen; i++) {
        let key = keys[i];
        let item = forms[key];

        formStates[key] = new FormState(key, item.save);

        item.save.on('click', item.saveEvent);

        let validForm = item.id.validator();
        validForm.on('invalid.bs.validator', (event) => { formStates[key].onValidate(key, event); });
        validForm.on('valid.bs.validator', (event) => { formStates[key].onValidate(key, event); });
    }
}

function registerEvents() {
    tabAdder.addParagraphButton.on('click', addRow);
    tabAdder.addSocialButton.on('click', addRow);
    tabAdder.addPortfolioButton.on('click', addRow);
    tabAdder.addHeaderButton.on('click', addRow);

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

    services.createUpload(formData, success, errorMessage);
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

function successMessage(message) {
    $('#success').html("<div class='alert alert-success'>");
    $('#success > .alert-success').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
        .append("</button>");
    $('#success > .alert-success').append($("<strong>").text(message.Data));
    $('#success > .alert-success').append('</div>');
}

function errorMessage(message) {
    $('#success').html("<div class='alert alert-danger'>");
    $('#success > .alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
        .append("</button>");
    $('#success > .alert-danger').append($("<strong>").text(message.Error));
    $('#success > .alert-danger').append('</div>');
}

function completeForm(formKey, formState) {
    return () => {
        setTimeout(function () {
            formState.submitDisabled(formKey, false);
        }, 1000);
    }
}

function trySaveAbout() {
    const formKey = "frmAbout";
    forms.frmAbout.id.validator('validate');

    if (formStates[formKey].isFormValid(formKey)) {
        submitAbout(formKey);
    }
}

function submitAbout(formKey) {
    let formState = formStates[formKey];
    let sections = getList(forms.frmAbout.lstAbout, "About");

    formState.submitDisabled(formKey, true);

    for (let i = 0; i < sections.length; i++) {
        let data = sections[i];

        services.updateAboutSection(data, successMessage, errorMessage);
    }

    completeForm(formKey, formState)();
}

function trySaveSite() {
    const formKey = "frmBasicSite";
    forms.frmBasicSite.id.validator('validate');

    if (formStates[formKey].isFormValid(formKey)) {
        submitSite(formKey);
    }
}

function submitSite(formKey) {
    let siteForm = forms.frmBasicSite;
    let formState = formStates[formKey];
    formState.submitDisabled(formKey, true);

    let data = {
        Id: currentID,
        Title: siteForm.txtTitle.val(),
        Description: siteForm.txtDescription.val(),
        ContactEmail: siteForm.txtEmail.val(),
        ContactPhone: siteForm.txtPhone.val(),
        URL: siteForm.txtURL.val(),
        ImageID: siteForm.image.data('id'),
        StyleSheet: siteForm.txtStylesheet.val()
    };

    services.updateSite(data, successMessage, errorMessage, completeForm(formKey, formState));
}

function trySaveHeader() {
    const formKey = "frmHeader";
    forms.frmHeader.id.validator('validate');

    if (formStates[formKey].isFormValid(formKey)) {
        submitHeader(formKey);
    }
}

function submitHeader(formKey) {
    let formState = formStates[formKey];
    let sections = getList(forms.frmHeader.lstHeader, "Header");

    formState.submitDisabled(formKey, true);

    for (let i = 0; i < sections.length; i++) {
        let data = sections[i];

        services.updateHeaderItem(data, successMessage, errorMessage);
    }

    completeForm(formKey, formState)();
}

function trySavePortfolio() {
    const formKey = "frmPortfolio";
    forms.frmPortfolio.id.validator('validate');

    if (formStates[formKey].isFormValid(formKey)) {
        submitPortfolio(formKey);
    }
}

function submitPortfolio(formKey) {
    let formState = formStates[formKey];
    let sections = getList(forms.frmPortfolio.lstPortfolio, "Portfolio");

    formState.submitDisabled(formKey, true);

    for (let i = 0; i < sections.length; i++) {
        let data = sections[i];

        services.updatePortfolioItem(data, successMessage, errorMessage);
    }

    completeForm(formKey, formState)();
}

function trySaveSocial() {
    const formKey = "frmSocialmedia";
    forms.frmSocialmedia.id.validator('validate');

    if (formStates[formKey].isFormValid(formKey)) {
        submitSocial(formKey);
    }
}

function submitSocial(formKey) {
    let formState = formStates[formKey];
    let sections = getList(forms.frmSocialmedia.lstSocial, "Social");

    formState.submitDisabled(formKey, true);

    for (let i = 0; i < sections.length; i++) {
        let data = sections[i];

        services.updateSocialLink(data, successMessage, errorMessage);
    }

    completeForm(formKey, formState)();
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
    let id = child.id.replace('liHeader', '');
    let recordID = $(`#uplHeaderImg${id}`).data('itemid');
    let imageID = $(`#uplHeaderImg${id}`).data('id');
    let heading = $(`#txtHeaderHeading${id}`).val();
    let text = $(`#txtHeaderText${id}`).val();

    return { ID: recordID, ImageID: imageID, Heading: heading, Text: text };
}

function removeRow(e) {
    let confirmed = confirm("Are you sure you want to remove this item?");

    if (confirmed)
        $(e.target.parentNode).remove();
}

function addRow(e) {
    let confirmed = confirm("All unsaved data will be lost. Continue?");

    if (confirmed) {
        let type = e.target.id.replace('btnAdd', '');
        const funcs = {
            "Social": addSocialRow,
            "Portfolio": addPortfolioRow,
            "Paragraph": addParagraphRow,
            "Header": addHeaderRow
        };

        let rowFunc = funcs[type];
        rowFunc();
    }
}

function addSocialRow() {
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

    services.createSocialLink(data, success);
}

function addPortfolioRow() {
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

    services.createPortfolioItem(data, success);
}

function addParagraphRow() {
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

function addHeaderRow() {
    let data = {
        ImageID: 0,
        Heading: 'heading',
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