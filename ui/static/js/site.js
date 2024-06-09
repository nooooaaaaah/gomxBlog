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

// Function to add copy button to each code block
function addCopyButton() {
  var codeBlocks = document.querySelectorAll(".mx-auto.break-words pre");
  codeBlocks.forEach(function (block) {
    // Wrap the code block in a container
    var codeContainer = document.createElement("div");
    codeContainer.className = "relative border border-midnight-700";
    block.parentNode.insertBefore(codeContainer, block);
    codeContainer.appendChild(block);

    // Style the pre block for better visuals
    block.className += " bg-opacity-95 rounded-lg";

    // Create the copy button
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

// Function to copy content of the code block
function copyContent(block) {
  var content = block.innerText;
  var copyTextarea = document.createElement("textarea");
  copyTextarea.value = content;
  document.body.appendChild(copyTextarea);
  copyTextarea.select();
  document.execCommand("copy");
  document.body.removeChild(copyTextarea);
  alert("Code copied to clipboard!");
}

if (window.location.pathname.includes("/blog")) {
  document.addEventListener("DOMContentLoaded", (event) => {
    // Highlight code
    document.querySelectorAll("pre code").forEach((block) => {
      hljs.highlightElement(block);
    });

    // Add line numbers
    document.querySelectorAll("pre code").forEach((block) => {
      hljs.lineNumbersBlock(block);
    });
  });
  addCopyButton();
}
