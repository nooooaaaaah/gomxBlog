// Utility Functions
function updateSidebarUrl(state) {
  const url = new URL(window.location);
  url.searchParams.set("sidebar", state);
  window.history.replaceState({ path: url.href }, "", url.href);
}

// Sidebar Logic
function getSidebarState() {
  const urlParams = new URLSearchParams(window.location.search);
  return urlParams.get("sidebar") !== "closed" ? "open" : "closed";
}

function setSidebarAndButtonState(sidebar, openSidebarBtn, state, mainContent) {
  sidebar.classList.toggle("animate-slideOut", state === "closed");
  sidebar.classList.toggle("animate-slideIn", state === "open");
  openSidebarBtn.classList.toggle("rotate-180", state === "open");
  sidebar.classList.toggle("w-54", state === "open");
  mainContent.classList.toggle("animate-contentShiftRight", state === "open");
  mainContent.classList.toggle("animate-contentShiftLeft", state === "closed");
}

function updateUrlWithSidebarState() {
  const sidebar = document.getElementById("sidebar");
  const newState = sidebar.classList.contains("animate-slideIn")
    ? "open"
    : "closed";
  updateSidebarUrl(newState);
}

function toggleSidebar() {
  const sidebar = document.getElementById("sidebar");
  const openSidebarBtn = document.getElementById("openSidebarBtn");
  const mainContent = document.getElementById("dynamic-content");
  const currentState = getSidebarState();
  const newState = currentState === "open" ? "closed" : "open";

  setSidebarAndButtonState(sidebar, openSidebarBtn, newState, mainContent);
  updateSidebarUrl(newState);
}

// Event Listeners
document
  .getElementById("openSidebarBtn")
  .addEventListener("click", toggleSidebar);

document.addEventListener("DOMContentLoaded", () => {
  const sidebar = document.getElementById("sidebar");
  const openSidebarBtn = document.getElementById("openSidebarBtn");
  const mainContent = document.getElementById("dynamic-content");

  // Determine if the current page is a blog post by checking the URL path
  const path = window.location.pathname;
  const isBlogPost =
    path.startsWith("/blogs/") && path.length > "/blogs/".length;

  // Determine the initial state of the sidebar based on the URL or force it to be closed on blog pages
  let sidebarState;
  if (isBlogPost) {
    sidebarState = "closed"; // Force sidebar to be closed on blog posts
  } else {
    // Default to open if not a blog post
    const urlParams = new URLSearchParams(window.location.search);
    sidebarState = urlParams.get("sidebar") === "closed" ? "closed" : "open";
  }

  // Set the sidebar and main content to the determined state
  setSidebarAndButtonState(sidebar, openSidebarBtn, sidebarState, mainContent);

  updateSidebarUrl(sidebarState);
});
