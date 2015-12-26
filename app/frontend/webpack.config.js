
module.exports = {
  entry: "./frontend.js",
  output: {
    path: '../../public/js/',
    filename: "frontend.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}