this.master = this.master || {};
this.master.entry = this.master.entry || {};
(function () {
'use strict';

$(document).ready(function () {
    registerEvents();
});

function registerEvents() {
    $('[data-toggle=offcanvas]').click(click_OffCanvas);
}

function click_OffCanvas() {
    $(this).toggleClass('visible-xs text-center');
    $(this).find('i').toggleClass('glyphicon-chevron-right glyphicon-chevron-left');
    $('.row-offcanvas').toggleClass('active');
    $('#lg-menu').toggleClass('hidden-xs').toggleClass('visible-xs');
    $('#xs-menu').toggleClass('visible-xs').toggleClass('hidden-xs');
    $('#btnShow').toggle();
}

}());
