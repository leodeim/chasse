/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './src/**/*.{js,jsx,ts,tsx}',
    'node_modules/flowbite-react/**/*.{js,jsx,ts,tsx}',
  ],
  theme: {
    colors: {
      'cGreen': '#8ba28c',
      'cDarkGreen': '#7E937F',
      'cDarkDarkGreen': '#687a69',
      'cYellow': '#f0d9b5',
      'cDarkYellow': '#d1bd9d',
      'cBlack': '#000000',
      'cGrey': '#7b7b7b',
      'cRed': '#c9635d',
      'cBlue': '#59a4de'
    },
    extend: {},
  },
  plugins: [
    require('flowbite/plugin')
]
}
