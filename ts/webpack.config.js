const { VueLoaderPlugin } = require('vue-loader');
const path = require('path');

module.exports = {
  mode: 'development',
  entry: './index.ts',
  output: {
    path: path.resolve(__dirname, '../static'),
    filename:'index.js'
  },
  resolve: {
    extensions: [ '.ts', '.js' ],
    alias: {
      'vue$': 'vue/dist/vue.esm.js'
    }
  },
  module: {
    rules: [
        {
            test: /\.ts$/,
            exclude: /node_modules/,
            loader: 'ts-loader',
            options: {
              appendTsSuffixTo: [/\.vue$/],
            },
          },

          { test: /\.vue$/,
            loader: 'vue-loader' },

          {
            test: /\.css$/,
            use: [
                "css-loader",
            ],
        },
    ]
  },
  plugins:[
    new VueLoaderPlugin()
  ]
}