const path = require('path')
const merge = require('webpack-merge')
const uglifyJSPlugin = require('uglifyjs-webpack-plugin')
const htmlWebpackPlugin = require('html-webpack-plugin')
const htmlWebpackHarddiskPlugin = require('html-webpack-harddisk-plugin')

const common = require('./webpack.common.js')

const ENV = process.env.NODE_ENV || 'production'

module.exports = merge(common, {
  mode: 'production',

  plugins: [
    new htmlWebpackPlugin({
      template: path.resolve(__dirname, 'index.ejs'),
      filename: path.resolve(__dirname, 'dist/index.html'),
      inject: false,
      alwaysWriteToDisk: true,
      'ENV': ENV,
      'VERSION': process.env.VERSION || '',
      'API_URL': process.env.API_URL,
    }),
    new htmlWebpackHarddiskPlugin(),
    new uglifyJSPlugin()
  ]
})
