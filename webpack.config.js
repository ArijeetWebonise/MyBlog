
const Path = require('path');
const Webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');
require('babel-polyfill');

const DIST_DIR = Path.resolve(__dirname, 'build');
const SRC_DIR = Path.resolve(__dirname, 'src');
const PUBLIC_DIR = Path.resolve(__dirname, 'public');

const config = {
  entry: ['babel-polyfill', `${SRC_DIR}/index.jsx`],
  output: {
    filename: 'bundle.js',
    path: DIST_DIR,
    publicPath: '/',
  },
  devtool: 'source-map',
  devServer: {
    contentBase: DIST_DIR,
    hot: true,
    inline: true,
    open: true,
    historyApiFallback: true,
    disableHostCheck: true,
  },
  resolve: {
    extensions: ['.js', '.jsx'],
  },
  module: {
    rules: [
      {
        test: [/\.jsx$/, /\.js$/],
        use: ['babel-loader'],
        exclude: [/node_modules/],
      },
      {
        test: /\.(css|scss|sass)$/,
        use: ExtractTextPlugin.extract({
          use: [{
            loader: 'css-loader',
            options: {
              sourceMap: true,
            },
          }, {
            loader: 'sass-loader',
            options: {
              sourceMap: true,
            },
          }],
          fallback: 'style-loader',
        }),
      },
      {
        test: /\.(eot|ttf|woff|woff2)$/,
        use: 'file-loader?name=[name]-[hash].[ext]&outputPath=assets/fonts/&publicPath=/',
      },
      {
        test: /\.(jpg|png|svg|gif)$/,
        use: 'file-loader?name=[name]-[hash].[ext]&publicPath=/',
      },
    ],
  },
  plugins: [
    new Webpack.HotModuleReplacementPlugin(),
    new HtmlWebpackPlugin({
      template: `${PUBLIC_DIR}/index.html`,
      filename: `${DIST_DIR}/index.html`,
    }),
    new ExtractTextPlugin({
      filename: 'styles/style.css',
      allChunks: true,
      disable: true,
    }),
  ],
};

module.exports = config;
