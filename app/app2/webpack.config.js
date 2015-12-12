
module.exports = {
  entry: "./app2.js",
  output: {
    path: '../../public/js/',
    filename: "app2.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}