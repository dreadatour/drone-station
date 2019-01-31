const path = require('path')
const merge = require('webpack-merge')
const htmlWebpackPlugin = require('html-webpack-plugin')
const htmlWebpackHarddiskPlugin = require('html-webpack-harddisk-plugin')

const common = require('./webpack.common.js')

require('dotenv').config()

const ENV = process.env.NODE_ENV || 'development'

module.exports = merge(common, {
  mode: 'development',

  devtool: 'inline-source-map',

  module: {
    rules: [
      { enforce: 'pre', test: /\.js$/, loader: 'source-map-loader' }
    ]
  },

  plugins: [
    new htmlWebpackPlugin({
      template: path.resolve(__dirname, 'index.ejs'),
      filename: path.resolve(__dirname, 'dist/index.html'),
      'ENV': ENV,
      'VERSION': process.env.VERSION || '',
      'API_URL': process.env.API_URL,
    }),
    new htmlWebpackHarddiskPlugin(),
  ],

  devServer: {
    contentBase: 'dist',
    publicPath: '/',
    historyApiFallback: true,
    host: '127.0.0.1',
    port: 8081,
    proxy: {
      '/api/': {
        target: 'http://127.0.0.1:8080',
        secure: false
      }
    }
  }
})
