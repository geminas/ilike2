module.exports = {
  entry: "./skin2.js",
  output: {
    path: './',
    filename: "skin2_test.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}