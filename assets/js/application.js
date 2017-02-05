require("bootstrap/dist/js/bootstrap.js");

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

