const path = require('path')
const copyWebpackPlugin = require('copy-webpack-plugin')

module.exports = {
  bail: true,

  entry: {
    app: './src/index.tsx'
  },

  output: {
    filename: 'static/js/drone-station.js',
    path: path.resolve(__dirname, 'dist'),
    publicPath: '/'
  },

  resolve: {
    modules: [
      path.resolve(__dirname, 'src'),
      path.resolve(__dirname, 'node_modules'),
    ],
    extensions: ['.ts', '.tsx', '.js', '.json']
  },

  module: {
    rules: [
      {
        test: /\.tsx?$/,
        loader: 'awesome-typescript-loader'
      },
    ]
  },

  plugins: [
    new copyWebpackPlugin([
      {
        from: 'static',
        to: path.resolve(__dirname, 'dist/static')
      }
    ]),
  ]
}
