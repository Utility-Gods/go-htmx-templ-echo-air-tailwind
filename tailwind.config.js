/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["template/**/*.templ"],
  theme: {
    extend: {
      fontFamily: {
        serif: ["Roboto", "serif"],
        sans: ["Roboto", "sans-serif"],
        mono: ["Roboto", "monospace"],
      },
    },
  },
  plugins: [],
};
