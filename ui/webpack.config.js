const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const root = path.join(__dirname);


module.exports = {
  context: path.join(root, 'src'),
  mode: 'development',
  entry: 'index.tsx',
  resolve: {
    extensions: ['.tsx', '.ts', '.js', '.jsx'],
    modules: [
      path.resolve('./src'),
      path.resolve('./node_modules'),
    ],
  },
  module: {
    rules: [
      {
        exclude: /node_modules/,
        test: /\.(ts|tsx)$/,
        use: [{
          loader: 'babel-loader',
          options: { cacheDirectory: true }
        }],
      },
    ]
  },
  plugins: [
    new HtmlWebpackPlugin()
  ]
};
