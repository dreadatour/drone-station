'use strict'

const path = require('path')
const express = require('express')
const http = require('http')

if ((process.env.NODE_ENV || 'development') === 'development') {
  require('dotenv').config()
}

const httpPort = process.env.HTTP_PORT || 8081

const app = express()
app.use('/', express.static(path.resolve(__dirname, 'dist')))

http.createServer(app).listen(httpPort)
console.info(`Listening http requests using ${ httpPort } port...`)
