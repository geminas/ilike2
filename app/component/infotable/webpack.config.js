module.exports = {
  entry: "./infotable.js",
  output: {
    path: './',
    filename: "infotable_test.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}