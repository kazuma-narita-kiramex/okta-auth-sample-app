require('dotenv').config()

export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'Okta Auth Sample App',
    htmlAttrs: {
      lang: 'ja',
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },

  publicRuntimeConfig: {
    SECRET_KEY: process.env.SECRET_KEY,
    OAUTH_ISSUER: process.env.OAUTH_ISSUER,
    CLIENT_ID: process.env.CLIENT_ID,
    BACKEND_SERVER: process.env.BACKEND_SERVER || 'http://localhost:3001"',
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: ['~/assets/css/style.css'],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    '@nuxt/typescript-build',
    // https://go.nuxtjs.dev/tailwindcss
    '@nuxtjs/tailwindcss',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    '@nuxtjs/auth-next',
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    // Workaround to avoid enforcing hard-coded localhost:3000: https://github.com/nuxt-community/axios-module/issues/308
    baseURL: '/',
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {},

  auth: {
    redirect: {
      login: '/login',
      logout: '/',
      rewriteRedirects: false,
      fullPathRedirect: true,
    },
    strategies: {
      cognito: {
        scheme: 'oauth2',
        endpoints: {
          authorization: process.env.OAUTH_ISSUER + '/login',
          token: process.env.OAUTH_ISSUER + '/oauth2/token',
          userInfo: process.env.OAUTH_ISSUER + '/oauth2/userInfo',
          logout: process.env.OAUTH_ISSUER + '/logout',
        },
        token: {
          property: 'access_token',
          type: 'Bearer',
          maxAge: 1800,
        },
        responseType: 'token',
        grantType: 'authorization_code',
        clientId: process.env.CLIENT_ID,
        scope: ['openid', 'email'],
        codeChallengeMethod: 'S256',
        autoLogout: true,
      },
    },
  },
}
