import * as lookup from './pathLookup';

export function updateSite(data, success, fail, complete) {
    lookup.buildPath('Folio.API', "site").then((buildPath) => {
        $.ajax({
            url: buildPath,
            type: "PUT",
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(data),
            cache: false,
            success: success,
            error: fail,
            complete: complete
        });
    });
}

export function createUpload(data, success, fail, complete){
    lookup.buildPath('Artifact.API', 'upload').then((buildPath) => {
        $.ajax({
            url: buildPath,
            type: 'POST',
            contentType: false,
            processData: false,
            data: data,
            success: success,
            error: fail,
            complete: complete
        });
    });
}

export function createSocialLink(data, success, fail){
    lookup.buildPath('Folio.API', 'social').then((buildPath) => {
        $.ajax({
            url: buildPath,
            type: "POST",
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(data),
            cache: false,
            success: success,
            error: fail
        });
    });
}

export function createAboutSection(data, success, fail){
    lookup.buildPath('Folio.API', 'about').then((buildPath) => {
        $.ajax({
            url: buildPath,
            type: "POST",
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(data),
            cache: false,
            success: success,
            error: fail
        });
    });
}

export function createPortfolioItem(data, success, fail){
    lookup.buildPath('Folio.API', 'portfolio').then((buildPath) => {
        $.ajax({
            url: buildPath,
            type: "POST",
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify(data),
            cache: false,
            success: success,
            error: fail
        });
    });
}