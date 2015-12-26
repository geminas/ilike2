module.exports = {
  entry: "./skin1.js",
  output: {
    path: './',
    filename: "skin1_test.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}