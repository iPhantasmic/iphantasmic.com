/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './internal/**/*.templ',
    './internal/**/*.go',
    './internal/posts/**/*.md',
    './assets/js/**/*.js'
  ],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        page: 'var(--bg)',
        surface: 'var(--surface)',
        ink: 'var(--text)',
        muted: 'var(--muted)',
        accent: 'var(--accent)'
      }
    }
  },
  plugins: []
}
