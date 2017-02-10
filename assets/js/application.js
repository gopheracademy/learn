require("bootstrap/dist/js/bootstrap.js");

const showSlide = (module, index) => {
  $(".slide").hide();

  let hash = `#${module}-${index}`;
  window.location.hash = hash;

  let $el = $(`#${module} [data-index='${index}']`);
  if ($el.length) {
    $el.show();
    $(`#collapse${module}`).collapse("show");
  } else {
    $("#welcome").show();
  }
};

const advanceSlide = (s, offest) => {
  let index = s.data("index");
  let module = s.data("module");
  index += offest;
  showSlide(module, index);
};

$(() => {

  $(".slide-selector").click((e) => {
    let a = $(e.target);
    advanceSlide(a, 0);
  });

  $(".slide-next a").click((e) => {
    e.preventDefault();
    let a = $(e.target);
    if (!a.hasClass(".slide-next")) {
      a = a.closest(".slide-next");
    }
    advanceSlide(a, 1);
  });

  $(".slide-previous a").click((e) => {
    e.preventDefault();
    let a = $(e.target);
    if (!a.hasClass(".slide-previous")) {
      a = a.closest(".slide-previous");
    }
    advanceSlide(a, -1);
  });


  $(".highlight pre").each((i, block) => {
    let html = block.innerHTML;
    html = html.replace(/\t/g, "  ");
    block.innerHTML = html;
    hljs.highlightBlock(block);
  });

  let hash = window.location.hash;
  if (hash != "") {
    let p = hash.split("-");
    $(".collapse").collapse("hide");
    showSlide(p[0].replace("#", ""), p[1]);
  }

});
