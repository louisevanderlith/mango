import FormState from './formState.js';
import lookup from './pathLookup.js';

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
    startCarousel();
});

function registerEvents() {
    form.sendButton.on('click', trySend);

    $('a[data-toggle="tab"]').on('click', tabClick);

    let validForm = form.id.validator();
    validForm.on('invalid.bs.validator', fs.onValidate);
    validForm.on('valid.bs.validator', fs.onValidate);

    $(document).on('keyup', pressEnter);
}

function trySend(e) {
    form.id.validator('validate');

    if (fs.isFormValid()) {
        let toEmail = e.target.data.to;
        submitSend(toEmail);
    }
}

function pressEnter(e) {
    if (e.key !== 'Enter')
        return;

    trySend();
    e.preventDefault();
}

function submitSend(toEmail) {
    fs.submitDisabled(true);

    var firstName = form.name.val();

    if (firstName.indexOf(' ') >= 0) {
        firstName = name.split(' ').slice(0, -1).join(' ');
    }

    $.ajax({
        url: lookup.buildPath('Comms.API', "message"),
        type: "POST",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify({
            Body: form.message.val(),
            Email: form.email.val(),
            Name: form.name.val(),
            Phone: form.contact.val(),
            To: toEmail
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
            $('#success > .alert-danger').append($("<strong>").text("Sorry " + firstName + ", it seems that our mail server is not responding. Please try again later!"));
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

function startCarousel(){
    let idx = 0;
    let sliders = $('.masthead>.slider');

    if (sliders.length === 1){
        showNext(sliders, idx);
    } else {
        setInterval(() => {
            idx = showNext(sliders, idx)
        }, 3000); 
    }
}

function showNext(carouselImages, idx) {
    const mastHeader = $('header.masthead');
    const header = $(mastHeader[0]);
    const image = $(carouselImages[idx]).data('img');

    header.css('background-image', `url("${image}")`);

    const carouselLimit = carouselImages.length - 1;
    return idx === carouselLimit ? 0 : (idx + 1);
}

$('#name').focus(function () {
    $('#success').html('');
});