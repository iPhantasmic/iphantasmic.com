(function () {
  function ready(callback) {
    if (document.readyState === "loading") {
      document.addEventListener("DOMContentLoaded", callback)
      return
    }
    callback()
  }

  ready(function () {
    var form = document.querySelector("[data-search-form]")
    if (!form) return

    var input = form.querySelector("[data-search-input]")
    if (!input) return

    var timer
    var controller

    function currentResults() {
      return document.querySelector("[data-search-results]")
    }

    function setBusy(busy) {
      var results = currentResults()
      if (!results) return
      results.setAttribute("aria-busy", busy ? "true" : "false")
      results.classList.toggle("is-loading", busy)
    }

    function search() {
      var query = input.value.trim()
      var url = new URL(form.action, window.location.origin)
      if (query) {
        url.searchParams.set("q", query)
      }
      url.searchParams.set("partial", "1")

      if (controller) {
        controller.abort()
      }
      controller = new AbortController()
      setBusy(true)

      window.fetch(url.toString(), {
        headers: { "X-Requested-With": "fetch" },
        signal: controller.signal
      }).then(function (response) {
        if (!response.ok) {
          throw new Error("Search request failed")
        }
        return response.text()
      }).then(function (html) {
        var results = currentResults()
        if (results) {
          results.outerHTML = html
        }

        var pageURL = new URL(form.action, window.location.origin)
        if (query) {
          pageURL.searchParams.set("q", query)
        }
        window.history.replaceState({}, "", pageURL.pathname + pageURL.search)
      }).catch(function (error) {
        if (error.name !== "AbortError") {
          setBusy(false)
        }
      })
    }

    function queueSearch() {
      window.clearTimeout(timer)
      timer = window.setTimeout(search, 180)
    }

    input.addEventListener("input", queueSearch)
    form.addEventListener("submit", function (event) {
      event.preventDefault()
      window.clearTimeout(timer)
      search()
    })
  })
})()
