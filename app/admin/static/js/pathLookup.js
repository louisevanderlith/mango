let pastNames = {
    "Router.API": "https://router.avosa.co.za/v1/discovery/"
};

async function getRouterPath(serviceName) {
    const routerURL = await getServiceURL("Router.API");

    return `${routerURL}${instanceKey}/${serviceName}/true`;
}

function doLookup(serviceName) {
    return new Promise(async (resolve) => {
        $.ajax({
            url: await getRouterPath(serviceName),
            type: "GET",
            contentType: "application/json; charset=utf-8",
            cache: true,
            success: resolve,
            error: function (err) {
                console.error(err);
            }
        });
    });
}

export async function getServiceURL(serviceName) {
    return pastNames[serviceName] || await doLookup(serviceName);
}

export async function buildPath(serviceName, controller, params) {
    let url = await getServiceURL(serviceName);
    let result = `${url}v1/${controller}`;

    if (params) {
        for (let i = 0; i < params.length; i++) {
            result += "/" + params[i];
        }
    }

    return result;
}