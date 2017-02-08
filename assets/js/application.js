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
    index = index + offest;
    show(module, index);
  };

  function show(module, index) {
    $(".slide").hide();
    window.location.hash = `#${module}-${index}`;
    console.log("window.location.hash:", window.location.hash);
    let $el = $(`#${module} [data-index='${index}']`)
    if ($el.length) {
      $el.show();
      $(`#collapse${module}`).collapse("show");
    } else {
      $("#welcome").show();
    }
  }

  $(".highlight pre").each(function(i, block) {
    let html = block.innerHTML;
    html = html.replace(/\t/g, "  ");
    block.innerHTML = html;
    hljs.highlightBlock(block);
  });

  let hash = window.location.hash;
  if (hash != "") {
    let p = hash.split("-");
    $(".collapse").collapse("hide");
    show(p[0].replace("#", ""), p[1]);
  }

});
