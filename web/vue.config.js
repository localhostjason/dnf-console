const { defineConfig } = require('@vue/cli-service')

const path = require('path')

function resolve(dir) {
  return path.join(__dirname, dir)
}

const port = 8080 // dev port

module.exports = defineConfig({
  transpileDependencies: true,

  publicPath: '/static',
  outputDir: 'dist',
  assetsDir: 'lib',
  lintOnSave: process.env.NODE_ENV === 'development',
  productionSourceMap: false,
  devServer: {
    port,
    open: false,
    hot: true,
    client: {
      overlay: {
        warnings: false,
        errors: true
      }
    },
    proxy: process.env.VUE_APP_BASE_API
  },
  configureWebpack: {
    // provide the app's title in webpack name field, so that
    // it can be accessed in index.html to inject the correct title.
    name: 'APP',
    resolve: {
      alias: {
        '@': resolve('src')
      },
      fallback: {
        path: require.resolve('path-browserify')
      }
    },
    performance: {
      hints: false
    }
  },
  chainWebpack: config => {
    config.resolve.alias.set('@', resolve('src'))

    // set svg-sprite-loader
    config.module.rule('svg').exclude.add(resolve('src/icons')).end()
    config.module
      .rule('icons')
      .test(/\.svg$/)
      .include.add(resolve('src/icons'))
      .end()
      .use('svg-sprite-loader')
      .loader('svg-sprite-loader')
      .options({
        symbolId: 'icon-[name]'
      })
      .end()
  },
  pluginOptions: {
    // 第三方插件配置
    // ...
  }
})
