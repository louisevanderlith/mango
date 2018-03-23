import * as lookup from './pathLookup';

function defaultSuccess(data) {
    console.log(data);
}

function defaultFail(data) {
    console.error(data);
}

function defaultComplete() {
}

function doRequest(url, method, data, success, fail, complete) {
    $.ajax({
        url: url,
        type: method,
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify(data),
        cache: false,
        success: success || defaultSuccess,
        error: fail || defaultFail,
        complete: complete || defaultFail
    });
}

function doMultipartRequest(url, data, success, fail, complete) {
    $.ajax({
        url: url,
        type: 'POST',
        contentType: false,
        processData: false,
        data: data,
        success: success || defaultSuccess,
        error: fail || defaultFail,
        complete: complete || defaultComplete
    });
}

export function createSite(data, success, fail, complete) {
    lookup.buildPath('Folio.API', "site").then((buildPath) => {
        doRequest(buildPath, "POST", data, success, fail, complete);
    });
}

export function updateSite(data, success, fail, complete) {
    lookup.buildPath('Folio.API', "site").then((buildPath) => {
        doRequest(buildPath, "PUT", data, success, fail, complete);
    });
}

export function createUpload(data, success, fail, complete) {
    lookup.buildPath('Artifact.API', 'upload').then((buildPath) => {
        doMultipartRequest(buildPath, data, success, fail, complete);
    });
}

export function createSocialLink(data, success, fail, complete) {
    lookup.buildPath('Folio.API', 'social').then((buildPath) => {
        doRequest(buildPath, "POST", data, success, fail, complete);
    });
}

export function updateSocialLink(data, success, fail, complete) {
    lookup.buildPath('Folio.API', 'social').then((buildPath) => {
        doRequest(buildPath, "PUT", data, success, fail, complete);
    });
}

export function createAboutSection(data, success, fail, complete) {
    lookup.buildPath('Folio.API', 'about').then((buildPath) => {
        doRequest(buildPath, "POST", data, success, fail, complete);
    });
}

export function updateAboutSection(data, success, fail, complete) {
    lookup.buildPath('Folio.API', 'about').then((buildPath) => {
        doRequest(buildPath, 'PUT', data, success, fail, complete);
    });
}

export function createPortfolioItem(data, success, fail, complete) {
    lookup.buildPath('Folio.API', 'portfolio').then((buildPath) => {
        doRequest(buildPath, "POST", data, success, fail, complete);
    });
}

export function updatePortfolioItem(data, success, fail, complete) {
    lookup.buildPath('Folio.API', 'portfolio').then((buildPath) => {
        doRequest(buildPath, "PUT", data, success, fail, complete);
    });
}

export function createHeaderItem(data, success, fail, complete) {
    lookup.buildPath('Folio.API', 'header').then((buildPath) => {
        doRequest(buildPath, "POST", data, success, fail, complete);
    });
}

export function updateHeaderItem(data, success, fail, complete) {
    lookup.buildPath('Folio.API', 'header').then((buildPath) => {
        doRequest(buildPath, "PUT", data, success, fail, complete);
    });
}