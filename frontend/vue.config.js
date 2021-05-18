module.exports = {
  devServer: {
    public:'0.0.0.0',
    port:3000,
    overlay: {
      warnings: true,
      errors: true
    },
  },
  transpileDependencies: [
    'vuetify'
  ]
}
