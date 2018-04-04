"use strict"; // Start of use strict

$(document).ready(() => {
    registerMenu();

    $("#sidenavToggler").click(toggleSide);
    $(".navbar-sidenav .nav-link-collapse").click(removeToggleClass);
    $('body.fixed-nav .navbar-sidenav, body.fixed-nav .sidenav-toggler, body.fixed-nav .navbar-collapse').on('mousewheel DOMMouseScroll', preventScroll);

    $(document).scroll(function () {
        var scrollDistance = $(this).scrollTop();
        if (scrollDistance > 100) {
            $('.scroll-to-top').fadeIn();
        } else {
            $('.scroll-to-top').fadeOut();
        }
    });

    $(document).on('click', 'a.scroll-to-top', easeToTop);
});

function registerMenu() {
    $('.toggleMenu').on('click', toggleMenu);
}

function toggleMenu(e) {
    let tag = $(e.currentTarget);
    let subID = "#" + tag.data('sub');
    console.log(subID);
    let sub = $(subID);

    tag.toggleClass('collapsed');
    sub.toggleClass('collapse');
}

function toggleSide(e) {
    e.preventDefault();
    $("body").toggleClass("sidenav-toggled");
    $(".navbar-sidenav .nav-link-collapse").addClass("collapsed");
    $(".navbar-sidenav .sidenav-second-level, .navbar-sidenav .sidenav-third-level").removeClass("show");
}

function removeToggleClass(e) {
    e.preventDefault();
    $("body").removeClass("sidenav-toggled");
}

function preventScroll(e) {
    var e0 = e.originalEvent, delta = e0.wheelDelta || -e0.detail;
    this.scrollTop += (delta < 0 ? 1 : -1) * 30;
    e.preventDefault();
}

function easeToTop(event) {
    var $anchor = $(this);
    $('html, body').stop().animate({
        scrollTop: ($($anchor.attr('href')).offset().top)
    }, 1000, 'easeInOutExpo');
    event.preventDefault();
}