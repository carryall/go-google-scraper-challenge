const colors = require('tailwindcss/colors')

module.exports = {
  purge: [],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      colors: {
        primary: {
          100: '#eefcf9',
          200: '#def9f7',
          300: '#c7eeee',
          400: '#b0dbde',
          500: '#92c1c9',
          600: '#6a9eac',
          700: '#497b90',
          800: '#2e5a74',
          900: '#1c4160',
        },
        olive: {
          100: '#faf9e7',
          200: '#f5f4d0',
          300: '#e2e0af',
          400: '#c5c38e',
          500: '#9f9d64',
          600: '#888649',
          700: '#726f32',
          800: '#5c591f',
          900: '#4c4913',
        },
        nude: {
          100: '#fefcf9',
          200: '#fdf9f3',
          300: '#faf3eb',
          400: '#f5ece3',
          500: '#efe2d8',
          600: '#cdaf9d',
          700: '#ac816c',
          800: '#8a5744',
          900: '#723829',
        },
        brown: {
          100: '#fdf6e4',
          200: '#fbebca',
          300: '#f3d9ac',
          400: '#e8c593',
          500: '#d9a96f',
          600: '#ba8651',
          700: '#9c6537',
          800: '#7d4823',
          900: '#683315',
        },
        crimson: {
          100: '#fbe5d2',
          200: '#f8c5a7',
          300: '#ea9977',
          400: '#d56f52',
          500: '#b93822',
          600: '#9f2118',
          700: '#851112',
          800: '#6b0a13',
          900: '#580614',
        }
      }
    },
  },
  variants: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}
