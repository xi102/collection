module.exports = {
  devServer: {
    public:'0.0.0.0:3000',
    port:3000,
    overlay: {
      warnings: true,
      errors: true
    },
    // proxy: {
    //   '/upload': {
    //     target: 'http://0.0.0.0:8080',
    //     changeOrigin: true,
    //     ws: true,
    //     pathRewrite: {
    //       '^/upload': ''
    //     }
    //   }
    // }
  },
  transpileDependencies: [
    'vuetify'
  ]
}
