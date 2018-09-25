let pastNames = {
    "Router.API": "https://router.localhost/v1/discovery/"//"https://router.avosa.co.za/v1/discovery/"//
};

async function getRouterPath(serviceName) {
    const routerURL = await pathLookup.getServiceURL("Router.API");

    return `${routerURL}${instanceID}/${serviceName}/true`;
}

async function doLookup(serviceName) {
    const routerPath = await getRouterPath(serviceName);

    return new Promise((resolve) => {
        $.ajax({
            url: routerPath,
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

const pathLookup = {
    getServiceURL: async function (serviceName) {
        let serviceURL = pastNames[serviceName];

        if (!serviceURL) {
            serviceURL = await doLookup(serviceName);
        }

        return serviceURL;
    },
    buildPath: async function (serviceName, controller, params) {
        let url = await pathLookup.getServiceURL(serviceName);
        let result = url + 'v1/' + controller;

        if (params) {
            for (let i = 0; i < params.length; i++) {
                result += "/" + params[i];
            }
        }

        return result;
    }
};

export default pathLookup