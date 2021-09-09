module.exports = {
  css: {
    loaderOptions: {
      sass: {
        implementation: require("sass"),
        prependData: `@import "~@/scss/_variables.scss"; @import "~@/scss/tools/_functions.scss"; @import "~@/scss/tools/_mixins.scss"; @import "~@/scss/_fonts.scss";`,
      },
    },
  },
};
