// Snowpack Configuration File
// See all supported options: https://www.snowpack.dev/reference/configuration

/** @type {import("snowpack").SnowpackUserConfig } */
module.exports = {
  mount: {
    "lib/web/assets": "/",
  },
  plugins: ["@snowpack/plugin-postcss", "@snowpack/plugin-sass"],
  packageOptions: {
    /* ... */
  },
  devOptions: {
    tailwindConfig: "./tailwind.config.js",
  },
  buildOptions: {
    out: "static",
    watch: process.env.NODE_ENV === "development",
  },
};
