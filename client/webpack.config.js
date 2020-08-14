const webpack = require('webpack');
const path = require('path');

const config = {
  mode: 'development',
  entry: {
      index: './src/index.js',
      game: './src/game.js'
  },
  output: {
    path: path.resolve(__dirname, 'dist')
  }
};

module.exports = config;