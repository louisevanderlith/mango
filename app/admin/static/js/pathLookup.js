let pastNames = {
    "Router.API": "https://router.avosa.co.za/v1/discovery/"
};

async function getRouterPath(serviceName) {
    const routerURL = await this.getServiceURL("Router.API");

    return `${routerURL}${instanceKey}/${serviceName}`;
}

function doLookup(serviceName) {
    return new Promise((resolve) => {
        $.ajax({
            url: getRouterPath(serviceName),
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

export default {
    getServiceURL: async function (serviceName) {
        let serviceURL = pastNames[serviceName];

        if (!serviceURL) {
            serviceURL = await doLookup(serviceName);
        }

        return serviceURL
    }
}