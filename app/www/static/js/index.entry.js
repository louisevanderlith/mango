import FormState from './formState.js';
import config from './config.js';

const form = {
    id: $('#frmContact'),
    name: $('#txtName'),
    email: $('#txtEmail'),
    contact: $('#txtContact'),
    message: $('#txtMessage'),
    sendButton: $('#btnSend')
};

var fs = {};

$(document).ready(() => {
    fs = new FormState(form.sendButton);
    fs.submitDisabled(true);

    registerEvents();
});

function registerEvents() {
    form.sendButton.on('click', trySend);

    $("a[data-toggle=\"tab\"]").on('click', tabClick);

    let validForm = form.id.validator();
    validForm.on('invalid.bs.validator', fs.onValidate);
    validForm.on('valid.bs.validator', fs.onValidate);
}

function trySend() {
    form.id.validator('validate');

    if (fs.isFormValid()) {
        submitSend();
    }
}

function submitSend() {
    fs.submitDisabled(true);

    var firstName = form.name.val();

    if (firstName.indexOf(' ') >= 0) {
        firstName = name.split(' ').slice(0, -1).join(' ');
    }

    $.ajax({
        url: config.CommsAPI,
        type: "POST",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify({
            Body: form.message.val(),
            Email: form.email.val(),
            Name: form.name.val(),
            Phone: form.contact.val()
        }),
        cache: false,
        success: function () {
            // Success message
            $('#success').html("<div class='alert alert-success'>");
            $('#success > .alert-success').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
                .append("</button>");
            $('#success > .alert-success')
                .append("<strong>Your message has been sent. </strong>");
            $('#success > .alert-success')
                .append('</div>');
            //clear all fields
            form.id.trigger("reset");
        },
        error: function (err) {
            console.log(err);
            // Fail message
            $('#success').html("<div class='alert alert-danger'>");
            $('#success > .alert-danger').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
                .append("</button>");
            $('#success > .alert-danger').append($("<strong>").text("Sorry " + firstName + ", it seems that my mail server is not responding. Please try again later!"));
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

function tabClick(e) {
    e.preventDefault();
    $(this).tab("show");
}

$('#name').focus(function () {
    $('#success').html('');
});