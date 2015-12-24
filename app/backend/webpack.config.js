
module.exports = {
  entry: "./backend.js",
  output: {
    path: '../../public/js/',
    filename: "backend.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}