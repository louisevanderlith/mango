import FormState from './formState.js';
import * as lookup from './pathLookup';
import { buildPath } from './pathLookup';

const form = {
    panel: $('#pnlEdit'),
    id: $('#frmSite'),
    title: $('#txtTitle'),
    description: $('#txtDescription'),
    email: $('#txtEmail'),
    phone: $('#txtPhone'),
    url: $('#txtURL'),
    image: $('#uplImage'),
    imageID: 0,
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

$(document).ready(() => {
    fs = new FormState(form.saveButton);
    fs.submitDisabled(true);

    form.panel.hide();

    getImageURL();

    registerEvents();
});

function getImageURL() {
    lookup.buildPath('Artifact.API', "upload", ["file"]).then((url) => {
        imageURL = url;
    }).catch((err) => {
        alert('Image URL Error: ', err.Error);
    });
}

function registerEvents() {
    form.saveButton.on('click', trySave);
    form.editButton.on('click', edit);

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

function edit(e) {
    currentID = $(e.target).attr('data-rowid');

    if (currentID > 0) {
        lookup.buildPath('Folio.API', "site", [currentID]).then((buildPath) => {
            $.ajax({
                url: buildPath,
                type: "GET",
                contentType: "application/json; charset=utf-8",
                cache: false,
                success: fillForm,
                error: function (obj) {
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
    form.styleSheet.val(data.StyleSheet);

    setImageHolder(data.ImageID);

    setList(form.socialLinks, data.SocialLinks, socialRowHTML);
    setList(form.portfolio, data.PortfolioItems, portfolioRowHTML);
    setList(form.aboutSections, data.AboutSections, paragraphRowHTML);

    form.panel.show();
}

function setImageHolder(imageID) {
    let imageElem = `<input class="form-control" type="file" multiple="false" accept=".jpg, .jpeg, .png" id="uplImage" placeholder="Site Logo" required data-validation-required-message="Please provide an image." />`;

    if (imageID) {
        const imageSrc = `${imageURL}/${imageID}`;
        imageElem = `<input type="image" id="uplImage" class="form-control" src="${imageSrc}" alt="portfolio image" />`
    }

    form.imageHolder.html(imageElem);
}

function uploadFile(e) {
    const fileElem = e.target;
    const files = fileElem.files;

    if (files.length > 0) {
        let formData = new FormData();
        formData.append('file', files[0]);
        formData.append('info', JSON.stringify({
            For: 'logo',
            ItemName: 'Profile',
            ItemID: 99
        }));

        doUpload(formData);
    }
}

function doUpload(formData) {
    lookup.buildPath('Artifact.API', 'upload').then((buildPath) => {
        $.ajax({
            url: buildPath,
            type: 'POST',
            contentType: false,
            processData: false,
            data: formData,
            success: finishUpload,
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

function finishUpload(obj) {
    console.log(obj.Data);
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

function socialRowHTML(id, obj) {
    const result = `<li class="list-group-item" id="liSocial${id}">
        <button type="button" class="fa fa-close close removeRow" data-liID="liSocial${id}"></button>
        <div class="control-group">
            <div class="form-group floating-label-form-group controls">
                <label for="txtSocialIcon${id}">Icon Name <i class="fa ${obj.Icon}"></i></label>
                <input class="form-control" type="text" id="txtSocialIcon${id}" required data-validation-required-message="Please provide the icon's name. See font-awesome for options." value="${obj.Icon}" />
                <p class="help-block with-errors text-danger"></p>
            </div>
            <div class="form-group floating-label-form-group controls">
                <label for="txtSocialURL${id}">URL</label>
                <input class="form-control" type="url" id="txtSocialURL${id}" required data-validation-required-message="Please enter the social URL." value="${obj.URL}" />
                <p class="help-block with-errors text-danger"></p>
            </div>
        </div>
    </li>`;

    return result;
}

function portfolioRowHTML(id, obj) {
    let imgCtrl = `<input class="form-control" type="file" multiple="false" accept=".jpg, .jpeg, .png" id="uplPortfolioImg${id}" placeholder="Portfolio Image" required data-validation-required-message="Please provide this item's image." />`;

    if (obj.ImageID) {
        const imageSrc = `${imageURL}/${obj.ImageID}`;
        imgCtrl = `<input type="image" id="uplPortfolioImg${id}" class="form-control" src="${imageSrc}" alt="portfolio image" />`;
    }

    const result = `<li class="list-group-item" id="liPortfolio${id}">
        <button type="button" class="fa fa-close close removeRow" data-liID="liPortfolio${id}"></button>
        <div class="control-group">
            <div class="form-group floating-label-form-group controls">
                <label for="uplPortfolioImg${id}">Icon</label>
                ${imgCtrl}
                <p class="help-block with-errors text-danger"></p>
            </div>
            <div class="form-group floating-label-form-group controls">
                <label for="txtPortfolioName${id}">Name</label>
                <input type="text" class="form-control" id="txtPortfolioName${id}" required placeholder="Name" data-validation-required-message="Please enter a name." value="${obj.Name}"/>
                <p class="help-block with-errors text-danger"></p>
            </div>
            <div class="form-group floating-label-form-group controls">
                <label for="txtPortfolioURL${id}">URL</label>
                <input class="form-control" type="url" id="txtPortfolioURL${id}" required placeholder="URL" data-validation-required-message="Please enter the destination URL." value="${obj.URL}" />
            </div>
        </div>
    </li>`;

    return result;
}

function paragraphRowHTML(id, obj) {
    const result = `<li class="list-group-item" id="liAbout${id}">
        <button type="button" class="fa fa-close close removeRow" data-liID="liAbout${id}"></button>
        <div class="control-group">
            <div class="form-group floating-label-form-group controls">
                <label for="txtAboutParagraph${id}">Paragraph</label>
                <textarea class="form-control" id="txtAboutParagraph${id}" cols="40" rows="5" required data-validation-required-message="Please provide this paragraphs's text." placeholder="About Paragraph">${obj.SectionText}</textarea>
                <p class="help-block with-errors text-danger"></p>
            </div>
        </div>
    </li>`;

    return result;
}