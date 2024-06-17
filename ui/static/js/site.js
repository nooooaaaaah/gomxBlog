document.addEventListener("DOMContentLoaded", function () {
  // Navbar logic
  var currentPage = window.location.pathname;
  var navbar = document.getElementById("navbar");

  if (navbar && !currentPage.includes("/blog")) {
    navbar.classList.add("fixed");
  }

  // Modal logic
  const modal = document.getElementById("coffeeModal");
  const openBtn = document.getElementById("buyCoffeeBtn");
  const closeBtn = document.getElementById("closeModalBtn");

  if (openBtn && modal && closeBtn) {
    // Open modal on button click
    openBtn.onclick = function () {
      modal.classList.remove("hidden");
      modal.classList.add("flex");
    };

    // Close modal on close button click
    closeBtn.onclick = function () {
      modal.classList.add("hidden");
      modal.classList.remove("flex");
    };

    // Close modal on clicking outside the modal content
    window.onclick = function (event) {
      if (event.target == modal) {
        modal.classList.add("hidden");
        modal.classList.remove("flex");
      }
    };
  }
  function adjustIframeHeight() {
    const availableHeight = window.innerHeight - 120; // 120px for padding and header space
    iframe.style.height = `${availableHeight}px`;
  }

  // Adjust iframe height on window resize
  window.addEventListener("resize", adjustIframeHeight);

  // Menu toggle logic
  const menuToggle = document.getElementById("menu-toggle");
  const menuClose = document.getElementById("menu-close");
  const menuModal = document.getElementById("menu-modal");

  if (menuToggle && menuModal) {
    menuToggle.addEventListener("click", function () {
      menuModal.classList.toggle("hidden");
      menuModal.classList.toggle("opacity-0");
      menuModal.classList.toggle("opacity-100");
    });
  }

  if (menuClose && menuModal) {
    menuClose.addEventListener("click", function () {
      menuModal.classList.add("hidden");
      menuModal.classList.add("opacity-0");
      menuModal.classList.remove("opacity-100");
    });
  }

  document.querySelectorAll("#menu-modal a").forEach(function (link) {
    link.addEventListener("click", function () {
      menuModal.classList.add("hidden");
      menuModal.classList.add("opacity-0");
      menuModal.classList.remove("opacity-100");
    });
  });

  // Add copy button to code blocks
  function addCopyButton() {
    var codeBlocks = document.querySelectorAll(".mx-auto.break-words pre");
    codeBlocks.forEach(function (block) {
      var codeContainer = document.createElement("div");
      codeContainer.className = "relative border border-midnight-700";
      block.parentNode.insertBefore(codeContainer, block);
      codeContainer.appendChild(block);

      block.className += " bg-opacity-95 rounded-lg";

      var copyButton = document.createElement("button");
      copyButton.textContent = "Copy";
      copyButton.className =
        "absolute top-0 right-0 m-2 bg-transparent text-peach-400 border border-peach-400 rounded-lg py-1 px-3 transition-colors duration-200 hover:bg-peach-400 hover:text-midnight-900 focus:outline-none";
      copyButton.onclick = function () {
        copyContent(block);
      };
      codeContainer.appendChild(copyButton);
    });
  }

  function copyContent(block) {
    var content = block.innerText;
    navigator.clipboard
      .writeText(content)
      .then(function () {
        alert("Code copied to clipboard!");
      })
      .catch(function (err) {
        console.error("Could not copy text: ", err);
      });
  }

  if (window.location.pathname.includes("/blog")) {
    document.querySelectorAll("pre code").forEach((block) => {
      hljs.highlightElement(block);
    });
    document.querySelectorAll("pre code").forEach((block) => {
      hljs.lineNumbersBlock(block);
    });
    addCopyButton();
  }
});
