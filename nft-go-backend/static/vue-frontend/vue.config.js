const { defineConfig } = require('@vue/cli-service')

module.exports = defineConfig({
    transpileDependencies: true,

    // 输出目录配置
    outputDir: '../dist',

    // 静态资源目录
    assetsDir: 'assets',

    // 开发服务器配置
    devServer: {
        port: 8081,
        proxy: {
            '/api': {
                target: 'http://localhost:8080',
                changeOrigin: true
            }
        }
    },

    // 生产环境配置
    productionSourceMap: false,

    // CSS相关配置
    css: {
        loaderOptions: {
            sass: {
                additionalData: `
          @import "@/assets/styles/variables.scss";
        `
            }
        }
    },

    publicPath: '/',

    // 禁用eslint保存检查
    lintOnSave: false
}) 