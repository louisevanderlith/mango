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
    styleSheet: $('#txtStylesheet'),
    socialLinks: $('#lstSocial'),
    portfolio: $('#lstPortfolio'),
    aboutSections: $('#lstAbout'),
    saveButton: $('#btnSave'),
    editButton: $('#btnEdit'),
    removeSocial: $('.removeSocial'),
    removePortfolio: $('.removePortfolio'),
    removeParagraph: $('.removeParagraph'),
    addSocialButton: $('#btnAddSocial'),
    addPortfolioButton: $('#btnAddPortfolio'),
    addParagraphButton: $('#btnAddParagraph')
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

    form.removeParagraph.on('click', removeRow);
    form.removePortfolio.on('click', removeRow);
    form.removeSocial.on('click', removeRow);

    form.addParagraphButton.on('click', addParagraphRow);
    form.addSocialButton.on('click', addSocialRow);
    form.addPortfolioButton.on('click', addPortfolioRow);
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
    form.image.val(data.Image);
    form.styleSheet.val(data.StyleSheet);

    setList(form.socialLinks, data.SocialLinks, socialRowHTML);
    setList(form.portfolio, data.PortfolioItems, portfolioRowHTML);
    setList(form.aboutSections, data.AboutSections, paragraphRowHTML);

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
            Image: form.image.val(),
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
    //liID = $(e.target).attr('data-liID');
    //console.log(liID);

    e.target.parentNode.removeChild(e.target)
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
    console.log(obj);
    const result = `<li class="list-group-item" id="liSocial${id}">
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
        <button type="button" class="close float-right removeSocial" data-liID="liSocial${id}"></button>
    </li>`;

    return result;
}

function portfolioRowHTML(id, obj) {
    let imgCtrl = `<input class="form-control" type="file" accept=".jpg, .jpeg, .png" id="uplPortfolioImg${id}" placeholder="Portfolio Image" required data-validation-required-message="Please provide this item's image." />`;

    if (obj.Icon) {
        imgCtrl = `<input type="image" id="uplPortfolioImg${id}" class="form-control" src="${obj.Icon}" alt="portfolio image" />`
    }

    const result = `<li class="list-group-item" id="liPortfolio1">
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
        <button type="button" class="close float-right removePortfolio" data-liID="liPortfolio${id}"></button>
    </li>`;

    return result;
}

function paragraphRowHTML(id, obj) {
    const result = `<li class="list-group-item" id="liAbout${id}">
        <div class="control-group">
            <div class="form-group floating-label-form-group controls">
                <label for="txtAboutParagraph${id}">Paragraph</label>
                <textarea class="form-control" id="txtAboutParagraph${id}" cols="40" rows="5" required data-validation-required-message="Please provide this paragraphs's text." placeholder="About Paragraph">${obj.SectionText}</textarea>
                <p class="help-block with-errors text-danger"></p>
            </div>
        </div>
        <button type="button" class="close float-right removeParagraph" data-liID="liAbout${id}"></button>
    </li>`;

    return result;
}