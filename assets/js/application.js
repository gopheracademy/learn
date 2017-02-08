require("bootstrap/dist/js/bootstrap.js");

$(() => {

  $(".slide-selector").click((e) => {
    let a = $(e.target);
    advance(a, 0);
  });

  $(".slide-next").click((e) => {
    let a = $(e.target);
    let s = a.closest(".slide");
    advance(s, 1);
  });

  $(".slide-previous").click((e) => {
    let a = $(e.target);
    let s = a.closest(".slide");
    advance(s, -1);
  });

  function advance(s, offest) {
    let index = s.data("index");
    let module = s.data("module");
    $(".slide").hide();
    $(`#${module} [data-index='${index + offest}']`).show();
  };

  $(".highlight pre").each(function(i, block) {
    let html = block.innerHTML;
    html = html.replace(/\t/g, "  ");
    block.innerHTML = html;
    hljs.highlightBlock(block);
  });

});
