import { Amplify } from 'aws-amplify'
import { Plugin } from '@nuxt/types'

const amplify: Plugin = (context, inject) => {
  Amplify.configure({
    Auth: {
      region: context.nuxtState.config.cognito.Region,
      userPoolId: context.nuxtState.config.cognito.UserPoolId,
      userPoolWebClientId: context.nuxtState.config.cognito.ClientId,
      oauth: {
        domain: context.nuxtState.config.cognito.AppWebDomain,
        scope: context.nuxtState.config.cognito.TokenScopesArray,
        redirectSignIn: context.nuxtState.config.cognito.RedirectUriSignIn,
        redirectSignOut: context.nuxtState.config.cognito.RedirectUriSignOut,
        responseType: 'code',
      },
    },
  })
}

export default amplify
