
module.exports = {
  entry: "./task.js",
  output: {
    path: '../../public/js/',
    filename: "task.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}