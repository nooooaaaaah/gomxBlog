document.addEventListener("DOMContentLoaded", function () {
  var currentPage = window.location.pathname;
  var navbar = document.getElementById("navbar");

  if (!currentPage.includes("/blog")) {
    navbar.classList.add("fixed");
  }
});

document.getElementById("menu-toggle").addEventListener("click", function () {
  var menuModal = document.getElementById("menu-modal");
  menuModal.classList.toggle("hidden");
  if (!menuModal.classList.contains("hidden")) {
    menuModal.classList.add("opacity-100");
    menuModal.classList.remove("opacity-0");
  } else {
    menuModal.classList.add("opacity-0");
    menuModal.classList.remove("opacity-100");
  }
});

document.getElementById("menu-close").addEventListener("click", function () {
  var menuModal = document.getElementById("menu-modal");
  menuModal.classList.add("hidden");
  menuModal.classList.add("opacity-0");
  menuModal.classList.remove("opacity-100");
});

document.querySelectorAll("#menu-modal a").forEach(function (link) {
  link.addEventListener("click", function () {
    var menuModal = document.getElementById("menu-modal");
    menuModal.classList.add("hidden");
    menuModal.classList.add("opacity-0");
    menuModal.classList.remove("opacity-100");
  });
});
