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
      button.textContent = root.classList.contains("dark") ? "Light" : "Dark"
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
