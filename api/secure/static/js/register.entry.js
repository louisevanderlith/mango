var form = {
    id: $('#frmRegister'),
    name: $('#txtName'),
    contact: $('#txtContact'),
    email: $('#txtEmail'),
    password: $('#txtPassword'),
    confirmPass: $('#txtConfirmPass'),
    registerButton: $('#btnRegister')
};

$(document).ready(() => {
    form.id.jqBoostrapValidation(getValidation())
});

function getValidation(){
    return {
        preventSubmit: true,
        submitError: function ($form, event, errors) {
          // additional error messages or events
        },
        submitSuccess: function ($form, event) {
          event.preventDefault(); // prevent default submit behaviour
          
          $this = form.registerButton;
          $this.prop("disabled", true); // Disable submit button until AJAX call is complete to prevent duplicate messages
          submitRegister();
        },
        filter: function () {
          return $(this).is(":visible");
        },
      };
}

function submitRegister(){
    $.ajax({
        url: "/Register",
        type: "POST",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify({
            Name: form.name.val(),
            Email: form.email.val(),
            ContactNumber: form.contact.val(),
            Password: form.password.val(),
            PasswordRepeat: form.confirmPass.val()
        }),
        cache: false,
        success: function () {
          // Success message
          $('#success').html("<div class='alert alert-success'>");
          $('#success > .alert-success').html("<button type='button' class='close' data-dismiss='alert' aria-hidden='true'>&times;")
            .append("</button>");
          $('#success > .alert-success')
            .append("<strong>Thank you. You have been successfully registered.</strong>");
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
          $('#success > .alert-danger').append($("<strong>").text("Sorry, it seems something went wrong. Please try again."));
          $('#success > .alert-danger').append('</div>');
          //clear all fields
          form.id.trigger("reset");
        },
        complete: function () {
          setTimeout(function () {
            form.registerButton.prop("disabled", false); // Re-enable submit button when AJAX call is complete
          }, 1000);
        }
      });
}