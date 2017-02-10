require("bootstrap/dist/js/bootstrap.js");
require("./slides.js");

$(() => {
  $(".highlight pre").each((i, block) => {
    let html = block.innerHTML;
    html = html.replace(/\t/g, "  ");
    block.innerHTML = html;
    hljs.highlightBlock(block);
  });

  let rx = new RegExp(/(\/[^\/]+).*/);
  let res = rx.exec(window.location.pathname);
  if (res.length >= 2) {
    let n = $(`.nav li[data-match="${res[1]}"]`);
    if (n.length > 0) {
      $(".nav li").removeClass("active");
      n.addClass("active");
    }
  }
});
