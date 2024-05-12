function toggleSidebar() {
  const sidebar = document.getElementById("sidebar");
  const openSidebarBtn = document.getElementById("openSidebarBtn");
  const isSidebarClosed = sidebar.classList.contains("animate-slideOut");
  const mainContent = document.getElementById("dynamic-content");
  const newState = isSidebarClosed ? "open" : "closed";

  setSidebarAndButtonState(sidebar, openSidebarBtn, newState, mainContent);
}

function setSidebarAndButtonState(sidebar, openSidebarBtn, state, mainContent) {
  sidebar.classList.toggle("animate-slideOut", state === "closed");
  sidebar.classList.toggle("animate-slideIn", state === "open");
  openSidebarBtn.classList.toggle("rotate-180", state === "open");
  sidebar.classList.toggle("w-54", state === "open");
  mainContent.classList.toggle("animate-contentShiftRight", state === "open");
  mainContent.classList.toggle("animate-contentShiftLeft", state === "closed");
}

// Add click event listener to the open sidebar button
document
  .getElementById("openSidebarBtn")
  .addEventListener("click", toggleSidebar);

document.addEventListener("DOMContentLoaded", () => {
  const sidebar = document.getElementById("sidebar");
  const openSidebarBtn = document.getElementById("openSidebarBtn");
  const urlParams = new URLSearchParams(window.location.search);
  const sidebarState = urlParams.get("sidebar") || "closed"; // Default to closed if undefined
  const mainContent = document.getElementById("dynamic-content");

  setSidebarAndButtonState(sidebar, openSidebarBtn, sidebarState, mainContent);
});
