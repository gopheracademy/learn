// poly fill for some of the old jquery stuff this theme uses:
jQuery.fn.load = function(callback){ $(window).on("load", callback) };

require("bootstrap/dist/js/bootstrap.js");
require("./jquery.isotope.min.js");
require("./jquery.touchSwipe.min.js");
require("./jquery.isotope.min.js");
require("./functions.min.js");
$(() => {
  activateSideNav();
});

function activateSideNav() {
  let loc = window.location;
  let path = loc.pathname;
  $(".nav li").removeClass("active");
  $(`.nav a[href='${path}']`).closest("li").addClass("active");
  $(`#contentTabs :eq(0)`).addClass("active")
  $(`#tabContents :eq(0)`).addClass("active")
}

