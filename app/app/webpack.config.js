module.exports = {
  entry: "./app.js",
  output: {
    path: '../../public/js/',
    filename: "app.js"
  },

  module: {
    loaders: [
    
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}