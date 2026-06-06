(function () {
  var storageKey = "theme"
  var root = document.documentElement
  var saved = window.localStorage.getItem(storageKey)
  var prefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches

  if (saved === "dark" || (!saved && prefersDark)) {
    root.classList.add("dark")
  }

  function updateButton() {
    var label = root.classList.contains("dark") ? "Switch to light mode" : "Switch to dark mode"
    document.querySelectorAll("[data-theme-toggle]").forEach(function (button) {
      button.setAttribute("aria-label", label)
      button.setAttribute("title", label)
    })
  }

  window.addEventListener("DOMContentLoaded", function () {
    updateButton()

    document.querySelectorAll("[data-theme-toggle]").forEach(function (button) {
      button.addEventListener("click", function () {
        root.classList.toggle("dark")
        window.localStorage.setItem(storageKey, root.classList.contains("dark") ? "dark" : "light")
        updateButton()
      })
    })
  })
})()
