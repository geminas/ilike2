module.exports = {
  entry: "./formbuilder.js",
  output: {
    path: './',
    filename: "formbuilder_test.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}