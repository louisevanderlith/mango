let templates = {
    uploadHTML: _getUploadHTML(id),
    imageHTML: _getImageHTML(id),
    _getUploadHTML: function(id){
        return ``;
    },
    _getImageHTML: function(id){
        return ``;
    }
};

export function imageUploadControl(id, obj){
    let imgCtrl = `<input class="form-control" type="file" multiple="false" data-for="logo" data-name="Portfolio" data-id="${id}" accept=".jpg, .jpeg, .png" id="uplPortfolioImg${id}" placeholder="Portfolio Image" required data-validation-required-message="Please provide this item's image." />`;

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

export function socialMediaControl(id, obj){
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

export function htmlMultilineControl(id, obj){
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