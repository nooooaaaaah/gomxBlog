function updateSidebarUrl(e) {
  let t = new URL(window.location);
  t.searchParams.set("sidebar", e),
    window.history.replaceState({ path: t.href }, "", t.href);
}
function getSidebarState() {
  let e = new URLSearchParams(window.location.search);
  return "closed" !== e.get("sidebar") ? "open" : "closed";
}
function setSidebarAndButtonState(e, t, a, n) {
  e.classList.toggle("animate-slideOut", "closed" === a),
    e.classList.toggle("animate-slideIn", "open" === a),
    t.classList.toggle("rotate-180", "open" === a),
    e.classList.toggle("w-54", "open" === a),
    n.classList.toggle("animate-contentShiftRight", "open" === a),
    n.classList.toggle("animate-contentShiftLeft", "closed" === a);
}
function updateUrlWithSidebarState() {
  let e = document.getElementById("sidebar"),
    t = e.classList.contains("animate-slideIn") ? "open" : "closed";
  updateSidebarUrl(t);
}
function toggleSidebar() {
  let e = document.getElementById("sidebar"),
    t = document.getElementById("openSidebarBtn"),
    a = document.getElementById("dynamic-content"),
    n = getSidebarState(),
    i = "open" === n ? "closed" : "open";
  setSidebarAndButtonState(e, t, i, a), updateSidebarUrl(i);
}
document
  .getElementById("openSidebarBtn")
  .addEventListener("click", toggleSidebar),
  document.addEventListener("DOMContentLoaded", () => {
    let e = document.getElementById("sidebar"),
      t = document.getElementById("openSidebarBtn"),
      a = document.getElementById("dynamic-content"),
      n = window.location.pathname,
      i = n.startsWith("/blogs/") && n.length > 7,
      d = window.innerWidth >= 1024,
      s;
    if (i) updateSidebarUrl((s = "closed"));
    else if (d) {
      let l = new URLSearchParams(window.location.search);
      (s = l.has("sidebar") ? getSidebarState() : "open"),
        l.has("sidebar") || updateSidebarUrl("open");
    } else updateSidebarUrl((s = "closed"));
    setSidebarAndButtonState(e, t, s, a);
  });
