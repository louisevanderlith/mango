import FormState from './formState.js';

const form = {
    id: $('#frmRegister'),
    name: $('#txtName'),
    contact: $('#txtContact'),
    email: $('#txtEmail'),
    password: $('#txtPassword'),
    confirmPass: $('#txtConfirmPass'),
    registerButton: $('#btnRegister')
};

var fs = {};

$(document).ready(() => {
    fs = new FormState(form.registerButton);
    fs.submitDisabled(true);

    registerEvents();
});

function registerEvents() {
    form.registerButton.on('click', tryRegister);

    let validForm = form.id.validator();
    validForm.on('invalid.bs.validator', fs.onValidate);
    validForm.on('valid.bs.validator', fs.onValidate);

    $(document).on('keyup', pressEnter);
}

function pressEnter(e) {
    if (e.key !== 'Enter')
        return;

    tryRegister();
    e.preventDefault();
}

function tryRegister() {
    form.id.validator('validate');

    if (fs.isFormValid()) {
        submitRegister();
    }
}

function getApp() { 
    let appUrl = localStorage.getItem('return');
    let ip = localStorage.getItem('ip');
    let location = localStorage.getItem('location');

    let result = {
        Name: appUrl,
        IP: ip,     
        Location: location, 
        InstanceID: instanceID   
    };

    return result;
}

function submitRegister() {
    fs.submitDisabled(true);

    $.ajax({
        url: "/v1/register",
        type: "POST",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify({
            App: getApp(),
            Name: form.name.val(),
            Email: form.email.val(),
            ContactNumber: form.contact.val(),
            Password: form.password.val(),
            PasswordRepeat: form.confirmPass.val()
        }),
        cache: false,
        success: function (res) {
            // Success message
            $('#success').html("<div class='alert alert-success'>");
            $('#success > .alert-success').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
                .append("</button>");
            $('#success > .alert-success')
                .append("<strong>"+res.Data+"</strong>");
            $('#success > .alert-success')
                .append('</div>');
            //clear all fields
            form.id.trigger("reset");
        },
        error: function (res) {
            // Fail message
            $('#success').html("<div class='alert alert-danger'>");
            $('#success > .alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
                .append("</button>");
            $('#success > .alert-danger').append($("<strong>").text(res.Error));
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