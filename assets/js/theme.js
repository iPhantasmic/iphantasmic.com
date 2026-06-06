(function () {
  var storageKey = "theme"
  var root = document.documentElement
  var saved = window.localStorage.getItem(storageKey)
  var prefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches

  if (saved === "dark" || (!saved && prefersDark)) {
    root.classList.add("dark")
  }

  function updateButton() {
    var button = document.querySelector("[data-theme-toggle]")
    if (button) {
      var label = root.classList.contains("dark") ? "Switch to light mode" : "Switch to dark mode"
      button.setAttribute("aria-label", label)
      button.setAttribute("title", label)
    }
  }

  window.addEventListener("DOMContentLoaded", function () {
    updateButton()

    var button = document.querySelector("[data-theme-toggle]")
    if (!button) return

    button.addEventListener("click", function () {
      root.classList.toggle("dark")
      window.localStorage.setItem(storageKey, root.classList.contains("dark") ? "dark" : "light")
      updateButton()
    })
  })
})()
