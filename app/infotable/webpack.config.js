
module.exports = {
  entry: "./infotable.js",
  output: {
    path: '../../public/js/',
    filename: "infotable.js"
  },

  module: {
    loaders: [
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}