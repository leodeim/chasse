/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './src/**/*.{js,jsx,ts,tsx}',
    'node_modules/flowbite-react/**/*.{js,jsx,ts,tsx}',
  ],
  theme: {
    colors: {
      'colorMain': '#9fa8b7',
      'colorMainDark': '#919bab',
      'colorMainLight': '#b6becc',
      'colorSecondary': '#ede4e4',
      'colorSecondaryDark': '#e3d8d8',
      'colorDetails': '#735764',
      'colorRed': '#c9635d',
      'colorBlue': '#59a4de',
    },
    extend: {},
  },
  plugins: [
    require('flowbite/plugin')
]
}
