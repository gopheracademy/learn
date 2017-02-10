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

  $(document).keydown((e) => {
    switch (e.which) {
      case 37:
        //left
        a = $(".slide-previous:visible");
        advanceSlide(a, -1);
        break;

      case 39:
        //right
        a = $(".slide-next:visible");
        advanceSlide(a, 1);
        break;

      default:
        //exit this handler for other keys
        return;
    }
    //prevent the default action (scroll / move caret)
    e.preventDefault();
  });

  let hash = window.location.hash;
  if (hash != "") {
    let p = hash.split("-");
    $(".collapse").collapse("hide");
    showSlide(p[0].replace("#", ""), p[1]);
  }

});
