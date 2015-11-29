module.exports = {
  entry: "./smartform.js",
  output: {
    path: './',
    filename: "smartform_test.js"
  },

  module: {
    loaders: [
    
      { test: /\.vue$/, loader: "vue-loader" }
    ]
  }
}