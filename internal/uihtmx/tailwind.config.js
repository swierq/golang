/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: "selector",
  content: ["./ui/**/*.{html,templ,go}", "./styles.go"],
  safelist: ["text-ctp-green", "border-ctp-green"],
  theme: {
    extend: {},
  },
  plugins: [
    require("@catppuccin/tailwindcss")({
      prefix: "ctp",
    }),
  ],
};
